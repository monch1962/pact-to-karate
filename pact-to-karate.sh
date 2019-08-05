#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

INPUT=$(echo $(cat))

CONSUMER=$(echo $INPUT |
    jq '.consumer.name')

PROVIDER=$(echo $INPUT |
    jq '.provider.name')

NUM_INTERACTIONS=$(echo $INPUT |
    jq '.interactions | length')

echo "Feature: Consumer $CONSUMER talking to provider $PROVIDER"
echo ""
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

    STATUS=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].response.status')

    RESPONSE_BODY=$(echo $INPUT |
        jq --raw-output --argjson i "$I" '.interactions[$i].response.body') 

    echo "  Scenario: $DESCRIPTION"
    echo "    Given url '$URL'"
    if [ "$REQUEST_BODY" != 'null' ]; then echo "    And request $REQUEST_BODY"; fi
    echo "    When method $METHOD"
    echo "    Then status $STATUS"
    if [ "$RESPONSE_BODY" != 'null' ]; then echo "    And match response == $RESPONSE_BODY"; fi
    echo ""
done

