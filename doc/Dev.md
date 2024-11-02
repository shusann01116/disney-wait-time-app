# TDR Portal

## 技術

使用する技術

アジリティを考慮して、今回は Next.js ですべてを完結させる

### 共通

- dependabot
  - 依存関係の更新を自動でマージしてくれる

### フロントエンド

- Next.js: https://nextjs.org/
- shadcn/ui: https://ui.shadcn.com/
- Tailwind CSS: https://tailwindcss.com/
- TypeScript: https://www.typescriptlang.org/
- Storybook: https://storybook.js.org/
- msw: https://mswjs.io/
- Vitest: https://vitest.dev/
  - ブラウザモードでテストできる
- lucide: https://lucide.dev/
  - アイコン
- eslint: https://eslint.org/
- prettier: https://prettier.io/
- prisma: https://www.prisma.io/
  - ORM マッパー
  - データベースのスキーマをコードで管理できる

### インフラ

- vercel: https://vercel.com/

#### CI/CD

- GitHub Actions: https://github.com/features/actions
  - eslint
  - tsc
  - vitest

## ドメイン

このアプリで達成したいこと

### ユーザージャーニー

1. 各施設の待ち時間を確認できる
2. ショー・パレードのスケジュールを簡単に確認できる
3. フードメニューと価格、場所を確認できる

## TODO

### data

- [ ] Prisma のスキーマ定義
- [ ] データベースのマイグレーション

### ui

- [ ] Figma デザイン起こし
- [ ] Storybook セットアップ
- [ ] vitest セットアップ
