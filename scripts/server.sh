#!/bin/sh

if [ $1 = "start" ]; then
  echo "Starting server .."

  if [ ! -z $2 ]; then
    export PORT="$2"
  fi
  GIN_MODE=release go run ./cmd/port_forwarder/main.go &
  sleep 5
  open "http://localhost:$PORT"
fi

if [ $1 = "stop" ]; then
  echo "Stoping server .."
  kill -9 `cat run.pid`
  rm run.pid
  echo "Server stopped .."
fi

# start
