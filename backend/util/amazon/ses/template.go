package ses

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

type passwordReset struct {
	Token string
}

// generatePwRstStr はパスワードリセットメール本文のテンプレートです
func generatePwRstStr() (string, error) {
	frontURI := os.Getenv("FRONT_URI")
	if frontURI == "" {
		return "", fmt.Errorf("FRONT_URI環境変数が設定されていません。詳細は document/aws-ses-setup.md を参照してください")
	}
	return "URLは " + frontURI + "/auth/password/reset/{{ .Token }} です", nil
}

func pwRstTemp(cfg *passwordReset) (string, error) {
	templateStr, err := generatePwRstStr()
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("psw-rst-template").Parse(templateStr)
	if err != nil {
		return "", err
	}

	// バッファにテンプレートの結果を出力
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, *cfg)
	if err != nil {
		return "", err
	}

	// バッファの内容を文字列として返す
	return buf.String(), nil
}
