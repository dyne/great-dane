#!/bin/bash
DOMAIN=${2:-"zenr.io"}
FQDN="${1:-"test.${DOMAIN}"}"
ALGOS=( NSEC3RSASHA1 RSASHA256 RSASHA512 ECDSAP256SHA256 ECDSAP384SHA384 ED25519 RSASHA1 )
for ALGO in "${ALGOS[@]}"
  do
    echo $ALGO
    dnssec-keygen -a ${ALGO} ${FQDN}
  done
