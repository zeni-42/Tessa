#!/bin/bash
NAME="tessa"

echo "Uninstalling $NAME..."

if [ -f "/usr/bin/$NAME" ]; then
    sudo rm /usr/bin/$NAME
    echo "Removed /usr/bin/$NAME"
else
    echo "$NAME binary not found in /usr/bin"
fi

if [ -n "$ZSH_VERSION" ]; then
    RCFILE="$HOME/.zshrc"
elif [ -n "$BASH_VERSION" ]; then
    RCFILE="$HOME/.bashrc"
else
    RCFILE="$HOME/.profile"
fi

if grep -q "alias $NAME=" "$RCFILE"; then
    sed -i.bak "/alias $NAME=/d" "$RCFILE"
    echo "Removed alias from $RCFILE (backup saved as $RCFILE.bak)"
else
    echo "No alias found in $RCFILE"
fi

. "$RCFILE"

echo "$NAME has been uninstalled!"
