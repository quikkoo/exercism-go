#!/bin/sh

set -e

MODULES="      \
  anagram      \
  bob          \
  etl          \
  hello-world  \
  word-count   \
"

export PATH=$PATH:$HOME/bin

task_check() {
  MODULE=$1

  go vet -v $MODULE
  golint $MODULE/src/$MODULE
}

task_test() {
  MODULE=$1

  task_check $MODULE

  mkdir -p coverage
  go test -v -cover -covermode=count -coverprofile=$MODULE/coverage/coverage.cov $MODULE
  go tool cover -func=$MODULE/coverage/coverage.cov
  go tool cover -html=$MODULE/coverage/coverage.cov -o $MODULE/coverage/coverage.html
}

task_all() {
  MODULE=$1

  task_test $MODULE
}

task_clean() {
  MODULE=$1

  rm -rf $MODULE/coverage
}

if [ -z $1 ]
then
  TASK=all
else
  TASK=$1
fi

export GOPATH=$HOME
go get -u github.com/golang/lint/golint

for MODULE in $MODULES
do
  export GOPATH=$HOME:"$(pwd)/$MODULE"
  case $TASK in
    check)
      task_check $MODULE
      ;;
    test)
      task_test $MODULE
      ;;
    all)
      task_all $MODULE
      ;;
    clean)
      task_clean $MODULE
      ;;
  esac
done
