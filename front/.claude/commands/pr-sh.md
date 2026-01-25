# Command: /pr-sh

# Purpose: AI-powered Pull Request Draft Creation

> **使い分けのガイド:**
> - `/pr`: Claude AIが直接gitコマンドとgh CLIを実行してPRを作成します。各ステップでユーザーの確認を求めるため、より安全ですが手動操作が必要です。
> - `/pr-sh`: シェルスクリプト（ai-pr-draft-ja.sh）を自動実行してPRを作成します。完全に自動化されていますが、ユーザー確認なしに実行されます。
>
> 初めての利用や重要な変更の場合は `/pr` を、日常的な作業や信頼できる環境では `/pr-sh` を推奨します。

## Command Execution

- **Execute**: immediate
- **Purpose**: "Create an AI-generated draft PR for the current branch using Japanese prompts"
- **Legend**: Generated based on symbols used in command

## Examples

`/pr-sh`

> Executes `./ai-pr-draft-ja.sh` to create a new draft pull request with Japanese AI prompts.

`/ai-pr-draft --debug`

> Executes `DEBUG=1 ./ai-pr-draft-ja.sh` to run the script in debug mode.

## Core Operation

This command's only function is to execute the `ai-pr-draft-ja.sh` script.

All logic for PR creation, AI communication, and git operations is contained within the script itself. Please refer to the `ai-pr-draft-ja.sh` file for implementation details.

This Japanese version uses Japanese prompts for Claude AI and searches for Japanese PR templates (`PULL_REQUEST_TEMPLATE_JA.md`).

## Flags

### `--debug`

- **Purpose**: Enables debug mode for the `ai-pr-draft-ja.sh` script.
- **Action**: Sets the `DEBUG=1` environment variable before executing the script.