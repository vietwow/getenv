#!/bin/bash

docker login && docker build -t vietwow/getenv . && docker push vietwow/getenv
