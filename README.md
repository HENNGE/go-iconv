# go-iconv

[![Go](https://github.com/HENNGE/go-iconv/actions/workflows/go.yml/badge.svg)](https://github.com/HENNGE/go-iconv/actions/workflows/go.yml)

`go-iconv` is a GNU [libiconv](https://www.gnu.org/software/libiconv/) wrapper for Go.

## Run the tests

**Linux (via Docker)**:
```console
docker build -f Dockerfile.test -t local/go-iconv-test .
docker run -it -v .:/usr/src/app local/go-iconv-test bash
go test -v
```

**Linux (Native)**:
```console
# consult with Dockerfile.test on how to install libiconv on your system
go test -v
```

**Darwin (macOS)**:
```console
# ensure you have installed libiconv via homebrew
brew install libiconv

# ensure you have CGO variables to use brew's libiconv
export CGO_CFLAGS="${CGO_CFLAGS} -I${HOMEBREW_PREFIX}/opt/libiconv/include"
export CGO_LDFLAGS="${CGO_LDFLAGS} -L${HOMEBREW_PREFIX}/opt/libiconv/lib"

go test -v
```

## Acknowledgement

- `HENNGE/go-iconv` is based on https://github.com/bu-/go-iconv
  - which is based on https://github.com/mattn/go-iconv
  - which is based on https://github.com/sloonz/go-iconv
  - which is based on https://github.com/oibore/go-iconv (original)

**Caveats**: The original author doesn't seem to put LICENSE notice. We also just treat it as-is.
