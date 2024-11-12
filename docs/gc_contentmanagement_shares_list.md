## gc contentmanagement shares list

Gets a list of shares.  You must specify at least one filter (e.g. entityId).

### Synopsis

Gets a list of shares.  You must specify at least one filter (e.g. entityId).

```
gc contentmanagement shares list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --entityId string          Filters the shares returned to only the entity specified by the value of this parameter.
      --expand strings           Which fields, if any, to expand. Valid values: member
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
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

* [gc contentmanagement shares](gc_contentmanagement_shares.html)	 - /api/v2/contentmanagement/shares


