# s3-static-api

s3-static-api は、Amazon S3 および S3 互換のオブジェクトストレージで静的ファイルを操作するための REST API を提供します。

## 環境変数

このプロジェクトは以下の環境変数を必要とします。これらの環境変数を記載した`.env`ファイルを作成してください。

- `STORAGE_ENDPOINT`: オブジェクトストレージのベースのエンドポイント
- `STORAGE_ACCESS_KEY`: オブジェクトストレージのアクセスキー ID
- `STORAGE_SECRET_KEY`: オブジェクトストレージのシークレットアクセスキー
- `STORAGE_BUCKET`: オブジェクトストレージのバケット名
- `STORAGE_USE_PATH_STYLE`: パススタイル形式のエンドポイントか否か(bool 型。デフォルトは`false`)
  - `true`: パススタイル形式のエンドポイント
  - `false`: バケット名を含むドメインのエンドポイント
- `PORT`: API サーバのポート(デフォルトは`8080`)

## Docker Compose を使用したローカル環境での起動

ローカル環境での Docker Compose の起動方法を以下に示します。
OSS のオブジェクトストレージ `minIO` のコンテナも同時に起動します。

```bash
docker compose build
docker compose up -d
```

- API サーバ: 8080 番ポート
- minIO: 9000 番ポート(admin / adminpass)

## ファイルのアップロード

アップロードしたいファイルは、`internal/storage/static`ディレクトリに格納してください。これらのファイルは、API を通じて オブジェクトストレージ にアップロードされます。

## 使い方

このプロジェクトの基本的な使い方を説明します。

### 準備

- `internal/storage/static`フォルダにアップロードするファイルを格納
- `.env.sample`からコピーした`.env`ファイルに必要な設定値を記入

### ビルド

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api-server/main.go
```

### サーバ起動

```bash
./main
```

環境変数設定ファイル名を`.env`にする場合は以下のように起動してください。
(以下は`.env.dev`とした例)

```bash
ENV_FILE=.env.dev ./main
```

API は以下のエンドポイントを提供します:

- `GET /`: 指定したキーのオブジェクトを取得します。キーはクエリパラメータとして渡します。
  - 例: `GET /?key=my-object-key`
- `PUT /`: 指定したキーで新しいオブジェクトを作成または更新します。キーおよびオブジェクト ID へのタイムスタンプ付与要否はリクエストボディに JSON 形式で含めます。
  - 例: `PUT /` with JSON body `{"key": "my-object-key","add_timestamp":true}`
