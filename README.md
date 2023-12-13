
# 1. 概要

このリポジトリは以下の要件に従い作成されました。

1. `Contentful API` からパンの情報を取得し、DBに保存する CLI
2. 以下の要件を満たすgraphQLサーバ
  - 1で保存したデータを一覧で取得できる
  - id を指定して、1で保存したデータ中の任意のデータを取得できる


# 2. 環境構築

以下のコマンドもしくは、vscodeのdevcontainerでコンテナを起動させる

```bash
cd .devcontainer
docker compose up # or docker-compose up
```


# 3. DBの初期化

以下のコマンドを実行してDBを初期化する

```bash
docker container exec -it panforyou-dev bash
cd /workspace/scripts/
make setup_db
exit
```


# 4. サーバーを起動させる

以下のコマンドを実行。
その後、`http://localhost:8080/`にアクセスして正常に起動していることを確認

```bash
docker container exec -it panforyou-dev bash
cd /workspace/scripts/
make test
```


# 5. CLIの実行

`ENTRY_ID` を指定してコマンドを実行する (`xxx` には適切な値を入力する)

コマンド実行後に`Contentful API` から取得した情報がdatabaseに保存されます

※ 3でサーバ起動時にcliも同時にコンパイルされています

```bash
docker container exec -it panforyou-dev bash
cd /workspace/bin
export CONTENTFUL_ENDPINT=https://cdn.contentful.com
export CONTENTFUL_SPACES=xxxxxxxxxxxx
export CONTENTFUL_ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
./cli -e [ENTRY_ID]
```


# 6. クエリの実行

`http://localhost:8080/` にアクセスして以下のqueryを実行


## 6.1. CLIで保存したデータを一覧で取得

```graphql
query findBreads {
  breads{
    id
    name
    createdAt
  }
}
```


## 6.2. id を指定してCLIで保存したデータ中の任意のデータを取得

```graphql
query findBread {
  bread(id: "4Li6w5uVbJNVXYVxWjWVoZ") {
    id
    name
    createdAt
  }
}
```


# 7. 環境破棄

以下のコマンドもしくは、vscodeのdevcontainerでコンテナを停止させる

```bash
cd .devcontainer
docker compose down # or docker-compose down
```
