## gc journey outcomes list

Retrieve all outcomes.

### Synopsis

Retrieve all outcomes.

```
gc journey outcomes list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --outcomeIds strings       IDs of outcomes to return. Use of this parameter is not compatible with pagination, sorting or querying. A maximum of 20 outcomes are allowed per request.
      --pageNumber string        Page number (default "1")
      --pageSize string          Page size (default "25")
      --queryFields queryValue   Outcome field(s) to query on. Requires queryValue to also be set.
      --queryValue queryFields   Value to query on. Requires queryFields to also be set.
      --sortBy -                 Field(s) to sort by. The response can be sorted by any first level property on the Outcome response. Prefix with - for descending (e.g. sortBy=displayName,-createdDate).
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

* [gc journey outcomes](gc_journey_outcomes.html)	 - /api/v2/journey/outcomes


