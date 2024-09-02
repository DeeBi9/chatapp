#!/bin/bash

# URL to send requests to
url="localhost:8080/conn"

# Send 1000 requests to the URL
for i in {1..1000}
do
   curl $url
done