#!/bin/sh
set -eux

LIBICONV_VERSION=1.18
LIBICONV_TMPDIR=$(mktemp -d)

cd "${LIBICONV_TMPDIR}"

wget https://ftp.gnu.org/pub/gnu/libiconv/libiconv-${LIBICONV_VERSION}.tar.gz
tar xf libiconv-${LIBICONV_VERSION}.tar.gz
cd libiconv-${LIBICONV_VERSION}

./configure --prefix=/usr/local --enable-extra-encodings
make
make install
ldconfig
