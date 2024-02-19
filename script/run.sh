#!/bin/sh
if [ -f "./signaling" ]; then
  rm ./signaling  
fi

if [ -e "./log" ]; then
  rm -f ./log/*
fi

bash ./script/build.sh && ./signaling