## gc telephony providers edges trunks listedgetrunks

Get the list of available trunks for the given Edge.

### Synopsis

Get the list of available trunks for the given Edge.

```
gc telephony providers edges trunks listedgetrunks [edgeId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for listedgetrunks
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --sortBy string            Value by which to sort
      --sortOrder string         Sort order
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
      --trunkBaseId string       Filter by Trunk Base Ids
      --trunkType string         Filter by a Trunk type Valid values: EXTERNAL, PHONE, EDGE
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

* [gc telephony providers edges trunks](gc_telephony_providers_edges_trunks.html)	 - /api/v2/telephony/providers/edges/trunks /api/v2/telephony/providers/edges/{edgeId}/trunks


