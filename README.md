
# 1. 環境構築

以下のコマンドもしくは、vscodeのdevcontainerでコンテナを起動させる

```bash
cd .devcontainer
docker compose up # or docker-compose up
```


# 2. DBの初期化

以下のコマンドを実行してDBを初期化する

```bash
docker container exec -it panforyou-dev bash
cd /workspace/scripts/
make setup_db
exit
```


# 3. サーバーを起動させる

以下のコマンドを実行。
その後、`http://localhost:8080/`にアクセスして正常に起動していることを確認

```bash
docker container exec -it panforyou-dev bash
cd /workspace/scripts/
make test
```


# 4. CLIの実行

`ENTRY_ID` を指定してコマンドを実行する (`xxx` には適切な値を入力する)

※ 3でサーバ起動時にcliも同時にコンパイルされています

```bash
docker container exec -it panforyou-dev bash
cd /workspace/bin
export CONTENTFUL_ENDPINT=https://cdn.contentful.com
export CONTENTFUL_SPACES=xxxxxxxxxxxx
export CONTENTFUL_ACCESS_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
./cli -e [ENTRY_ID]
```


# 5. クエリの実行

`http://localhost:8080/` にアクセスして以下のqueryを実行


## 5.1. CLIで保存したデータを一覧で取得

```graphql
query findBreads {
  breads{
    id
    name
    createdAt
  }
}
```


## 5.2. d を指定してCLIで保存したデータ中の任意のデータを取得

```graphql
query findBread {
  bread(id: "4Li6w5uVbJNVXYVxWjWVoZ") {
    id
    name
    createdAt
  }
}
```
