## gc analytics flows aggregates jobs results list

Fetch a page of results for an async aggregates query

### Synopsis

Fetch a page of results for an async aggregates query

```
gc analytics flows aggregates jobs results list [jobId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --cursor string            Cursor token to retrieve next page
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
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

* [gc analytics flows aggregates jobs results](gc_analytics_flows_aggregates_jobs_results.html)	 - /api/v2/analytics/flows/aggregates/jobs/{jobId}/results


