#!/bin/bash
export IMAGE=quay.io/omular/pcmk:test
# export IMAGE=quay.io/abeekhof/pacemaker

export CLUSTER_DOMAIN="example.com"
export SECRETS_DIR="/etc/secret-volume"
export AUTHKEY=$(cat $1 | base64 | tr -d '\n')
# export FENCEKEY=$(cat /etc/cluster/fence_xvm.key | base64 | tr -d '\n')
export NODENAME=$2

export NODE1NAME=$3
# export NODE1DOMAIN=$(grep baseDomain install-config-$2.yaml | cut -d: -f2 | tr -d ' \t')
export NODE1DOMAIN=$CLUSTER_DOMAIN
export NODE1ADDR=$(dig api.$NODE1NAME.$NODE1DOMAIN | grep -v -e "^;" | cut -s -dA -f2 | tr -d ' \t')

export NODE2NAME=$4
# export NODE2DOMAIN=$(grep baseDomain install-config-$3.yaml | cut -d: -f2 | tr -d ' \t')
export NODE2DOMAIN=$CLUSTER_DOMAIN
export NODE2ADDR=$(dig api.$NODE2NAME.$NODE2DOMAIN | grep -v -e "^;" | cut -s -dA -f2 | tr -d ' \t')
	    
envsubst < pod-cluster.yaml.in > pod-$NODENAME.yaml
