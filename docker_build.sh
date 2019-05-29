#!/bin/bash

docker login && docker build -t vietwow/getevn . && docker push vietwow/getevn
