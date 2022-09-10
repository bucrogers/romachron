#!/bin/bash

URL=http://localhost:3000


if [ "$1" = "stop" ]
then
  docker-compose down
else
  # Start server, frontend locally
  docker-compose up -d

  # Launch browser on frontend
  if which open > /dev/null
  then
    open ${URL}
  elif which xdg-open > /dev/null
    xdg-open ${URL}
  elif which gnome-open > /dev/null
  then
    gnome-open $URL}
  else
    echo "Error: Can't find browser"    
  fi
fi
