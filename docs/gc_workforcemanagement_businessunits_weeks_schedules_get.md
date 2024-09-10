## gc workforcemanagement businessunits weeks schedules get

Get the metadata for the schedule, describing which management units and agents are in the scheduleSchedule data can then be loaded with the query route

### Synopsis

Get the metadata for the schedule, describing which management units and agents are in the scheduleSchedule data can then be loaded with the query route

```
gc workforcemanagement businessunits weeks schedules get [businessUnitId] [weekId] [scheduleId] [flags]
```

### Options

```
      --expand string   expand Valid values: managementUnits.agents
  -h, --help            help for get
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

* [gc workforcemanagement businessunits weeks schedules](gc_workforcemanagement_businessunits_weeks_schedules.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules


