FROM fluent/fluentd:v1.5-1

USER root

RUN  apk add --no-cache --update --virtual .build-deps \
        sudo build-base ruby-dev \
    && sudo gem install fluent-plugin-multiprocess \
    && sudo gem install fluent-plugin-multi-format-parser \
    && sudo gem install fluent-plugin-elasticsearch \
    && sudo gem install fluent-plugin-file-alternative \
    && apk del .build-deps \
    && rm -rf /home/fluent/.gem/ruby/2.5.0/cache/*.gem
