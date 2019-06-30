* Docker version 1.12.2, build bb80604

# docker help pull

Docker Imageをリポジトリからダウンロードする。

```
Usage:  docker pull [OPTIONS] NAME[:TAG|@DIGEST]

Pull an image or a repository from a registry

Options:
  -a, --all-tags                Download all tagged images in the repository
      --disable-content-trust   Skip image verification (default true)
      --help                    Print Usage
```

リポジトリ参照先はデフォルトでDocker Hub.リンク先から検索も可。

https://hub.docker.com/

# 任意のリポジトリからダウンロード

NAMEの前に`host:port/`を付与すると任意のリポジトリにアクセスできる。

```
docker pull 10.1.20.35:5000/webapp:dev
```

# 保存先

Docker imageの保存先は`/var/lib/docker`配下のどこか。

| パス | 内容 |
|-----|-----|
| /var/lib/docker/image/aufs/repositories.json | イメージと対応するハッシュ値の対応json |
| /var/lib/docker/aufs | イメージ本体。 |

