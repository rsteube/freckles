FROM golang:alpine

ENV GOPATH /go
RUN ln -s /freckles/cmd/freckles/freckles /usr/local/bin/freckles

RUN apk add --no-cache curl git github-cli

# carapace-bin
RUN curl -L https://github.com/rsteube/carapace-bin/releases/download/v0.12.6/carapace-bin_0.12.6_Linux_x86_64.tar.gz | tar -xvz \
 && mv carapace /usr/local/bin/carapace

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo -e "\
eval (carapace _carapace|slurp)\n\
eval (freckles _carapace|slurp)" \
  > /root/.elvish/rc.elv

RUN echo -e "[credential \"https://github.com\"]\n\
        helper = !gh auth git-credential\n"\
      > /root/.gitconfig
