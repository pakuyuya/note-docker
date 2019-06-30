チートシート

* Docker version 1.12.2, build bb80604

## イメージ系

```bash
# 【取得】イメージをリポジトリから取得
docker pull <イメージ名>[:タグ]

# 【一覧】イメージの一覧表示
docker images

# 【削除】メージを消す
docker rmi <イメージ名またはイメージD＞
```

## コンテナ起動・停止系

```bash
# 【作成】イメージをコンテナ化して走らせる
docker run --name <コンテナ名> -d <イメージIDまたはイメージ名>

# 【実行中一覧】走ってるコンテナの一覧
docker ps

# 【一覧】走ってる or 止まってるコンテナの一覧
docker ps -a

# 【停止】コンテナを止める
docker stop <コンテナ名またはコンテナID>

# 【停止】コンテナを再開する
docker start <コンテナ名またはコンテナID>

# 【削除】コンテナを消す（ボリュームもついでに掃除）
docker rm -v <コンテナ名またはコンテナID>
```

## コンテナ操作系

```bash
# 【ターミナル】走ってるコンテナにTTYで入る
docker exec -it <コンテナ名またはコンテナID> bash

# 【bash実行】走ってるコンテナに任意のコマンド
docker exec -d <コンテナ名またはコンテナID> <コマンド>

# 【ログ】走ってるコンテナの、Dockerファイル定義CMDで実行された標準出力ログを見る
docker logs <コンテナ名またはコンテナID>

# 【ファイル配置】ファイルをローカルから転送する
docker cp <ローカルパス> <コンテナID>:<コンテナ上の絶対パス>

# 【ファイル配置】ファイルをローカルへ転送する
docker cp <ローカルパス> <コンテナID>:<コンテナ上の絶対パス>
```

## Dockefile系

```bash
# Dockerfileをビルドする
docker build -t <イメージ名> <Dockerファイルが配置されているパス>
```
