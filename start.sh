#!/bin/bash

# Run Go server in the background
go run main.go > go_output.log 2>&1 &

# Run Deno server in the background
cd web/static/js && deno task start --allow-net --allow-read &

# Wait for user input to terminate
wait
