#!/bin/bash

# Lefthook のインストールとセットアップ

echo "Setting up lefthook..."

# lefthookがインストールされているか確認
if ! command -v lefthook &> /dev/null; then
    echo "Installing lefthook..."
    
    # OSに応じてインストール方法を選択
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
        curl -1sLf 'https://dl.cloudsmith.io/public/evilmartians/lefthook/setup.rpm.sh' | sudo -E bash
        sudo yum install lefthook -y || sudo apt-get install lefthook -y
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        # macOS
        brew install lefthook
    else
        # その他の環境: Goでインストール
        echo "Installing lefthook via Go..."
        go install github.com/evilmartians/lefthook@latest
    fi
fi

# lefthookのインストール
lefthook install

echo "✓ Lefthook setup completed!"
echo ""
echo "To install Go development tools, run:"
echo "  cd backend && make install-tools"
