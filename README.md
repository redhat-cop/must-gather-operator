# Resource Locker Operator

[![Build Status](https://travis-ci.org/redhat-cop/must-gather-operator.svg?branch=master)](https://travis-ci.org/redhat-cop/must-gather-operator) [![Docker Repository on Quay](https://quay.io/repository/redhat-cop/must-gather-operator/status "Docker Repository on Quay")](https://quay.io/repository/redhat-cop/must-gather-operator)

The must gather operator helps collecting must gather information on a cluster and uploading it to a case.
To use the operator a cluster administrator can create the following must gather CR:

```yaml
apiVersion: redhatcop.redhat.io/v1alpha1
kind: MustGather
metadata:
  name: example-mustgather
spec:
  caseID: '02527285'
  caseManagementAccountSecretRef:
    name: case-management-creds
```

this request will collect the standard must gather info and upload it to case `#02527285` using the credentials found in the `caseManagementCreds` secret.

A more complex example:

```yaml
apiVersion: redhatcop.redhat.io/v1alpha1
kind: MustGather
metadata:
  name: full-mustgather
spec:
  caseID: '02527285'
  caseManagementAccountSecretRef:
    name: case-management-creds
  serviceAccountRef:
    name: must-gather-admin
  mustGatherImages:
  - quay.io/kubevirt/must-gather:latest
  - quay.io/ocs-dev/ocs-must-gather
```

in this example we are using a specific service account (which must have cluster admin permissions as per must-gather requirements) and we are specifying a couple of additional must gather images to be run for the `kubevirt` and `ocs` subsystem. If not specified serviceAccountRef.Name will default to `default`. Also the standard must gather image: `quay.io/openshift/origin-must-gather:latest` is always added by default.

## Deploying the Operator

This is a cluster-level operator that you can deploy in any namespace, `must-gather-operator` is recommended.

You can either deploy it using [`Helm`](https://helm.sh/) or creating the manifests directly.

### Deploying with Helm

Here are the instructions to install the latest release with Helm.

```shell
oc new-project must-gather-operator

helm repo add must-gather-operator https://redhat-cop.github.io/must-gather-operator
helm repo update
export must_gather_operator_chart_version=$(helm search must-gather-operator/must-gather-operator | grep must-gather-operator/must-gather-operator | awk '{print $2}')

helm fetch must-gather-operator/must-gather-operator --version ${must_gather_operator_chart_version}
helm template must-gather-operator-${must_gather_operator_chart_version}.tgz --namespace must-gather-operator | oc apply -f - -n must-gather-operator

rm must-gather-operator-${must_gather_operator_chart_version}.tgz
```

### Deploying directly with manifests

Here are the instructions to install the latest release creating the manifest directly in OCP.

```shell
git clone git@github.com:redhat-cop/must-gather-operator.git; cd must-gather-operator
oc apply -f deploy/crds/redhatcop.redhat.io_mustgathers_crd.yaml
oc new-project must-gather-operator
oc -n must-gather-operator apply -f deploy
```

### Meeting the operator requirements

In order to run the operator needs a secret and a service accout to be created by the admin as follows (this assume the operato is running in the `must-gather-operator` namespace)

```shell
oc create sa must-gather-admin -n must-gather-operator
oc adm policy add-cluster-role-to-user cluster-admin -z must-gather-admin -n must-gather-operator
oc create secret generic case-management-creds --from-literal=username=<username> --from-literal=password=<password>
```

## Local Development

Execute the following steps to develop the functionality locally. It is recommended that development be done using a cluster with `cluster-admin` permissions.

```shell
go mod download
```

optionally:

```shell
go mod vendor
```

Using the [operator-sdk](https://github.com/operator-framework/operator-sdk), run the operator locally:

```shell
oc apply -f deploy/crds/redhatcop.redhat.io_mustgathers_crd.yaml
oc new-project must-gather-operator
OPERATOR_NAME='must-gather-operator' operator-sdk --verbose up local --namespace "must-gather-operator"
```

## Release Process

To release execute the following:

```shell
git tag -a "<version>" -m "release <version>"
git push upstream <version>
```

use this version format: vM.m.z
