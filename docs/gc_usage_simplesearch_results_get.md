## gc usage simplesearch results get

Get the results of a usage search. Number of records to be returned is limited to 20,000 results.

### Synopsis

Get the results of a usage search. Number of records to be returned is limited to 20,000 results.

```
gc usage simplesearch results get [executionId] [flags]
```

### Options

```
      --after string   The cursor that points to the end of the set of entities that has been returned
  -h, --help           help for get
      --pageSize int   The max number of entities to be returned per request. Maximum page size of 1000
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

* [gc usage simplesearch results](gc_usage_simplesearch_results.html)	 - /api/v2/usage/simplesearch/{executionId}/results


