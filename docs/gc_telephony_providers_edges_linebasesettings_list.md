## gc telephony providers edges linebasesettings list

Get a listing of line base settings objects

### Synopsis

Get a listing of line base settings objects

```
gc telephony providers edges linebasesettings list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Fields to expand in the response, comma-separated Valid values: properties
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size (default 25)
      --sortBy string            Value by which to sort
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

* [gc telephony providers edges linebasesettings](gc_telephony_providers_edges_linebasesettings.html)	 - /api/v2/telephony/providers/edges/linebasesettings

