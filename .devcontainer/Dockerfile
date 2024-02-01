ARG GO_VERSION=1.20.5
FROM golang:${GO_VERSION} AS build

WORKDIR /workspace
RUN git clone https://github.com/gonagutor/dnd-frontend

ARG NODE_VERSION="v18.12.1"

RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash
RUN export NVM_DIR="$([ -z "${XDG_CONFIG_HOME-}" ] && printf %s "${HOME}/.nvm" || printf %s "${XDG_CONFIG_HOME}/nvm")"
RUN nvm install ${NODE_VERSION} 2>&1
