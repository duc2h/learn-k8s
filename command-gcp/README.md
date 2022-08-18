#### Some command on gcp

##### Find who deleted the pod by audit log

```
resource.type="k8s_cluster"
resource.labels.cluster_name="staging-2"
resource.labels.location="asia-southeast1-b"
logName="projects/staging-manabie-online/logs/cloudaudit.googleapis.com%2Factivity"
protoPayload.methodName="io.k8s.core.v1.pods.delete" # action
protoPayload.response.metadata.namespace="stag-jprep-services"
protoPayload.response.metadata.name=~"zeus" # search name of pod
protoPayload.authenticationInfo.principalEmail=~"@manabie.com" # search `@manabie.com` prefix email.
```

##### get all service account use kms key

```
gcloud kms keys get-iam-policy stag-manabie-zeus --keyring=separate-kms-test --location=global
```

##### get all roles of serviceaccount

```
gcloud projects get-iam-policy staging-manabie-online  \
--flatten="bindings[].members" \
--format='table(bindings.role)' \
--filter="bindings.members:stag-yasuo@staging-manabie-online.iam.gserviceaccount.com"
```