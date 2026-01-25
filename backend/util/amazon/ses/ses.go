package ses

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"os"
)

const (
	pwResetTitle = "パスワードリセット"
)

type emailContent struct {
	Subject string
	To      string
	//Bcc *[]string
	Text string
}

// ValidateSESConfig AWS SES設定が正しく設定されているかチェックします
// パスワードリセット機能を使用する前に呼び出すことを推奨します
func ValidateSESConfig() error {
	var missingVars []string

	if os.Getenv("AWS_REGION") == "" {
		missingVars = append(missingVars, "AWS_REGION")
	}
	if os.Getenv("AWS_SES_ACCESS_KEY") == "" {
		missingVars = append(missingVars, "AWS_SES_ACCESS_KEY")
	}
	if os.Getenv("AWS_SES_ACCESS_SECRET") == "" {
		missingVars = append(missingVars, "AWS_SES_ACCESS_SECRET")
	}
	if os.Getenv("AWS_SES_FROM") == "" {
		missingVars = append(missingVars, "AWS_SES_FROM")
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("以下のAWS SES環境変数が設定されていません: %v\n設定方法の詳細は document/aws-ses-setup.md を参照してください", missingVars)
	}

	return nil
}

func SendPasswordReset(ctx context.Context, email string, token string) error {
	//メールテンプレートを作る
	msg, err := pwRstTemp(&passwordReset{
		Token: token,
	})
	if err != nil {
		return fmt.Errorf("メールテンプレートの生成に失敗しました: %w", err)
	}
	//メールを送信する
	err = sendMail(ctx, &emailContent{
		Subject: pwResetTitle,
		To:      email,
		Text:    msg,
	})
	return err
}

func sendMail(ctx context.Context, content *emailContent) error {
	// 環境変数の検証
	if err := ValidateSESConfig(); err != nil {
		return err
	}

	region := os.Getenv("AWS_REGION")
	accessKey := os.Getenv("AWS_SES_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SES_ACCESS_SECRET")
	from := os.Getenv("AWS_SES_FROM")

	// 1. AWSの設定を読み込む
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region), config.WithCredentialsProvider(
		credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
	))
	if err != nil {
		return fmt.Errorf("AWS設定の読み込みに失敗しました: %w", err)
	}

	// 2. SESクライアントを作成
	client := sesv2.NewFromConfig(cfg)

	input := &sesv2.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{content.To},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Body: &types.Body{
					Text: &types.Content{
						Data: &content.Text,
					},
				},
				Subject: &types.Content{
					Data: &content.Subject,
				},
			},
		},
		FromEmailAddress: &from,
	}

	_, err = client.SendEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("メール送信に失敗しました (送信先: %s): %w", content.To, err)
	}

	return nil
}
