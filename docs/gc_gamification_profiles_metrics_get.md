## gc gamification profiles metrics get

Performance profile gamified metric by id

### Synopsis

Performance profile gamified metric by id

```
gc gamification profiles metrics get [profileId] [metricId] [flags]
```

### Options

```
  -h, --help             help for get
      --workday string   The objective query workday. If not specified, then it retrieves the current objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
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

* [gc gamification profiles metrics](gc_gamification_profiles_metrics.html)	 - /api/v2/gamification/profiles/{sourceProfileId}/metrics /api/v2/gamification/profiles/{profileId}/metrics

