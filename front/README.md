# front

This template should help get you started developing with Vue 3 in Vite.

## ⚠️ 重要：自動生成ファイルについて

**以下のファイルは自動生成されるため、直接編集しないでください：**

- `src/graphQL/auto-generated.ts` - GraphQLスキーマから自動生成されるTypeScript型定義ファイル

これらのファイルを変更する必要がある場合は：

1. 元のGraphQLスキーマファイル（`../backend/graph/schema/*.graphqls`）またはGraphQLクエリ/ミューテーションファイル（`src/graphQL/`のサブディレクトリ内の`.ts`ファイル、例：`Auth/`, `User/`, `Liquor/`など）を修正してください
   - **注意：`auto-generated.ts`自体は修正しないでください**
2. 以下のコマンドを実行して、自動生成ファイルを再生成してください：
   ```sh
   npm run codegen
   ```

**決して自動生成ファイルを直接編集しないでください。変更は次回のコード生成時に上書きされます。**

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

## Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Run Unit Tests with [Vitest](https://vitest.dev/)

```sh
npm run test:unit
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```
