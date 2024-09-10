## gc gamification scorecards values trends get

Values trends of the requesting user or group

### Synopsis

Values trends of the requesting user or group

```
gc gamification scorecards values trends get [flags]
```

### Options

```
      --endWorkday string         End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --filterType string         Filter type for the query request. If not set, then the request is for the requesting user. Valid values: PerformanceProfile, Division
  -h, --help                      help for get
      --referenceWorkday string   Reference workday for the trend. Used to determine the profile of the user as of this date. If not set, then the user`s current profile will be used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
      --startWorkday string       Start workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --timeZone string           Timezone for the workday. Defaults to UTC
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

* [gc gamification scorecards values trends](gc_gamification_scorecards_values_trends.html)	 - /api/v2/gamification/scorecards/values/trends


