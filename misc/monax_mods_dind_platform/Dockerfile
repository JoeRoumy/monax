FROM quay.io/monax/monax:0.18.0-dind
MAINTAINER Monax <support@monax.io>

env HELM_VERSION 2.8.2
env KUBERNETES_VERSION 1.8.3

RUN apk add --update --no-cache \
  g++ \
  libc6-compat \
  make \
  nodejs \
  nodejs-npm \
  python \
  py-crcmod

# via: https://kubernetes.io/docs/tasks/tools/install-kubectl/
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v$KUBERNETES_VERSION/bin/linux/amd64/kubectl && \
  chmod +x ./kubectl && \
  mv ./kubectl /usr/local/bin/kubectl

# via: https://docs.helm.sh/using_helm/#installing-helm
RUN curl -LO https://storage.googleapis.com/kubernetes-helm/helm-v$HELM_VERSION-linux-amd64.tar.gz && \
  tar -xzf helm-v$HELM_VERSION-linux-amd64.tar.gz && \
  mv linux-amd64/helm /usr/local/bin/helm && \
  rm -rf helm-* && \
  rm -rf linux-amd64

# via: https://cloud.google.com/storage/docs/gsutil_install
RUN curl -LO https://storage.googleapis.com/pub/gsutil.tar.gz && \
  tar -xzf gsutil.tar.gz -C /var && \
  rm -rf gsutil.tar.gz
ENV PATH "$PATH:/var/gsutil"
