## gc telephony calls metrics get

Get the concurrent call metrics for a given organization.

### Synopsis

Get the concurrent call metrics for a given organization.

```
gc telephony calls metrics get [flags]
```

### Options

```
  -h, --help                help for get
      --metricType string   Flag to indicate metric type to fetch. Valid values: cloud, premises
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

* [gc telephony calls metrics](gc_telephony_calls_metrics.html)	 - /api/v2/telephony/calls/metrics


