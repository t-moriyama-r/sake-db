package ses

import (
	"os"
	"testing"
)

func TestValidateSESConfig(t *testing.T) {
	// テスト前の環境変数を保存
	originalRegion := os.Getenv("AWS_REGION")
	originalAccessKey := os.Getenv("AWS_SES_ACCESS_KEY")
	originalSecretKey := os.Getenv("AWS_SES_ACCESS_SECRET")
	originalFrom := os.Getenv("AWS_SES_FROM")

	// テスト後に環境変数を復元
	defer func() {
		os.Setenv("AWS_REGION", originalRegion)
		os.Setenv("AWS_SES_ACCESS_KEY", originalAccessKey)
		os.Setenv("AWS_SES_ACCESS_SECRET", originalSecretKey)
		os.Setenv("AWS_SES_FROM", originalFrom)
	}()

	tests := []struct {
		name          string
		setupEnv      func()
		expectedError bool
		expectedInMsg string
	}{
		{
			name: "すべての環境変数が設定されている場合",
			setupEnv: func() {
				os.Setenv("AWS_REGION", "ap-northeast-1")
				os.Setenv("AWS_SES_ACCESS_KEY", "test-access-key")
				os.Setenv("AWS_SES_ACCESS_SECRET", "test-secret-key")
				os.Setenv("AWS_SES_FROM", "test@example.com")
			},
			expectedError: false,
		},
		{
			name: "AWS_REGIONが未設定の場合",
			setupEnv: func() {
				os.Unsetenv("AWS_REGION")
				os.Setenv("AWS_SES_ACCESS_KEY", "test-access-key")
				os.Setenv("AWS_SES_ACCESS_SECRET", "test-secret-key")
				os.Setenv("AWS_SES_FROM", "test@example.com")
			},
			expectedError: true,
			expectedInMsg: "AWS_REGION",
		},
		{
			name: "AWS_SES_ACCESS_KEYが未設定の場合",
			setupEnv: func() {
				os.Setenv("AWS_REGION", "ap-northeast-1")
				os.Unsetenv("AWS_SES_ACCESS_KEY")
				os.Setenv("AWS_SES_ACCESS_SECRET", "test-secret-key")
				os.Setenv("AWS_SES_FROM", "test@example.com")
			},
			expectedError: true,
			expectedInMsg: "AWS_SES_ACCESS_KEY",
		},
		{
			name: "AWS_SES_ACCESS_SECRETが未設定の場合",
			setupEnv: func() {
				os.Setenv("AWS_REGION", "ap-northeast-1")
				os.Setenv("AWS_SES_ACCESS_KEY", "test-access-key")
				os.Unsetenv("AWS_SES_ACCESS_SECRET")
				os.Setenv("AWS_SES_FROM", "test@example.com")
			},
			expectedError: true,
			expectedInMsg: "AWS_SES_ACCESS_SECRET",
		},
		{
			name: "AWS_SES_FROMが未設定の場合",
			setupEnv: func() {
				os.Setenv("AWS_REGION", "ap-northeast-1")
				os.Setenv("AWS_SES_ACCESS_KEY", "test-access-key")
				os.Setenv("AWS_SES_ACCESS_SECRET", "test-secret-key")
				os.Unsetenv("AWS_SES_FROM")
			},
			expectedError: true,
			expectedInMsg: "AWS_SES_FROM",
		},
		{
			name: "すべての環境変数が未設定の場合",
			setupEnv: func() {
				os.Unsetenv("AWS_REGION")
				os.Unsetenv("AWS_SES_ACCESS_KEY")
				os.Unsetenv("AWS_SES_ACCESS_SECRET")
				os.Unsetenv("AWS_SES_FROM")
			},
			expectedError: true,
			expectedInMsg: "AWS_REGION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数をセットアップ
			tt.setupEnv()

			// テスト実行
			err := ValidateSESConfig()

			// エラーチェック
			if tt.expectedError {
				if err == nil {
					t.Errorf("エラーが期待されましたが、エラーが返されませんでした")
				} else if tt.expectedInMsg != "" {
					// エラーメッセージに期待される文字列が含まれているか確認
					if err.Error() == "" || !containsString(err.Error(), tt.expectedInMsg) {
						t.Errorf("エラーメッセージに '%s' が含まれていることが期待されましたが、実際のメッセージは: %s", tt.expectedInMsg, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが返されました: %v", err)
				}
			}
		})
	}
}

func TestGeneratePwRstStr(t *testing.T) {
	// テスト前の環境変数を保存
	originalFrontURI := os.Getenv("FRONT_URI")

	// テスト後に環境変数を復元
	defer func() {
		os.Setenv("FRONT_URI", originalFrontURI)
	}()

	tests := []struct {
		name          string
		frontURI      string
		expectedError bool
		expectedInMsg string
	}{
		{
			name:          "FRONT_URIが設定されている場合",
			frontURI:      "https://example.com",
			expectedError: false,
		},
		{
			name:          "FRONT_URIが未設定の場合",
			frontURI:      "",
			expectedError: true,
			expectedInMsg: "FRONT_URI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 環境変数をセットアップ
			os.Setenv("FRONT_URI", tt.frontURI)

			// テスト実行
			result, err := generatePwRstStr()

			// エラーチェック
			if tt.expectedError {
				if err == nil {
					t.Errorf("エラーが期待されましたが、エラーが返されませんでした")
				} else if tt.expectedInMsg != "" {
					if !containsString(err.Error(), tt.expectedInMsg) {
						t.Errorf("エラーメッセージに '%s' が含まれていることが期待されましたが、実際のメッセージは: %s", tt.expectedInMsg, err.Error())
					}
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが返されました: %v", err)
				}
				// 結果が期待通りか確認
				if !containsString(result, tt.frontURI) {
					t.Errorf("結果に FRONT_URI '%s' が含まれていることが期待されましたが、実際の結果は: %s", tt.frontURI, result)
				}
			}
		})
	}
}

// containsString は文字列に部分文字列が含まれているかチェックする
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || s[0:len(substr)] == substr || containsStringHelper(s, substr))
}

func containsStringHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
