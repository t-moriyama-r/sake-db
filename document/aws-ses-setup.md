# AWS SES 設定ガイド

## 概要
このアプリケーションでは、パスワードリセット機能にAmazon Simple Email Service (SES)を使用しています。
SESを正しく設定しないと、パスワードリセットメールが送信できません。

## 必要な環境変数

バックエンドの`.env`ファイルに以下の環境変数を設定する必要があります：

### AWS SES関連
- `AWS_SES_ACCESS_KEY`: AWS IAMユーザーのアクセスキーID
- `AWS_SES_ACCESS_SECRET`: AWS IAMユーザーのシークレットアクセスキー
- `AWS_SES_FROM`: 送信元メールアドレス（SESで検証済みのメールアドレス）
- `AWS_REGION`: AWSリージョン（例: `ap-northeast-1`）

### フロントエンド関連
- `FRONT_URI`: フロントエンドのURL（パスワードリセットリンクに使用）

## AWS SES セットアップ手順

### 1. AWS IAMユーザーの作成

1. AWS Management Consoleにログイン
2. IAMサービスに移動
3. 「ユーザー」→「ユーザーを追加」をクリック
4. ユーザー名を入力（例: `sake-db-ses-user`）
5. 「アクセスキー - プログラムによるアクセス」を選択
6. 「次のステップ: アクセス許可」をクリック

### 2. SES権限の付与

1. 「既存のポリシーを直接アタッチ」を選択
2. 以下のポリシーをアタッチ：
   - `AmazonSESFullAccess`（開発環境の場合）
   - または、本番環境では最小権限の原則に従ったカスタムポリシーを作成

カスタムポリシーの例：
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ses:SendEmail",
        "ses:SendRawEmail"
      ],
      "Resource": "*"
    }
  ]
}
```

### 3. アクセスキーの取得

1. ユーザー作成の最後のステップで、アクセスキーIDとシークレットアクセスキーが表示されます
2. これらの値を安全に保存してください（**この画面を閉じると二度と表示されません**）
3. `.env`ファイルの対応する環境変数に設定します

### 4. メールアドレスの検証（サンドボックス環境の場合）

SESは初期状態でサンドボックスモードになっています。サンドボックスでは、検証済みのメールアドレスにのみメールを送信できます。

#### 送信元メールアドレスの検証
1. SESコンソールに移動
2. 「Email Addresses」→「Verify a New Email Address」をクリック
3. `AWS_SES_FROM`に設定するメールアドレスを入力
4. 検証メールが送信されるので、リンクをクリックして検証を完了

#### 受信者メールアドレスの検証（開発・テスト時のみ）
開発環境では、テストで使用するメールアドレスも同様に検証が必要です。

### 5. 本番環境への移行（サンドボックスの解除）

本番環境で任意のメールアドレスに送信するには、SESのサンドボックスを解除する必要があります：

1. SESコンソールで「Sending Statistics」に移動
2. 「Request Production Access」をクリック
3. 申請フォームに必要事項を記入：
   - メール送信のユースケース
   - バウンス率とクレームの管理方法
   - オプトアウトの処理方法
4. AWSの承認を待つ（通常24時間以内）

## 設定例

```bash
# backend/.env または backend/.env.dev

# AWS SES設定
AWS_SES_ACCESS_KEY=AKIAIOSFODNN7EXAMPLE
AWS_SES_ACCESS_SECRET=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
AWS_SES_FROM=liquor@trc.mixh.jp
AWS_REGION=ap-northeast-1

# フロントエンドURL（パスワードリセットリンクに使用）
FRONT_URI=https://localhost
```

## トラブルシューティング

### メールが送信されない場合

1. **環境変数の確認**
   - すべての必要な環境変数が設定されているか確認
   - アクセスキーとシークレットキーが正しいか確認

2. **メールアドレスの検証**
   - 送信元メールアドレスがSESで検証済みか確認
   - サンドボックスモードの場合、受信者メールアドレスも検証済みか確認

3. **IAM権限の確認**
   - IAMユーザーにSESの送信権限があるか確認

4. **リージョンの確認**
   - `AWS_REGION`が正しく設定されているか確認
   - SESで使用可能なリージョンか確認

5. **ログの確認**
   - バックエンドのログでエラーメッセージを確認
   - CloudWatch Logsでより詳細なエラー情報を確認

### よくあるエラー

#### `MessageRejected: Email address is not verified`
- 原因: 送信元または受信先のメールアドレスが検証されていない（サンドボックスモード）
- 解決: メールアドレスを検証するか、本番アクセスを申請

#### `InvalidClientTokenId: The security token included in the request is invalid`
- 原因: アクセスキーIDが間違っているか、無効
- 解決: 正しいアクセスキーを設定

#### `SignatureDoesNotMatch: The request signature we calculated does not match the signature you provided`
- 原因: シークレットアクセスキーが間違っている
- 解決: 正しいシークレットキーを設定

## セキュリティのベストプラクティス

1. **環境変数の管理**
   - `.env`ファイルは`.gitignore`に追加（すでに追加済み）
   - 本番環境では環境変数を安全に管理（AWS Systems Manager Parameter Store、AWS Secrets Manager等を使用）

2. **最小権限の原則**
   - IAMユーザーには必要最小限の権限のみを付与
   - 本番環境では、SES送信のみの権限に制限

3. **キーのローテーション**
   - 定期的にアクセスキーをローテーション
   - 古いキーは無効化

4. **監視とアラート**
   - SESの送信統計を監視
   - バウンス率やクレーム率が高い場合はアラートを設定

## 参考リンク

- [AWS SES 公式ドキュメント](https://docs.aws.amazon.com/ses/)
- [AWS SES サンドボックスの移動](https://docs.aws.amazon.com/ses/latest/dg/request-production-access.html)
- [AWS IAM ベストプラクティス](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html)
