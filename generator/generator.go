package generator

import (
	"bytes"
	"encoding/base64"
	"image/png"

	errs "github.com/Basic-Components/qrcodegenerator/errs"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// Encode 输出编码base64的编码
func Encode(info string) (string, error) {
	var result string
	code, err := qr.Encode(info, qr.L, qr.Unicode)
	if err != nil {
		return result, err
	}
	if info != code.Content() {
		return result, errs.DataDifferentError
	}
	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		return result, err
	}
	buff := new(bytes.Buffer)
	err = png.Encode(buff, code)
	if err != nil {
		return result, err
	}
	result = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buff.Bytes())
	return result, nil
}
