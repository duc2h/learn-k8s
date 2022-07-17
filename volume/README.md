#### Note:

##### Go
When parsing a config from an file (yaml, json). `struct or field` is used for map the value from a file
should be uppercase with the first character, the reason due to `encapsulation in go`. We cannot unmarshal a private field.

##### Secret
The secret is only stored in memory instead of physical storage of worker node where the pod is created, when the worker node 
is die the secret will be removed. K8s will encode your content by base64.