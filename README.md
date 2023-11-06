# s3-static-api

s3-static-api は、Amazon S3 の静的ファイルを操作するための REST API を提供します。

## 環境変数

このプロジェクトは以下の環境変数を必要とします。これらの環境変数を記載した`.env`ファイルを作成してください。

- `STORAGE_ENDPOINT`: オブジェクトストレージのベースのエンドポイント
- `STORAGE_ACCESS_KEY`: オブジェクトストレージのアクセスキー ID
- `STORAGE_SECRET_KEY`: オブジェクトストレージのシークレットアクセスキー
- `STORAGE_BUCKET`: オブジェクトストレージのバケット名
- `PORT`: API サーバのポート(デフォルトは`8080`)

## Docker Compose を使用した起動

ローカル環境での Docker Compose の起動方法を以下に示します。

```bash
docker-compose up
```

## ファイルのアップロード

アップロードしたいファイルは、`static`ディレクトリに格納してください。これらのファイルは、API を通じて S3 にアップロードされます。

## 使い方

このプロジェクトの基本的な使い方を説明します。

### ビルド

このプロジェクトをビルドするには、`static`フォルダにアップロードするファイルを格納し、`.env`ファイルを準備してから以下のコマンドを実行します。

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api-server/main.go
```

### サーバ起動

```bash
./main
```

API は以下のエンドポイントを提供します:

- `GET /`: 指定したキーのオブジェクトを取得します。キーはクエリパラメータとして渡します。例: `GET /?key=my-object-key`
- `PUT /`: 指定したキーで新しいオブジェクトを作成または更新します。キーはリクエストボディに JSON 形式で含めます。例: `PUT /` with JSON body `{"key": "my-object-key"}`
