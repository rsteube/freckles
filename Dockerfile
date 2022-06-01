FROM golang:alpine

ENV GOPATH /go
RUN ln -s /freckles-bin/cmd/freckles/freckles /usr/local/bin/freckles

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo "\
eval (freckles _carapace|slurp)" \
  > /root/.elvish/rc.elv
