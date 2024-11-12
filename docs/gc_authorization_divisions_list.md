## gc authorization divisions list

Retrieve a list of all divisions defined for the organization

### Synopsis

Retrieve a list of all divisions defined for the organization

```
gc authorization divisions list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           variable name requested by expand list
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               Optionally request specific divisions by their IDs
      --name string              Search term to filter by division name
      --nextPage string          next page token
      --objectCount string       Include the count of objects contained in the division Valid values: true, false
      --pageNumber string        The page number requested (default "1")
      --pageSize string          The total page size requested (default "25")
      --previousPage string      Previous page token
      --sortBy string            variable name requested to sort by
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

* [gc authorization divisions](gc_authorization_divisions.html)	 - /api/v2/authorization/divisions


