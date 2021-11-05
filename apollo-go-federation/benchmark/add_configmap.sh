kubectl create configmap assets \
   --from-file=./assets/single_query.json \
   --from-file=./assets/composite_query.json \
   --from-file=./assets/test_single.txt \
   --from-file=./assets/test_composite.txt \
   --from-file=./assets/test_single_apollo.txt \
   --from-file=./assets/test_composite_apollo.txt
