#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

INPUT=$(echo $(cat))

CONSUMER=$(echo $INPUT |
    jq '.consumer.name')

PROVIDER=$(echo $INPUT |
    jq '.provider.name')

NUM_INTERACTIONS=$(echo $INPUT |
    jq '.interactions | length')

TIMESTAMP=$(date +%Y-%m-%dT%H:%M:%S%z)

for INTERACTION in $(seq $NUM_INTERACTIONS)
do
    #echo $INTERACTION
    I="$(($INTERACTION - 1))"
    DESCRIPTION=$(echo $INPUT | 
        jq --raw-output --argjson i "$I" '.interactions[$i].description') 

    URL=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.path') 

    METHOD=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.method | ascii_downcase') 

    REQUEST_BODY=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.body')

    REQUEST_HEADERS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.headers')

    STATUS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].response.status')

    RESPONSE_BODY=$(echo $INPUT |
        jq --argjson i "$I" '.interactions[$i].response.body')

    RESPONSE_HEADERS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].response.headers')
done

echo $REQUEST_HEADERS
echo $REQUEST_BODY
echo $RESPONSE_HEADERS
echo $RESPONSE_BODY
#exit
DESTINATION=$(echo $URL|sed -e 's/:.*//g')
SCHEME=$(echo $URL | sed -e 's/http.\{0,1\}:\/\///g' | sed -e 's/\/.*//g')
echo "{}" |
    jq --arg TIMESTAMP $TIMESTAMP '.meta.timeExported = ($TIMESTAMP)' |
    jq '.data.pairs[0].request.path[0].matcher = "exact"' |
    jq --arg PATH1 $URL '.data.pairs[0].request.path[0].value = ($PATH1)' |
    jq '.data.pairs[0].request.method[0].matcher = "exact"' |
    jq --arg VERB $METHOD '.data.pairs[0].request.method[0].value = ($VERB | ascii_upcase)' |
    jq '.data.pairs[0].request.destination[0].matcher = "glob"' |
    jq '.data.pairs[0].request.destination[0].value = ".*"' |
    jq '.data.pairs[0].request.scheme[0].matcher = "glob"' |
    jq '.data.pairs[0].request.scheme[0].value = "http.*"' |
    jq --arg HTTP_CODE $STATUS '.data.pairs[0].response.status = ($HTTP_CODE|tonumber)' |
    jq '.data.pairs[0].request.body.matcher = "exact"' |
    #jq --arg REQUEST_BODY $REQUEST_BODY '.data.pairs[0].request.body.value = ($REQUEST_BODY)' |
    #jq --arg REQUEST_HEADERS $REQUEST_HEADERS '.data.pairs[0].request.headers = ($REQUEST_HEADERS)' |
    #jq --arg RESPONSE_BODY $RESPONSE_BODY '.data.pairs[0].response.body = ($RESPONSE_BODY)' |
    jq '.data.pairs[0].response.encodedBody = false' |
    jq --arg RESPONSE_HEADERS $RESPONSE_HEADERS '.data.pairs[0].response.headers = ($RESPONSE_HEADERS)' |
    jq '.data.pairs[0].response.templated = false' |
    jq '.meta.schemaVersion = "v5"' |
    jq '.meta.hoverflyVersion = "v0.17.7"' |
    jq '.data.globalActions.delays = []'
