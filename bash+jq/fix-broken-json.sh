#!/bin/bash
PACTS=$1
for PACT in $PACTS
do
	#echo $PACT
	hjson -j $PACT
done
