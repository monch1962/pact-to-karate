#!/bin/bash
# First get hjson
GET=https://github.com/hjson/hjson-go/releases/download/v3.0.0/linux_amd64.tar.gz
curl -sSL $GET | sudo tar -xz -C /usr/local/bin
