FROM docker:17.06
MAINTAINER Monax <support@monax.io>

ENV USER monax
ENV MONAX_PATH /home/$USER/.monax
ENV TARGET monax
ENV INSTALL_BASE /usr/local/bin

RUN addgroup -S $USER && adduser -s /bin/bash -S -D -u 1000 $USER $USER && \
  echo "alias ls='ls --color=never'" > /home/monax/.profile && \
  mkdir $MONAX_PATH && \
  mkdir $MONAX_PATH/apps && \
  mkdir $MONAX_PATH/bundles && \
  mkdir $MONAX_PATH/chains && \
  mkdir $MONAX_PATH/chains/account-types && \
  mkdir $MONAX_PATH/chains/chain-types && \
  mkdir $MONAX_PATH/keys && \
  mkdir $MONAX_PATH/keys/data && \
  mkdir $MONAX_PATH/keys/names && \
  mkdir $MONAX_PATH/services && \
  chown --recursive $USER:$USER $MONAX_PATH && \
  adduser $USER ping

RUN apk add --no-cache \
  bash \
  coreutils \
  curl \
  git \
  nodejs \
  nodejs-npm \
  jq

COPY ./target/"$TARGET".dockerartefact $INSTALL_BASE/$TARGET
COPY ./monax-entrypoint.sh $INSTALL_BASE

RUN chmod +x $INSTALL_BASE/$TARGET && \
  chmod +x $INSTALL_BASE/monax-entrypoint.sh && \
  chgrp 1000 $INSTALL_BASE/$TARGET && \
  unset INSTALL_BASE && \
  unset TARGET

# Finalize
WORKDIR /home/$USER
ENTRYPOINT ["/usr/local/bin/monax-entrypoint.sh"]
