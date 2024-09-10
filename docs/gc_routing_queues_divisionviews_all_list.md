## gc routing queues divisionviews all list

Get a paged listing of simplified queue objects, sorted by name.  Can be used to get a digest of all queues in an organization.

### Synopsis

Get a paged listing of simplified queue objects, sorted by name.  Can be used to get a digest of all queues in an organization.

```
gc routing queues divisionviews all list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           Page number (default 1)
      --pageSize int             Page size [max value is 500] (default 25)
      --sortOrder string         Sort order Valid values: asc, desc
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

* [gc routing queues divisionviews all](gc_routing_queues_divisionviews_all.html)	 - /api/v2/routing/queues/divisionviews/all


