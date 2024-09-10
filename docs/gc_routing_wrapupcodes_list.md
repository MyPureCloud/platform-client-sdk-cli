## gc routing wrapupcodes list

Get list of wrapup codes.

### Synopsis

Get list of wrapup codes.

```
gc routing wrapupcodes list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionId strings       Filter by division ID(s)
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               Filter by wrapup code ID(s)
      --name s name (            Wrapup codes name (Sort by` param is ignored unless this field is provided)
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --sortBy string            Sort by Valid values: name, id
      --sortOrder string         Sort order Valid values: ascending, descending
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc routing wrapupcodes](gc_routing_wrapupcodes.html)	 - /api/v2/routing/wrapupcodes


