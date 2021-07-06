#!/bin/bash
cluster=$1; shift
# kubectl --kubeconfig=./$cluster/auth/kubeconfig $*
export KUBECONFIG=/root/kubeconfig-$cluster
/root/oc/oc $*
