# GitHub Copilot 自動コードレビューの設定方法

このドキュメントでは、GitHub Copilotによる自動コードレビュー機能を有効にする方法を説明します。

## 概要

GitHub Copilotのコードレビュー機能は、プルリクエストが作成されたときに自動的にコードをレビューし、改善案やバグの可能性などをコメントとして提供します。

**重要**: GitHub Copilotのコードレビューは、GitHub Actionsワークフローではなく、**リポジトリルールセット**を通じて設定します。

## 前提条件

- GitHub Copilot Pro、Business、またはEnterpriseのサブスクリプション
- リポジトリの管理者権限

## 設定手順

### 1. リポジトリ設定にアクセス

1. GitHubでリポジトリページを開く
2. 「Settings」タブをクリック
3. 左サイドバーの「Code and automation」セクションから「Rules」をクリック

### 2. ルールセットの作成

1. 「Rulesets」タブをクリック
2. 「New ruleset」ボタンをクリック
3. 「New branch ruleset」を選択

### 3. Copilot自動レビューの有効化

1. ルールセット名を入力（例: "Copilot自動レビュー"）
2. 「Target branches」セクションで、レビューを適用するブランチを選択
   - 通常は `main` または `develop` ブランチへのPRに適用
   - デフォルトブランチを選択することを推奨
3. 「Require code review by Copilot」オプションを有効にする
4. オプション設定：
   - ✅ **Review new pushes**: 新しいコミットがプッシュされるたびに再レビュー
   - ✅ **Review draft pull requests**: ドラフトPRもレビュー対象にする
5. ルールセットの状態を **Active** に設定
6. 「Create」をクリックして保存

### 4. 動作確認

1. 新しいブランチを作成してコード変更を行う
2. プルリクエストを作成
3. Copilotが自動的にレビューを開始することを確認
4. レビューが完了すると、プルリクエストにコメントが投稿される

## よくある質問

### Q: ワークフローの承認が必要と表示される

A: GitHub Actionsワークフローではなく、リポジトリルールセットを使用してCopilotレビューを設定してください。誤ったワークフローファイル（`.github/workflows/copilot-review.yml`など）がある場合は削除してください。

### Q: レビューが自動的に開始されない

A: 以下を確認してください：
1. リポジトリルールセットが正しく設定されているか
2. ルールセットが「Active」状態になっているか
3. 対象ブランチが正しく設定されているか
4. Copilotのサブスクリプションが有効であるか

### Q: エラーで失敗する

A: GitHub Actionsワークフローを使用している場合、そのワークフローを削除し、代わりにリポジトリルールセットを使用してください。

### Q: 特定のファイルをレビューから除外したい

A: Copilotは以下のファイルを自動的に無視します：
- lockファイル（package-lock.json、yarn.lock など）
- ログファイル
- バイナリファイル
- 生成されたコード

## 参考リンク

- [GitHub公式ドキュメント: Copilot自動レビューの設定](https://docs.github.com/en/copilot/how-tos/use-copilot-agents/request-a-code-review/configure-automatic-review)
- [GitHub Copilotコードレビューについて](https://docs.github.com/en/copilot/concepts/agents/code-review)
- [リポジトリルールでの自動レビュー設定](https://github.blog/changelog/2025-09-10-copilot-code-review-independent-repository-rule-for-automatic-reviews/)

## 注意事項

- Copilotのレビューは「Comment」のみで、マージをブロックしたり、必須のコードレビュー要件を満たすものではありません
- Copilotの提案は参考情報として扱い、最終的な判断は人間のレビュアーが行ってください
- レビューの品質を向上させるため、`.github/copilot-instructions.md`でレビュー基準をカスタマイズできます
