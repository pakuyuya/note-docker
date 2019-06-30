* Docker version 1.12.2, build bb80604

# docker commit

コンテナからDockerイメージを作成する。

```bash
root@yu:~# docker help commit

Usage:  docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]

Create a new image from a container's changes

Options:
  -a, --author string    Author (e.g., "John Hannibal Smith <hannibal@a-team.com>")
  -c, --change value     Apply Dockerfile instruction to the created image (default [])
      --help             Print usage
  -m, --message string   Commit message
  -p, --pause            Pause container during commit (default true)
```

# 例

```
docker commit webapp webapp:dev
```

# 確認

作成したイメージは `docker images`で確認できる。

```
docker images
```