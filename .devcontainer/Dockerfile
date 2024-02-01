ARG GO_VERSION=1.20.5
FROM golang:${GO_VERSION} AS build

WORKDIR /workspace
RUN git clone https://github.com/gonagutor/dnd-frontend

RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
ARG NODE_VERSION="v18.12.1"
RUN umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1