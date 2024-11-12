## gc flows datatables list

Retrieve a list of datatables for the org

### Synopsis

Retrieve a list of datatables for the org

```
gc flows datatables list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionId strings       division ID(s)
      --expand string            Expand instructions for the result Valid values: schema
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --name string              Filter by Name. The wildcard character * is supported within the filter. Matches are case-insensitive.
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --sortBy string            Sort by Valid values: id, name
      --sortOrder string         Sort order
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

* [gc flows datatables](gc_flows_datatables.html)	 - /api/v2/flows/datatables


