# Golang version

This Golang version contains the code necessary to consume a Pact JSON file, and emit the equivalent Karate test cases & stub definitions

## To build

`$ docker build -t pact-karate .`

will create a tiny (~3Mb!) Docker image containing this executable

## To run

`$ cat ../../pacts/sample-pact-extended2.v2.json | docker run -i pact-karate:latest` will run this Docker image against the Pact JSON file, and emit Karate test cases & stub
