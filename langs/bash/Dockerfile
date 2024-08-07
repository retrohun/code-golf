FROM alpine:3.20 AS builder

RUN apk add --no-cache build-base curl patch

RUN curl https://ftp.gnu.org/gnu/bash/bash-5.2.tar.gz | tar xz && cd bash-5.2 \
 && curl https://ftp.gnu.org/gnu/bash/bash-5.2-patches/bash52-0[01-32] | patch -p0

RUN cd bash-*                  \
 && ./configure                \
    --disable-command-timing   \
    --disable-debugger         \
    --disable-directory-stack  \
    --disable-help-builtin     \
    --disable-job-control      \
    --disable-net-redirections \
    --disable-progcomp         \
    --disable-select           \
    --enable-static-link       \
    --without-bash-malloc      \
 && make                       \
 && strip bash

FROM codegolf/lang-base-no-proc

COPY --from=0 /bash-*/bash /usr/bin/

ENTRYPOINT ["bash"]

CMD ["-c", "echo ${BASH_VERSION%\\([0-9]\\)-release}"]
