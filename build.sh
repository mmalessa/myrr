#!/bin/bash
OD="$(pwd)"

# Pushes application version into the build information.
RR_VERSION=1.8.2

# Hardcode some values to the core package
LDFLAGS="$LDFLAGS -X github.com/spiral/roadrunner/cmd/rr/cmd.Version=${RR_VERSION}"
LDFLAGS="$LDFLAGS -X github.com/spiral/roadrunner/cmd/rr/cmd.BuildTime=$(date +%FT%T%z)"
# remove debug info from binary as well as string and symbol tables
LDFLAGS="$LDFLAGS -s"

build() {
  echo Packaging "$1" Build
  bdir=roadrunner-${RR_VERSION}-$2-$3
  rm -rf builds/"$bdir" && mkdir -p builds/"$bdir"
  GOOS=$2 GOARCH=$3 ./build.sh

  if [ "$2" == "windows" ]; then
    mv rr builds/"$bdir"/rr.exe
  else
    mv rr builds/"$bdir"
  fi

  cp README.md builds/"$bdir"
  cp CHANGELOG.md builds/"$bdir"
  cp LICENSE builds/"$bdir"
  cd builds

  if [ "$2" == "linux" ]; then
    tar -zcf "$bdir".tar.gz "$bdir"
  else
    zip -r -q "$bdir".zip "$bdir"
  fi

  rm -rf "$bdir"
  cd ..
}

CGO_ENABLED=0 go build -trimpath -ldflags "$LDFLAGS" -o "$OD/myrr" main.go
