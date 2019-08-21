# pact-to-karate

Code to take Pact contracts as input and convert them to executable Karate (consumer-side) test cases & (provider-side) stubs
There are 2 versions of this code: one to run as a Golang executable (optionally within a Docker container), and another to run in a bash+jq environment

## Golang version

### Assumptions
Either you have a Go build environment to compile the executable, or a Docker environment which you can use to build a Docker image containing the Go executable

## bash+jq version

### Assumptions

- your execution environment is Linux
- you have recent version of `jq` installed
- you have a bash shell
- you either have access to the Internet to download `hjson`, or have `hjson` already installed

### To use

Assuming your Pact contracts are available in `./pacts`...

---

Sometimes Pact contracts are presented as invalid JSON due to a bug in the Pact broker - need to fix that...

First install HJSON, a tool to fix broken JSONs

`$ ./get-hjson.sh`

Now use it to fix the broken pacts

`$ fix-broken-json pacts/*.json`

---

Now convert your Pact contracts (with valid JSON) to Karate test cases

`$ cat pacts/sample-pact-extended.v2.json | ./pact-to-karate.sh`

The output should be an set of executable Karate test scenarios (for testing the provider of a consumer-provider pair), in a single feature file for each Pact contract JSON
