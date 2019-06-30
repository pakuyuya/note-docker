* Docker version 1.12.2, build bb80604

# 起動時にbashでつなぐ

なお、このときTTYから`exit`するとコンテナが終了する。

```
docker run -it imagename /bin/bash
```

# 後からdockerにつなぐ

## 汎用

`docker run`とは違い、こちらのTTYは`logout`しても消えないしlogoutしないとプロセスが残る。

```
docker exec -it コンテナID /bin/bash
```

## Dockerfile記述の CMDが bash のコンテナのみ

```
docker attach コンテナID
```
