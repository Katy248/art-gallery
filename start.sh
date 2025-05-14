#!/bin/bash

if [[ $1 == "server" ]]; then
    cd /root/art-gallery/server || exit 2
    
    export GOPATH='/root/go'
    export GOENV='/root/.config/go/env'
    go run .
    elif [[ $1 == "client" ]]; then
    cd /root/art-gallery/client || exit 2
    npm run dev
else
    echo "Usage: start.sh [server|client]"
    exit 1
fi
