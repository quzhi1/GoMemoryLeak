#!/bin/bash

for CALL_I in {1..10000}
do
  curl --location --request GET 'http://localhost:8080/append-slice'
  echo "Calling /append-slice $CALL_I/10000"
done
