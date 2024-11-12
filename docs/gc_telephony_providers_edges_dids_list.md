## gc telephony providers edges dids list

Get a listing of DIDs

### Synopsis

Get a listing of DIDs

```
gc telephony providers edges dids list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --didPoolId string         Filter by the DID Pool assignment
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --id strings               Filter by a specific list of ID`s
      --ownerId string           Filter by the owner of a phone number
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --phoneNumber string       Filter by phoneNumber
      --sortBy string            Sort by
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

* [gc telephony providers edges dids](gc_telephony_providers_edges_dids.html)	 - /api/v2/telephony/providers/edges/dids


