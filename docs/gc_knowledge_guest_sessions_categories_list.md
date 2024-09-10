## gc knowledge guest sessions categories list

Get categories

### Synopsis

Get categories

```
gc knowledge guest sessions categories list [sessionId] [flags]
```

### Options

```
      --after string                  The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate                  Automatically paginate through the results stripping page information
      --before string                 The cursor that points to the start of the set of entities that has been returned.
      --expand string                 The specified entity attribute will be filled. Supported value:Ancestors: every ancestors will be filled via the parent attribute recursively,but only the id, name, parentId will be present for the ancestors.
      --filtercondition string        Filter list command output based on a given condition or regular expression
  -h, --help                          help for list
      --includeDocumentCount string   If specified, retrieves the number of documents related to category. Valid values: true, false
      --isRoot string                 If specified, retrieves only the root categories. Valid values: true, false
      --name string                   Filter to return the categories that starts with the given category name.
      --pageSize string               Number of entities to return. Maximum of 200.
      --parentId string               If specified, retrieves the children categories by parent category ID.
      --sortBy string                 Name: sort by category names alphabetically; Hierarchy: sort by the full path of hierarchical category names alphabetically Valid values: Name, Hierarchy
  -s, --stream                        Paginate and stream data as it is being processed leaving page information intact
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

* [gc knowledge guest sessions categories](gc_knowledge_guest_sessions_categories.html)	 - /api/v2/knowledge/guest/sessions/{sessionId}/categories


