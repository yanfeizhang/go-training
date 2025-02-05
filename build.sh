#!/bin/bash


go build -gcflags=all='-l' -o output/bin/hook 5-performance/hook/main.go