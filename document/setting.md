# 環境構築
## ディレクトリ構成
モノレポ管理です
| ディレクトリ | 説明                              |
|------------------|-----------------------------------|
| backend          | APIサーバー（バックエンド）       |
| front            | フロントエンド                    |
| db               | DBサーバー（Dockerfile があるのみ）|
| proxy            | リバースプロキシサーバーの設定    |

(自分で用意するもの)
| ディレクトリ/名称 | 説明                              |
|------------------|-----------------------------------|
| .ssl          | リバースプロキシ用の認証ファイル置き場 そのうちサンプル置く       |
 

## 言語仕様・技術スタック
| ディレクトリ | 言語                              |
|------------------|-----------------------------------|
| backend          | Go / GraphQL     |
| front            | Vue3                    |
| db               | mongoDB|

全体的に、勉学のために採用した自分が手をつけたことがない領域がメインなので、合理性とか特にないです。


## 環境構築
(正直あんま構造覚えてないので、誰かが環境構築する際に一緒にやりたい...)<br>
### コンテナ
`docker compose up -d`を実行する

![image](https://github.com/user-attachments/assets/8b2ad7ce-5e21-4871-bd1e-fc0544b85d60)

この4つのイメージが作成されるはず。


### バックエンド
何かしらの手段で実行する。デバッグでも通常の実行でも可。<br>
`air`が導入されてるのでHMRを効かせたい方は`go run air`でも可。僕は使わなくなった...<br>

### フロント
コンテナが起動してたら`https://localhost/`でアクセスできるはずです。(裏で`npm run dev`が走ってる)<br>
HMRが効いているので開発中は特に何も意識しなくてOKのはず。ちなみにバックエンドが実行されてないとデータは描画されません。

### DB
シーダーとかはバックエンド側がよしなにやってくれます。ただ単にコンテナが起動してれば動きます。

### プロキシサーバー
これも起動さえしていれば特に何も意識しなくてOK。

### Git hooks (lefthook)
プロジェクトではlefthookを使用してGit hooksを管理しています。

#### 初回セットアップ
`front`ディレクトリで`npm install`を実行すると、`prepare`スクリプトによって自動的にlefthookのGit hooksがリポジトリルート(`.git/hooks/`)にインストールされます。

もし`npm install`後に`.git/hooks/`にフックが見当たらない場合は、リポジトリルートで以下のコマンドを手動で実行してください：
```bash
npx --prefix front lefthook install
```

#### 実行タイミング
pre-pushフックでは、フロントエンドのコードに対して`npm run lint`および`npm run format`が実行され、プッシュ前にLintとフォーマットが自動で行われます。lintコマンドは自動生成ファイルを除外するよう設定されています。

#### 設定ファイル
- `front/.lefthook.yml`: フロントエンドのGit hooks設定（実際の設定）
- `lefthook.yml`（リポジトリルート）: `front/.lefthook.yml`を拡張（extends）する最小限の設定ファイル
- 将来的に`backend/.lefthook.yml`も追加可能で、それぞれが独立して管理されます

#### 一時的に無効化する方法
特定のプッシュでフックを無効化したい場合：
```bash
LEFTHOOK=0 git push
```

### 自動PR生成ツール
プロジェクトでは、Claude Codeを使用した自動PR生成機能を提供しています。

#### 必要なツール

自動PR生成機能を利用するには、以下のツールが必要です：

##### GitHub CLI (`gh`)
GitHubのPR作成に使用します（`/pr`と`/pr-sh`の両方で必須）。

**インストール:**
- 公式サイト: https://cli.github.com/
- Windows: `winget install GitHub.cli` または `choco install gh`
- macOS: `brew install gh`
- Linux: パッケージマネージャーを使用（詳細は公式サイト参照）

**認証:**
```bash
gh auth login
```

##### Claude CLI (`claude`)
AI によるPR内容生成に使用します（`/pr-sh`のみで必要）。

**インストール:**
```bash
npm install -g @anthropic-ai/claude-cli
```

**認証:**
APIキーの設定が必要です。詳細はClaude CLIのドキュメントを参照してください。

##### jq
JSON処理ツールです（`/pr-sh`のみで必要）。

**インストール:**
- Windows:
  - winget: `winget install jqlang.jq`
  - Chocolatey: `choco install jq`
  - 手動: https://jqlang.github.io/jq/download/
- macOS: `brew install jq`
- Linux: `sudo apt install jq` または `sudo yum install jq`

**確認:**
```bash
jq --version
```

#### コマンドの使い分け

##### `/pr` コマンド
Claude AIが直接gitコマンドとgh CLIを実行してPRを作成します。

**特徴:**
- 各ステップでユーザーの確認を求めるため、より安全
- `gh`と`claude`のみで動作（`jq`不要）
- 初めての利用や重要な変更に推奨

##### `/pr-sh` コマンド
シェルスクリプト（`ai-pr-draft-ja.sh`）を自動実行してPRを作成します。

**特徴:**
- 完全に自動化されており、ユーザー確認なしで実行
- Claude Codeのコンテキストを節約（試験運用中）
- すべてのツール（`gh`, `claude`, `jq`）が必要
- 日常的な作業や信頼できる環境に推奨

**デバッグモード:**
```bash
/pr-sh --debug
```

## 諸々の経緯
### リバースプロキシサーバーを採用した理由
~~正直忘れたんですが~~ 勉強のために(といいつつケチっただけ)認証を自前で組むことにした際、CORS制約をクリアするための対策として採用した。

### 開発環境でバックエンドサーバーがパイパスな理由
Golandのデバッガが見やすかったので、あんまりコンテナ依存にしたくなかったためです。めちゃくちゃ個人的な理由です。
Docker上でGolandを動かす方法もあるのですが、やってることがオーバーなので採用しませんでした。
