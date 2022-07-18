package writer

import (
	"bytes"
	_ "embed"
	"go/format"
	"os"
	"text/template"

	"github.com/nikaera/generr/pkg/info"
)

//go:embed templates/errors.go.tmpl
var tmplString string

type CreateFileWithErrorWithCodesInput struct {
	Output         string
	Package        string
	ErrorWithCodes []*info.ErrorWithCode
}

func CreateFileWithErrorWithCodes(input CreateFileWithErrorWithCodesInput) error {
	fp, err := os.Create(input.Output)
	if err != nil {
		return err
	}
	defer fp.Close()

	var bb bytes.Buffer
	t, _ := template.New("errors").Parse(tmplString)
	if err := t.Execute(&bb, input); err != nil {
		return err
	}

	p, err := format.Source(bb.Bytes())
	if err != nil {
		return err
	}
	if _, err := fp.Write(p); err != nil {
		return err
	}

	return nil
}
