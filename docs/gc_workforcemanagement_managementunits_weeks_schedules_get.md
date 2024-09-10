## gc workforcemanagement managementunits weeks schedules get

Deprecated.  Use the equivalent business unit resource instead. Get a week schedule

### Synopsis

Deprecated.  Use the equivalent business unit resource instead. Get a week schedule

```
gc workforcemanagement managementunits weeks schedules get [managementUnitId] [weekId] [scheduleId] [flags]
```

### Options

```
      --expand string                 Which fields, if any, to expand Valid values: generationResults, headcountForecast
      --forceDownloadService string   Force the result of this operation to be sent via download service.  For testing/app development purposes Valid values: true, false
  -h, --help                          help for get
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

* [gc workforcemanagement managementunits weeks schedules](gc_workforcemanagement_managementunits_weeks_schedules.html)	 - /api/v2/workforcemanagement/managementunits/{managementUnitId}/weeks/{weekId}/schedules


