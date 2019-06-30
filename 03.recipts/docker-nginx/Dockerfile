FROM nginx:latest

RUN apt-get update
RUN apt-get install -y supervisor

# コンテンツ配置
RUN mkdir -p /var/www/html
ADD html/index.html /var/www/html/index.html

# nginx.confを上書き
ADD conf/nginx.conf /etc/nginx/nginx.conf

#nginx用のvirtualhost設定追加
ADD conf/site-http.conf /etc/nginx/conf.d/default.conf
ADD conf/site-ssl.conf /etc/nginx/conf.d/default-ssl.conf

# 証明書ファイル追加
RUN mkdir -p /usr/local/nginx/conf/cert
ADD cert/cert.pem /usr/local/nginx/conf/cert.pem
ADD cert/cert.key /usr/local/nginx/conf/cert.key

# supevisordファイル追加
ADD conf/supervisord.conf /etc/supervisord/supervisord.conf

EXPOSE 80 443

#開始スクリプト配置
ADD scripts/start.sh /start.sh

CMD [ "/start.sh" ]
