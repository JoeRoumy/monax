FROM quay.io/monax/monax:0.18.1-dind
MAINTAINER Monax <support@monax.io>

env HELM_VERSION 2.8.2
env KUBERNETES_VERSION 1.8.3
env CLOUD_SDK_VERSION 193.0.0

RUN apk add --update --no-cache \
  g++ \
  libc6-compat \
  make \
  nodejs \
  nodejs-npm \
  python \
  py-crcmod \
  openssh-client

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

# via: https://github.com/GoogleCloudPlatform/cloud-sdk-docker/blob/master/alpine/Dockerfile
RUN curl -LO https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz && \
  tar -xzf google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz && \
  mv google-cloud-sdk /var && \
  rm google-cloud-sdk-${CLOUD_SDK_VERSION}-linux-x86_64.tar.gz && \
  ln -s /var/google-cloud-sdk/lib /var/google-cloud-sdk/lib64 && \
  /var/google-cloud-sdk/bin/gcloud config set core/disable_usage_reporting true && \
  /var/google-cloud-sdk/bin/gcloud config set component_manager/disable_update_check true
ENV PATH "$PATH:/var/google-cloud-sdk/bin"
