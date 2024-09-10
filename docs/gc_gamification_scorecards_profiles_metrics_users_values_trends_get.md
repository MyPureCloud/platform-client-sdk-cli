## gc gamification scorecards profiles metrics users values trends get

Average performance values trends by metric of a division or a performance profile

### Synopsis

Average performance values trends by metric of a division or a performance profile

```
gc gamification scorecards profiles metrics users values trends get [profileId] [metricId] [flags]
```

### Options

```
      --endWorkday string         End workday of querying workdays range. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd - REQUIRED
      --filterId string           ID for the filter type. Only required when filterType is Division.
      --filterType string         Filter type for the query request. - REQUIRED Valid values: PerformanceProfile, Division
  -h, --help                      help for get
      --referenceWorkday string   Reference workday for the trend. Used to determine the associated metric definition. If not set, then the value of endWorkday is used. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
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

* [gc gamification scorecards profiles metrics users values trends](gc_gamification_scorecards_profiles_metrics_users_values_trends.html)	 - /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/values/trends /api/v2/gamification/scorecards/profiles/{profileId}/metrics/{metricId}/users/{userId}/values/trends


