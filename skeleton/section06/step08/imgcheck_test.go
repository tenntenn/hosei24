package imgcheck_test

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	imgcheck "github.com/tenntenn/hosei24/section06/step08"
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
		{"format-pattern ok", "72x72.png", []imgcheck.Rule{formatPattern(t, ".+g")}, nil},
		{"format-pattern ng", "72x72.png", []imgcheck.Rule{formatPattern(t, "jpe?g")}, []error{imgcheck.ErrFormat}},
		{"height ok", "72x72.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, nil},
		{"width ok", "72x72.png", []imgcheck.Rule{imgcheck.MaxWidth(72)}, nil},
		{"height ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxHeight(72)}, []error{imgcheck.ErrTooLarge}},
		{"width ng", "300x300.png", []imgcheck.Rule{imgcheck.MaxWidth(72)}, []error{imgcheck.ErrTooLarge}},
	}

	for _, tt := range cases {
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
				if !errors.As(err, &verr) || verr.Rule != tt.rules[i] {
					t.Errorf("期待したルールのエラーではありません: i = %d", i)
				}

				if !errors.Is(err, tt.wantErrs[i]) {
					t.Errorf("予期したエラーと異なります: want %v", tt.wantErrs[i])
				}
			}
		})
	}
}

func formatPattern(t *testing.T, pattern string) imgcheck.Rule {
	t.Helper()
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatal("予期しないエラー:", err)
	}
	return imgcheck.FormatPattern(re)
}
