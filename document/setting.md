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
| .ssl          | リバースプロキシ用のSSL証明書置き場（後述の「プロキシサーバー」セクションを参照）       |
 

## 言語仕様・技術スタック
| ディレクトリ | 言語                              |
|------------------|-----------------------------------|
| backend          | Go / GraphQL     |
| front            | Vue3                    |
| db               | mongoDB|

全体的に、勉学のために採用した自分が手をつけたことがない領域がメインなので、合理性とか特にないです。


## 環境構築
### コンテナ
`cd sake-db`<br/>
`docker compose up -d`を実行する

![image](https://github.com/user-attachments/assets/8b2ad7ce-5e21-4871-bd1e-fc0544b85d60)

この4つのイメージが作成されるはず。<br/>
<br/>
フロントのnpm install
`cd front`<br/>
`cd npm i`<br/>

バックエンド
`cd ../backend`<br/>
`go mod tidy`<br/>

プロキシサーバー<br/>
オレオレ証明書が必要。

#### SSL証明書の準備（初回のみ）
プロキシサーバーはHTTPSでの通信を行うため、自己署名SSL証明書が必要です。
以下の手順で証明書を生成してください：

1. ルートディレクトリに移動

2. 自己署名証明書を生成
 ```bash
MSYS_NO_PATHCONV=1 openssl req -x509 -nodes -days 36500 -newkey rsa:2048 \
-keyout .ssl/key.pem \
-out .ssl/cert.pem \
-subj "/C=JP/ST=Tokyo/L=Tokyo/O=Development/CN=localhost"
```
   
 このコマンドは以下を生成します：
 - `cert.pem`: SSL証明書
 - `key.pem`: 秘密鍵
 - 有効期限: 100年
 - CN（Common Name）: localhost

 このあたりは環境によって違うと思うので、各々調べてください。結果的に.ssl内に自己署名証明書があればOK。
 `cert.pem`と`key.pem`の2つのファイルが存在することを確認してください。

これでバックエンドで`go run main.go`を実行すると、`http://localhost:5173/`でアクセスできる。

### バックエンド
ローカル側で何かしらの手段で実行する。デバッグでも通常の実行でも可。基本的には`go run main.go`です。<br/>
`air`が導入されてるのでHMRを効かせたい方は`go run air`でも可。僕は使わなくなった...<br>

### フロント
コンテナが起動してたら`https://localhost/`でアクセスできるはずです。(裏で`npm run dev`が走ってる)<br/>
HMRが効いているので開発中は特に何も意識しなくてOKのはず。ちなみにバックエンドが実行されてないとデータは描画されません。

### DB
シーダーとかはバックエンド側がよしなにやってくれます。ただ単にコンテナが起動してれば動きます。

### プロキシサーバー
プロキシサーバーを起動するには、事前にSSL証明書の準備が必要です。
ローカルサーバーとの通信のために存在しています。(これがあることでローカルでデバッグが可能になる)<br/>

#### ブラウザでのアクセス
初回アクセス時、ブラウザで自己署名証明書の警告が表示されます。
これは開発環境で自己署名証明書を使用しているためです。
「詳細設定」→「安全でないサイトに進む」（Chromeの場合）などで続行してください。<br/>

### Git hooks (lefthook)
プロジェクトではlefthookを使用してGit hooksを管理しています。
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
Docker上でGolandを動かす方法もあるのですが、やってることがオーバーなのと、コンテナから開くと拡大率が100%から下がらず見づらかったという個人的な理由で採用しませんでした。
