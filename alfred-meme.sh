#! /bin/bash
# http://redsymbol.net/articles/unofficial-bash-strict-mode/

GOBIN=./bin/$(basename $0 .sh)
xattr "$GOBIN" | grep 'com.apple.quarantine' && {
    osascript -e 'display dialog "Alfred: Do you want to trust this version of paste ?" buttons {"No", "Yes"}' | grep 'returned:Yes' && xattr -d com.apple.quarantine "$GOBIN"
}

$GOBIN $*

exit 0


exit 0
