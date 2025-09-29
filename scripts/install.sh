#!/bin/bash
echo "Hi from deno"

NAME="tessa"

if ! command -v go >/dev/null 2>&1; then
    echo "Go is not installed or not in PATH"
    exit 1
fi

echo "building.."
go build -o $NAME ./cmd/main.go
if [ $? -ne 0 ]; then
    echo "Go build failed"
    exit 1
fi

echo "configuring.."
sudo mv $NAME /usr/bin

if [ -n "$ZSH_VERSION" ]; then
    RCFILE="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
    RCFILE="$HOME/.bashrc"
else
    RCFILE="$HOME/.profile"
fi

echo 'alias tessa="/usr/bin/$NAME"' >> $RCFILE

. "$RCFILE"

echo "All done!"