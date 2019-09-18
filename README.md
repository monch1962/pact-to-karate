[![Build Status](https://dev.azure.com/monch1962/monch1962/_apis/build/status/monch1962.pact-to-karate?branchName=master)](https://dev.azure.com/monch1962/monch1962/_build/latest?definitionId=7&branchName=master)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=alert_status)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=sqale_index)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=code_smells)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=coverage)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=ncloc)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=security_rating)](https://sonarcloud.io/dashboard?id=pact-to-karate)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=pact-to-karate&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=pact-to-karate)

# pact-to-karate

Code to take Pact contracts as input and convert them to executable (consumer-side) test cases & (provider-side) stubs.

### Assumptions

Either you have a Go build environment to compile the executable, or a Docker environment which you can use to build a Docker image containing the Go executable

### Templates

Test case & stub generation is delivered via Handlebars templates; the various template files are stored under `./templates`. The intention is that this collection of templates will grow and evolve over time to cover a broader range of tools (e.g. Mountebank stubs, Hoverfly stubs, Node/mocha/chai test cases, ...) and you'll be able to pick and choose which languages/tools/frameworks you want to use

To use one of these templates, you need to specify it via the environment variable TEMPLATE when you run either the Go code or the Docker container

To generate e.g. a set of Karate tests, you could run e.g. `$ cat sample-pacts/sample-pact-extended2.v2.json | TEMPLATE=karate-tests go run main.go`

To generate a Karate stub, you could run e.g. `$ cat sample-pacts/sample-pact-extended2.v2.json | TEMPLATE=karate-stub go run main.go`

### To use

To run interpreted version locally
`$ cat sample-pacts/sample-pact-extended2.v2.json | TEMPLATE=karate-stub go run main.go`

To run compiled version locally
`$ go build main.go`
`$ cat sample-pacts/sample-pact-extended2.v2.json | TEMPLATE=karate-stub ./main`

To run inside Docker container
`$ docker build -t pact-karate .`
`$ cat sample-pacts/sample-pact-extended2.v2.json | docker run -e TEMPLATE=karate-stub -i pact-karate:latest`
