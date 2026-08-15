## gc dependencies type id connections requires list

Get entities that the given entity requires

### Synopsis

Get entities that the given entity requires

```
gc dependencies type id connections requires list [entityType] [entityId] [flags]
```

### Options

```
      --afterSourceId string      Cursor for next page
      --afterSourceType string    Cursor for next page
  -a, --autopaginate              Automatically paginate through the results stripping page information
      --beforeSourceId string     Cursor for previous page
      --beforeSourceType string   Cursor for previous page
      --filtercondition string    Filter list command output based on a given condition or regular expression
  -h, --help                      help for list
      --pageSize string           Page size (max 100)
  -s, --stream                    Paginate and stream data as it is being processed leaving page information intact
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

* [gc dependencies type id connections requires](gc_dependencies_type_id_connections_requires.html)	 - /api/v2/dependencies/type/{entityType}/id/{entityId}/connections/requires


