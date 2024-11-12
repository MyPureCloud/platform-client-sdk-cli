## gc recordings retention query list

Query for recording retention data

### Synopsis

Query for recording retention data

```
gc recordings retention query list [flags]
```

### Options

```
  -a, --autopaginate                    Automatically paginate through the results stripping page information
      --cursor string                   Indicates where to resume query results (not required for first page)
      --filtercondition string          Filter list command output based on a given condition or regular expression
  -h, --help                            help for list
      --pageSize string                 Page size. Maximum is 500. (default "25")
      --retentionThresholdDays string   Fetch retention data for recordings retained for more days than the provided value. - REQUIRED
  -s, --stream                          Paginate and stream data as it is being processed leaving page information intact
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

* [gc recordings retention query](gc_recordings_retention_query.html)	 - /api/v2/recordings/retention/query


