## gc telephony providers edges didpools dids list

Get a listing of unassigned and/or assigned numbers in a set of DID Pools.

### Synopsis

Get a listing of unassigned and/or assigned numbers in a set of DID Pools.

```
gc telephony providers edges didpools dids list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               Filter by a specific list of DID Pools.  If this is not provided, numbers from all DID Pools will be returned.
      --numberMatch string       A number to filter the results by.
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --sortOrder string         Sort order
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --varType string           The type of numbers to return. - REQUIRED Valid values: ASSIGNED_AND_UNASSIGNED, UNASSIGNED
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

* [gc telephony providers edges didpools dids](gc_telephony_providers_edges_didpools_dids.html)	 - /api/v2/telephony/providers/edges/didpools/dids


