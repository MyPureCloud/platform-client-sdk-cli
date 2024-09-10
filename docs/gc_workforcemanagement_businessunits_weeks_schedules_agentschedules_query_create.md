## gc workforcemanagement businessunits weeks schedules agentschedules query create

Loads agent schedule data from the schedule. Used in combination with the metadata route

### Synopsis

Loads agent schedule data from the schedule. Used in combination with the metadata route

```
gc workforcemanagement businessunits weeks schedules agentschedules query create [businessUnitId] [weekId] [scheduleId] [flags]
```

### Options

```
  -d, --directory string              Directory path with files containing request bodies
  -f, --file string                   File name containing the JSON body
      --forceAsync string             Force the result of this operation to be sent asynchronously via notification. For testing/app development purposes Valid values: true, false
      --forceDownloadService string   Force the result of this operation to be sent via download service. For testing/app development purposes Valid values: true, false
  -h, --help                          help for create
  -b, --printrequestbody              Print the request body format of the API.
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

* [gc workforcemanagement businessunits weeks schedules agentschedules query](gc_workforcemanagement_businessunits_weeks_schedules_agentschedules_query.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/weeks/{weekId}/schedules/{scheduleId}/agentschedules/query


