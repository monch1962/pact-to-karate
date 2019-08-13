#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

INPUT=$(echo $(cat))

CONSUMER=$(echo $INPUT |
    jq '.consumer.name')

PROVIDER=$(echo $INPUT |
    jq '.provider.name')

NUM_INTERACTIONS=$(echo $INPUT |
    jq '.interactions | length')

echo "Feature: Provider $PROVIDER responding to consumer $CONSUMER"
echo ""
echo "Background:"
echo "  * configure cors = true"
echo ""
for INTERACTION in $(seq $NUM_INTERACTIONS)
do
    I="$(($INTERACTION - 1))"
    DESCRIPTION=$(echo $INPUT | 
        jq --raw-output --argjson i "$I" '.interactions[$i].description') 

    URL=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.path') 

    METHOD=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.method | ascii_downcase') 

    REQUEST_BODY=$(echo $INPUT |
        jq -c --raw-output --argjson i "$I" '.interactions[$i].request.body')

    STATUS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].response.status')

    RESPONSE_BODY=$(echo $INPUT |
        jq -c --raw-output --argjson i "$I" '.interactions[$i].response.body') 

    REQUEST_BODY_MATCHERS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].request.body |
        . as $data | 
        [path(.. | select(type != "object" and type != "array") )] | 
        map( . as $path | map("[" + (. | tojson) + "]") | join("") | "request\(.) == \($data | getpath($path) | tojson)" ) | 
        join(" && ")')

    #echo $REQUEST_BODY_MATCHERS

    if [ "$REQUEST_BODY_MATCHERS" != 'null' ]
    then
        echo "Scenario: pathMatches('${URL}') && methodIs('${METHOD}') && ${REQUEST_BODY_MATCHERS}"
    else
        echo "Scenario: pathMatches('${URL}') && methodIs('${METHOD}')"
    fi

    if [ "$RESPONSE_BODY" != 'null' ]; then echo "    * def response = $RESPONSE_BODY"; fi
    if [ "$STATUS" != 'null' ]; then echo "    * def responseStatus = $STATUS"; fi
    echo ""
done

echo "Scenario:"
echo "    * def responseStatus = 404"
