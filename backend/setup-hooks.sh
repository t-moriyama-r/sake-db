#!/bin/bash

# Lefthook のインストールとセットアップ（バックエンド専用）
# 注意: Lefthookはリポジトリルートにインストールされますが、バックエンドファイルのみをチェックします

echo "Setting up backend static analysis hooks..."

# ルートディレクトリに移動
cd "$(git rev-parse --show-toplevel)" || exit 1

# lefthookがインストールされているか確認
if ! command -v lefthook &> /dev/null; then
    echo "Installing lefthook..."
    
    # OSに応じてインストール方法を選択
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        echo "Installing via Go..."
        go install github.com/evilmartians/lefthook@latest
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        if command -v brew &> /dev/null; then
            brew install lefthook
        else
            echo "Installing via Go..."
            go install github.com/evilmartians/lefthook@latest
        fi
    else
        # その他の環境: Goでインストール
        echo "Installing lefthook via Go..."
        go install github.com/evilmartians/lefthook@latest
    fi
fi

# lefthookのインストール
GOPATH_BIN="$(go env GOPATH)/bin"
if command -v lefthook &> /dev/null; then
    lefthook install
    echo "✓ Lefthook setup completed!"
elif [ -f "$GOPATH_BIN/lefthook" ]; then
    "$GOPATH_BIN/lefthook" install
    echo "✓ Lefthook setup completed!"
else
    echo "✗ Failed to install lefthook"
    exit 1
fi

echo ""
echo "Lefthook is now configured to check backend files only."
echo "To install other Go development tools, run:"
echo "  cd backend && make install-tools"


