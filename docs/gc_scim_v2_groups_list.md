## gc scim v2 groups list

Get a list of groups

### Synopsis

Get a list of groups

```
gc scim v2 groups list [flags]
```

### Options

```
      --attributes strings           Indicates which attributes to include. Returns these attributes and the id, active, and meta attributes. Use attributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId
  -a, --autopaginate                 Automatically paginate through the results stripping page information
      --count string                 The requested number of items per page. A value of 0 returns totalResults. A page size over 25 may exceed internal resource limits and return a 429 error. For a page size over 25, use the excludedAttributes or attributes query parameters to exclude or only include secondary lookup values such as externalId,  roles, urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingLanguages, or urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingSkills. (default "25")
      --excludedAttributes strings   Indicates which attributes to exclude. Returns the default attributes minus excludedAttributes. Always returns id, active, and meta attributes. Use excludedAttributes to avoid expensive secondary calls for the default attributes. Valid values: id, displayName, members, externalId, meta, meta.version, meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:id, urn:ietf:params:scim:schemas:core:2.0:Group:meta, urn:ietf:params:scim:schemas:core:2.0:Group:meta.version, urn:ietf:params:scim:schemas:core:2.0:Group:meta.lastModified, urn:ietf:params:scim:schemas:core:2.0:Group:displayName, urn:ietf:params:scim:schemas:core:2.0:Group:members, urn:ietf:params:scim:schemas:core:2.0:Group:externalId
      --filter string                Filters results. If nothing is specified, returns all groups. Examples of valid values: id eq 5f4bc742-a019-4e38-8e2a-d39d5bc0b0f3, displayname eq Sales. - REQUIRED
      --filtercondition string       Filter list command output based on a given condition or regular expression
  -h, --help                         help for list
      --startIndex string            The 1-based index of the first query result. (default "1")
  -s, --stream                       Paginate and stream data as it is being processed leaving page information intact
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

* [gc scim v2 groups](gc_scim_v2_groups.html)	 - /api/v2/scim/v2/groups


