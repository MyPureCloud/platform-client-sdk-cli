## gc audits query results list

Get results of audit query

### Synopsis

Get results of audit query

```
gc audits query results list [transactionId] [flags]
```

### Options

```
      --allowRedirect string     Result sets with large amounts of data will respond with a download url Valid values: true, false
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Indicates where to resume query results (not required for first page)
      --expand strings           Which fields, if any, to expand Valid values: user
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageSize string          Indicates maximum number of results in response. Default page size is 25 results. The maximum page size is 500. (default "25")
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

* [gc audits query results](gc_audits_query_results.html)	 - /api/v2/audits/query/{transactionId}/results

