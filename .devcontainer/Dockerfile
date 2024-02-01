ARG GO_VERSION=1.20.5
FROM golang:${GO_VERSION} AS build

ARG NODE_VERSION="v18.12.1"
RUN su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"