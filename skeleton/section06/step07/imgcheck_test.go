package imgcheck_test

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	imgcheck "github.com/tenntenn/hosei24/section06/step07"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name  string
		file  string
		rules []imgcheck.Rule

		wantErrs []error
	}{
		{"format ok", "72x72.png", []imgcheck.Rule{imgcheck.Format("png")}, nil},
		{"format ng", "72x72.png", []imgcheck.Rule{imgcheck.Format("jpeg")}, []error{imgcheck.ErrFormat}},
		{"height ok", "72x72.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, nil},
		// TODO: 幅が72以下かチェックするルールが成功するパターンのテストケースを書く

		{"height ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, []error{imgcheck.ErrTooLarge}},
		{"width ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxWidth(72)}, []error{imgcheck.ErrTooLarge}},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f, err := os.Open(filepath.Join("testdata", tt.file))
			if err != nil {
				t.Fatal("予期しないエラー:", err)
			}
			t.Cleanup(func() { f.Close() })

			err = imgcheck.Validate(f, tt.rules...)
			for i := range tt.wantErrs {
				var verr *imgcheck.ValidationError
				if /* TODO: errが*imgcheck.ValidationError型に変換できるか試す*/ ||
					verr.Rule != tt.rules[i] {
					t.Errorf("期待したルールのエラーではありません: i = %d", i)
				}

				if !errors.Is(err, tt.wantErrs[i]) {
					t.Errorf("予期したエラーと異なります: want %v", tt.wantErrs[i])
				}
			}
		})
	}
}
