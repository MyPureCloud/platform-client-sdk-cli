## gc scim groups update

Modify a group

### Synopsis

Modify a group

```
gc scim groups update [groupId] [flags]
```

### Options

```
      --attributes strings           Indicates which attributes to include. Returns these attributes and the id, active, and meta attributes. Use attributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId
  -d, --directory string             Directory path with files containing request bodies
      --excludedAttributes strings   Indicates which attributes to exclude. Returns the default attributes minus excludedAttributes. Always returns id, active, and meta attributes. Use excludedAttributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId
  -f, --file string                  File name containing the JSON body
  -h, --help                         help for update
      --ifMatch string               The ETag of a resource in double quotes. Returned as header and meta.version with initial call to GET /api/v2/scim/groups/{groupId}. Example: 42. If the ETag is different from the version on the server, returns 400 with a scimType of invalidVers.
  -b, --printrequestbody             Print the request body format of the API.
```

### Options inherited from parent commands

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc scim groups](gc_scim_groups.html)	 - /api/v2/scim/groups


