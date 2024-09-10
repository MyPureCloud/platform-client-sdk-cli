## gc workforcemanagement managementunits weeks schedules list

Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit

### Synopsis

Deprecated.  Use the equivalent business unit resource instead. Get the list of schedules in a week in management unit

```
gc workforcemanagement managementunits weeks schedules list [managementUnitId] [weekId] [flags]
```

### Options

```
      --earliestWeekDate string       The start date of the earliest week to query in yyyy-MM-dd format
  -h, --help                          help for list
      --includeOnlyPublished string   Return only published schedules Valid values: true, false
      --latestWeekDate string         The start date of the latest week to query in yyyy-MM-dd format
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


