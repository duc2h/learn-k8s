#### Note:

##### Go
When parsing a config from an file (yaml, json). `struct or field` is used for map the value from a file
should be uppercase with the first character, the reason due to `encapsulation in go`. We cannot unmarshal a private field.
