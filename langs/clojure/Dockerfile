FROM alpine:3.20 AS builder

RUN apk add --no-cache curl

ENV VER=1.4.192

RUN curl -L https://github.com/babashka/babashka/releases/download/v$VER/babashka-$VER-linux-amd64-static.tar.gz | tar xz

FROM codegolf/lang-base

COPY --from=0 /bb /usr/bin/clojure

ENTRYPOINT ["clojure"]

CMD ["--version"]
