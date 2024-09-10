## gc knowledge knowledgebases operations list

Get operations

### Synopsis

Get operations

```
gc knowledge knowledgebases operations list [knowledgeBaseId] [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --interval string          Retrieves the operations modified in specified date and time range. If the after and before cursor parameters are within this interval, it would return valid data, otherwise it throws an error.The dates in the interval are represented in ISO-8601 format: YYYY-MM-DDThh:mm:ssZ/YYYY-MM-DDThh:mm:ssZ
      --pageSize string          Number of entities to return. Maximum of 200.
      --sourceId strings         If specified, retrieves operations associated with source ids, comma separated values expected.
      --status strings           If specified, retrieves operations with specified operation status, comma separated values expected.
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --userId strings           If specified, retrieves operations associated with user ids, comma separated values expected.
      --varType strings          If specified, retrieves operations with specified operation type, comma separated values expected. Valid values: Export, Import, Parse, Sync
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

* [gc knowledge knowledgebases operations](gc_knowledge_knowledgebases_operations.html)	 - /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/operations


