# pact-to-karate

Scripts to take Pact contracts as input and convert them to executable Karate test cases

Note that this is only relevant for provider-side testing, where these test cases can be used to test the provider meets all the relevant consumer contracts. For consumer-side testing, you need to turn the Pact contracts into stub definitions to test the consumer code against a set of provider stubs.

## Assumptions

- your execution environment is Linux
- you have recent version of `jq` installed
- you have a bash shell
- you either have access to the Internet to download `hjson`, or have `hjson` already installed

## To use

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

## To do

Create a Dockerfile that can take a volume of Pact JSONs, turn them into Karate test cases and emit them in the same volume. This is likely to be a good fit for a CI+Docker/Kubernetes environment
