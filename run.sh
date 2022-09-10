#!/bin/bash

URL=http://localhost:3000


if [ "$1" = "stop" ]
then
  docker-compose down
else
  # Pull image if not already present
  docker-compose pull

  # Start server, frontend locally, build from source with layer caching
  docker-compose up -d --build

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
