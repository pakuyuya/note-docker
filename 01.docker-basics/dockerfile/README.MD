* Docker version 1.12.2, build bb80604

## Dockerfileとは

Docker imageの作成を自動化するスクリプトファイル。

## サンプル

```
# ベースにするイメージ（自動的にpullする）
FROM ubuntu:latest

# 作成者
MAINTAINER it's me

# RUN ... コマンドを実行する。
RUN apt-get update
RUN apt-get install apache2
RUN apt-get install php5
RUN apt-get install libapache2-mod-php5

RUN service apache2 restart

# ポートの使用
EXPOSE 22 80

# CMD ... containerをrunするときに実行するコマンド
CMD "bash"
```


## 日本語リファレンス

非常に親切な日本語リファレンスがある。

http://docs.docker.jp/engine/reference/builder.html
