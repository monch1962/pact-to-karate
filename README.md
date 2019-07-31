# pact-to-karate

Scripts to take Pact contracts as input and convert them to executable Karate test cases

## Assumptions

- your execution environment is Linux
- you have `jq` installed
- you have a bash shell
- you either have access to the Internet to download `hjson`, or have `hjson` installed

## To use

Assuming your Pact contracts are available in `./pacts`...

Sometimes Pact contracts are presented as invalid JSON due to a bug in the Pact broker - need to fix that...

First install HJSON, a tool to fix broken JSONs

`$ ./get-hjson.sh`

Now use it to fix the broken pacts

`$ fix-broken-json pacts/*.json`

Now convert your Pact contracts (with valid JSON) to Karate test cases

`$ cat pacts/sample-pact-extended.v2.json | ./pact-to-karate.sh`

The output should be an set of executable Karate test scenarios (for testing the provider of a consumer-provider pair), in a single feature file for each Pact contract JSON
