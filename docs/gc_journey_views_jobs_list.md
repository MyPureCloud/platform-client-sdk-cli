## gc journey views jobs list

Get the jobs for an organization.

### Synopsis

Get the jobs for an organization.

```
gc journey views jobs list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --interval string          An absolute timeframe for filtering the jobs, expressed as an ISO 8601 interval.
      --pageNumber string        The number of the page to return (default "1")
      --pageSize string          Max number of entities to return (default "25")
      --statuses string          Job statuses to filter for
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

* [gc journey views jobs](gc_journey_views_jobs.html)	 - /api/v2/journey/views/jobs

