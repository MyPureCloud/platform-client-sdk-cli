## gc workforcemanagement businessunits get

Get business unit

### Synopsis

Get business unit

```
gc workforcemanagement businessunits get [businessUnitId] [flags]
```

### Options

```
      --expand strings                                     Include to access additional data on the business unit Valid values: settings, settings.timeZone, settings.startDayOfWeek, settings.shortTermForecasting, settings.scheduling, settings.notifications.scheduling, settings.learning, settings.coaching
  -h, --help                                               help for get
      --includeSchedulingDefaultMessageSeverities string   Whether to include scheduling default message severities Valid values: true, false
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

* [gc workforcemanagement businessunits](gc_workforcemanagement_businessunits.html)	 - /api/v2/workforcemanagement/businessunits


