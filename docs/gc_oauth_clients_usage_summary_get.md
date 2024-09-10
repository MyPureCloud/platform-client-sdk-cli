## gc oauth clients usage summary get

Get a summary of OAuth client API usage

### Synopsis

Get a summary of OAuth client API usage

```
gc oauth clients usage summary get [clientId] [flags]
```

### Options

```
      --days string   Previous number of days to query
  -h, --help          help for get
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

* [gc oauth clients usage summary](gc_oauth_clients_usage_summary.html)	 - /api/v2/oauth/clients/{clientId}/usage/summary


