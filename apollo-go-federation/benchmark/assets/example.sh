#!/bin/bash

for i in `seq 100 100 1000`; do
	cat test_single.txt | vegeta attack -rate=$i -duration=1m | vegeta report
	sleep 20
done
