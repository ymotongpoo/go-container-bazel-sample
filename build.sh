#!/bin/bash

statik -f -src=public
go build -o server
./server
