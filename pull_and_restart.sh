#!/bin/bash

cd ~/Area/

# verifie si ya des changements
git fetch

# compare les branches
LOCAL=$(git rev-parse HEAD)
REMOTE=$(git rev-parse @{u})

# si ya un changement il pull et relance
if [ $LOCAL != $REMOTE ]; then
    git pull
    git stash
    git stash drop
    # Restart Docker containers
    docker-compose down
    docker-compose up -d --build
    docker-compose up -d
    docker-compose up -d
fi
