#!/bin/bash

PIDS=$(ps aux | grep 'redis-server' | grep -E '7000|7001|7002|7003|7004|7005' | awk '{print $2}')

if [ -n "$PIDS" ]; then
  echo "Stopping: $PIDS"
  kill $PIDS
else
  echo "No Redis processes found"
fi
