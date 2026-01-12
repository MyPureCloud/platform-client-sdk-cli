## gc intents customerintents sourceintents list

Get source intents mapped to a customer intent

### Synopsis

Get source intents mapped to a customer intent

```
gc intents customerintents sourceintents list [customerIntentId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber string        The page number requested (default "1")
      --pageSize string          The total page size requested (default "25")
      --queryValue string        Search query value to filter results by
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

* [gc intents customerintents sourceintents](gc_intents_customerintents_sourceintents.html)	 - /api/v2/intents/customerintents/{customerIntentId}/sourceintents


