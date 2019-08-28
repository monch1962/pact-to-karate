[![Build Status](https://dev.azure.com/monch1962/monch1962/_apis/build/status/monch1962.pact-to-karate?branchName=master)](https://dev.azure.com/monch1962/monch1962/_build/latest?definitionId=7&branchName=master)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=alert_status)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=sqale_index)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=code_smells)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=coverage)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=ncloc)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=security_rating)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=pact-to-karate)

[![Actions Status](https://github.com/monch1962/pact-to-karate/workflows/Go/badge.svg)](https://github.com/monch1962/pact-to-karate/actions)

# pact-to-karate

Code to take Pact contracts as input and convert them to executable Karate (consumer-side) test cases & (provider-side) stubs.

There are 2 versions of this code: one to run as a Golang executable (optionally within a Docker container), and another to run in a bash+jq environment

## Golang version

### Assumptions
Either you have a Go build environment to compile the executable, or a Docker environment which you can use to build a Docker image containing the Go executable

### To use
Check https://github.com/monch1962/pact-to-karate/blob/master/golang/src/README.md

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

`$ cat pacts/sample-pact-extended.v2.json | ./pact-to-karate-tests.sh`

The output should be an set of executable Karate test scenarios (for testing the provider of a consumer-provider pair), in a single feature file for each Pact contract JSON
