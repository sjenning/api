#!/bin/bash

# This script is meant to be the entrypoint for OpenShift Bash scripts to import all of the support
# libraries at once in order to make Bash script preambles as minimal as possible. This script recur-
# sively `source`s *.sh files in this directory tree. As such, no files should be `source`ed outside
# of this script to ensure that we do not attempt to overwrite read-only variables.

set -o errexit
set -o nounset
set -o pipefail

API_GROUP_VERSIONS="\
apps/v1 \
authorization/v1 \
build/v1 \
config/v1 \
image/v1 \
kubecontrolplane/v1 \
legacyconfig/v1 \
network/v1 \
oauth/v1 \
openshiftcontrolplane/v1 \
operator/v1 \
operator/v1alpha1 \
project/v1 \
quota/v1 \
route/v1 \
security/v1 \
servicecertsigner/v1alpha1 \
template/v1 \
user/v1 \
webconsole/v1 \
"
API_PACKAGES="\
github.com/sjenning/api/apps/v1,\
github.com/sjenning/api/authorization/v1,\
github.com/sjenning/api/build/v1,\
github.com/sjenning/api/image/v1,\
github.com/sjenning/api/network/v1,\
github.com/sjenning/api/oauth/v1,\
github.com/sjenning/api/project/v1,\
github.com/sjenning/api/quota/v1,\
github.com/sjenning/api/route/v1,\
github.com/sjenning/api/security/v1,\
github.com/sjenning/api/template/v1,\
github.com/sjenning/api/user/v1\
"
