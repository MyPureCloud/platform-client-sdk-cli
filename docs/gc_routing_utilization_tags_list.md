## gc routing utilization tags list

Get list of utilization tags

### Synopsis

Get list of utilization tags

```
gc routing utilization tags list [flags]
```

### Options

```
  -a, --autopaginate                                 Automatically paginate through the results stripping page information
      --filtercondition string                       Filter list command output based on a given condition or regular expression
  -h, --help                                         help for list
      --name s name (Wildcard is supported, e.g.,    Utilization tags name (Wildcard is supported, e.g., tag1*`)
      --pageNumber string                            Page number (default "1")
      --pageSize string                              Page size (default "25")
      --sortOrder string                             Sort order by name Valid values: ascending, descending
  -s, --stream                                       Paginate and stream data as it is being processed leaving page information intact
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

* [gc routing utilization tags](gc_routing_utilization_tags.html)	 - /api/v2/routing/utilization/tags


