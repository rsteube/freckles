FROM golang

ENV GOPATH /go
RUN ln -s /dotfiles-bin/cmd/dotfiles/dotfiles /usr/local/bin/dotfiles

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo "\
eval (dotfiles _carapace|slurp)" \
  > /root/.elvish/rc.elv
