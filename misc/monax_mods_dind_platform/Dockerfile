FROM quay.io/monax/monax:0.18.0-dind
MAINTAINER Monax <support@monax.io>

env HELM_VERSION 2.7.2
env KUBERNETES_VERSION 1.8.3

RUN apk add --update --no-cache \
  g++ \
  libc6-compat \
  make \
  nodejs \
  nodejs-npm \
  python

# via: https://kubernetes.io/docs/tasks/tools/install-kubectl/
# RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl && \
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v$KUBERNETES_VERSION/bin/linux/amd64/kubectl && \
  chmod +x ./kubectl && \
  mv ./kubectl /usr/local/bin/kubectl

# via: https://docs.helm.sh/using_helm/#installing-helm
RUN curl -LO https://storage.googleapis.com/kubernetes-helm/helm-v$HELM_VERSION-linux-amd64.tar.gz && \
  tar -xzf helm-v$HELM_VERSION-linux-amd64.tar.gz && \
  mv linux-amd64/helm /usr/local/bin/helm && \
  rm -rf helm-* && \
  rm -rf linux-amd64
