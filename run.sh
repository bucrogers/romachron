#!/bin/bash

URL=http://localhost:3000


if [ $1 = "stop" ]
then
  docker-compose down
else
  # Start server, frontend locally
  docker-compose up -d

  # Launch browser on frontend
  if which xdg-open > /dev/null
  then
    xdg-open ${URL}
  elif which gnome-open > /dev/null
  then
    gnome-open $URL}
  fi
fi
