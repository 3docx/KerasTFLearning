#!/bin/bash

cd `dirname $0`
ROOT=$PWD/_dist

if [ -d "$ROOT" ]; then rm -rf $ROOT; fi

mkdir $ROOT

function build() {
  export GOOS=$1
  export GOARCH=$2

  echo build $GOOS $GOARCH
  OUT=$ROOT/${GOOS}_${GOARCH}
  mkdir $OUT
  cd $OUT
  go build github.com/hjson/hjson-go/hjson-cli
  if [[ $3 == "zip" ]]; then
    mv $OUT/hjson-cli.exe $OUT/hjson.exe
    zip -j ${OUT}.zip $OUT/*
  else
    mv $OUT/hjson-cli $OUT/hjson
    tar -czf ${OUT}.tar.gz -C $OUT .
  fi

}

# not all targets can be