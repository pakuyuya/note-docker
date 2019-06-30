FROM mysql:5.5

# MySQLイメージ用の環境変数設定
ENV MYSQL_ROOT_USER root
ENV MYSQL_ROOT_PASSWORD root
ENV MYSQL_DATABASE sample_db
ENV MYSQL_USER test
ENV MYSQL_PASSWORD password

EXPORT '3306:3306'

# MySQLイメージの初期化SQL用ディレクトリに配置
COPY ./init.d/init.sql /docker-entrypoint-initdb.d

# Docker 1.12だと、以下がうまくかなかった。（/init.d:/docker-entrypoint-initdb.dというボリュームが作成されてしまう。）
#VOLUME ./init.d:/docker-entrypoint-initdb.d

# docker run の -v フラグで上記パスを設定すると、意図通り./init.dが/docker-entrypoint-initdb.dにマウントされた。
# 仕様変更か何かでつぶされた？
