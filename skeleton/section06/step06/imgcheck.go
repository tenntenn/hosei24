package imgcheck

import (
	"errors"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// ベースとなるエラー
var (
	ErrFormat   = errors.New("画像フォーマットが違います")
	ErrTooLarge = errors.New("画像が大きすぎます")
)

// バリデーションエラー
type ValidationError struct {
	Rule Rule
	Err  error
}

// TODO: errorインタフェースを実装する
// 返す文字列はErrフィールドのErrorメソッドを呼び出して取得する


// バリデーションルールを表すインタフェース
type Rule interface {
	Validate(img image.Image, format string) error
}

type formatRule struct {
	format string
}

func (r *formatRule) Validate(img image.Image, format string) error {
	if r.format != format {
		return &ValidationError{
			Rule: r,
			Err:  ErrFormat,
		}
	}
	return nil
}

// 画像フォーマットをチェックするルール
func Format(format string) (rule Rule) {
	return &formatRule{format: format}
}

type maxSizeRule struct {
	height *int
	width  *int
}

func (r *maxSizeRule) Validate(img image.Image, _ string) error {
	bounds := img.Bounds()
	var err error

	if r.height != nil && bounds.Dy() > *r.height {
		err = errors.Join(err, &ValidationError{
			Rule: r,
			Err:  ErrTooLarge,
		})
	}

	// TODO: 幅のチェックを行う
	// ヒント：高さの処理とほとんど同じ

	return err
}

// 高さをチェックするルール
func MaxHeight(h int) Rule {
	return &maxSizeRule{height: &h}
}

// 幅をチェックするルール
func MaxWidth(w int) Rule {
	return &maxSizeRule{width: &w}
}

// 画像のバリデーションを行う
func Validate(r io.Reader, rules ...Rule) error {
	// 画像を読み込む
	img, format, err := image.Decode(r)
	switch {
	case err == image.ErrFormat:
		// 画像として読み込めなかった
		return nil
	case err != nil:
		return err
	}

	var rerr error
	for _, rule := range rules {
		rerr = errors.Join(rerr, rule.Validate(img, format))
	}

	return rerr
}

// ディレクトリ以下の画像ファイルのバリデーションを行う
func ValidateDir(root string, rules ...Rule) error {
	walkfunc := func(path string, info fs.FileInfo, err error) (rerr error) {

		// エラーが発生した
		if err != nil {
			return err
		}

		// ディレクトリ
		if info.IsDir() {
			return nil
		}

		// 変換前のファイルを開く
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		// 関数終了時にファイルを閉じる
		defer file.Close()

		// バリデーションをかける
		if err := Validate(file, rules...); err != nil {
			return err
		}
		return nil
	}
	return filepath.Walk(root, walkfunc)
}
