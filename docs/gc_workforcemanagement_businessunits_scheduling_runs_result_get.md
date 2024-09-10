## gc workforcemanagement businessunits scheduling runs result get

Get the result of a rescheduling operation

### Synopsis

Get the result of a rescheduling operation

```
gc workforcemanagement businessunits scheduling runs result get [businessUnitId] [runId] [flags]
```

### Options

```
      --expand strings              The fields to expand. Omitting will return an empty response - REQUIRED Valid values: headcountForecast, generationResults, agentSchedules
  -h, --help                        help for get
      --managementUnitIds strings   The IDs of the management units for which to fetch the reschedule results - REQUIRED
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

* [gc workforcemanagement businessunits scheduling runs result](gc_workforcemanagement_businessunits_scheduling_runs_result.html)	 - /api/v2/workforcemanagement/businessunits/{businessUnitId}/scheduling/runs/{runId}/result


