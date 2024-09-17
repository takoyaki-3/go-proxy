**Caddy Proxyに移行しました。現在は利用していません。**

## Reverse Proxy with Autocert

このリポジトリは、Go言語とAutocertを使用してHTTPS対応のリバースプロキシを実現するシンプルなプロジェクトです。docker-composeを使用して簡単にデプロイできます。

### 特徴

- HTTPS自動化: Autocertを使用して、Let's EncryptからSSL証明書を自動的に取得・更新します。
- 設定ファイルベース: `conf.json` ファイルでリバースプロキシの設定を管理します。
- Dockerによる容易なデプロイ: docker-composeを使って簡単に環境構築・起動ができます。
- ロギング機能: リクエストをロガーサーバーに送信し、アクセスログを記録します。


### ファイルツリー

```
├── docker-compose.yml
└── volume
    ├── go.mod
    ├── go.sum
    └── proxy.go

```

- `docker-compose.yml`: Dockerコンテナの設定ファイルです。
- `volume/`: アプリケーションのソースコードと設定ファイルが格納されています。
    - `volume/go.mod`: Goモジュールの依存関係を管理するファイルです。
    - `volume/go.sum`: Goモジュールのチェックサムを管理するファイルです。
    - `volume/proxy.go`: リバースプロキシのメインロジックを実装したGoのソースコードです。
    - `volume/conf.json`: リバースプロキシの設定ファイルを配置します。（後述）


### 設定ファイル

リバースプロキシの設定は `volume/conf.json` ファイルで行います。このファイルはJSON形式で、以下の構造を持ちます。

```json
[
  {
    "domain": "example.com",
    "host": "backend.example.com",
    "scheme": "https"
  },
  {
    "domain": "api.example.com",
    "host": "api.internal.example.com",
    "scheme": "http"
  }
]
```

- `domain`: リバースプロキシの対象となるドメイン名です。
- `host`: リクエストを転送する先のホスト名です。
- `scheme`: 転送先のスキーム (http または https) です。


### インストール

1. リポジトリをクローンします。
```bash
git clone https://github.com/<your-username>/<your-repository>.git
```

2. `volume/conf.json` ファイルを作成し、リバースプロキシの設定を記述します。

### 実行

1. Dockerコンテナを起動します。
```bash
docker-compose up -d
```

これで、指定したドメインへのHTTPSアクセスが、設定されたバックエンドサーバーに転送されるようになります。


### コマンド実行例

#### Dockerコンテナの起動
```bash
docker-compose up -d
```

#### Dockerコンテナの停止
```bash
docker-compose down
```

#### Dockerコンテナのログの確認
```bash
docker-compose logs -f file_server
```

### 注意点

- Autocertを使用するには、指定したドメインのDNSレコードを適切に設定する必要があります。
- ロガーサーバー (`logger.api.takoyaki3.com`) はデフォルトで設定されています。必要に応じて変更してください。

### その他

このREADMEは既存のREADMEとリポジトリの最新情報に基づいて生成されました。詳細な情報や最新の情報については、リポジトリのソースコードを参照してください。