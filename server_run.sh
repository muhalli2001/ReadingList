#!/bin/bash

kill -9 $(lsof -t -i tcp:4000)

go run ./cmd/api


