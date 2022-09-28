#!/bin/bash

# WSL note: For launch-browser behavior via xdg-open, install https://github.com/cpbotha/xdg-open-wsl

URL=http://localhost:3000

set -euo pipefail

if [ "${1:-}" = "stop" ]
then
  docker-compose down
else
  if [ "${1:-}" != "--build" ]
  then
    # Pull image if not already present
    set +e # allow offline failure of pull
    docker-compose pull
    set -e
  fi

  # Start server, frontend locally, build from source with layer caching
  export DOCKER_TZ=$(ls -la /etc/localtime | sed -r  "s/.*\/([^\/]+)\/([^\/]+)$/\1\/\2/")
  echo "DOCKER_TZ: ${DOCKER_TZ}"
  docker-compose up -d ${1:-}

  # Launch browser on frontend
  if which open > /dev/null
  then
    open ${URL}
  elif which xdg-open > /dev/null
  then
    xdg-open ${URL}
  elif which gnome-open > /dev/null
  then
    gnome-open $URL}
  else
    echo "Error: Can't find browser"    
  fi
fi
