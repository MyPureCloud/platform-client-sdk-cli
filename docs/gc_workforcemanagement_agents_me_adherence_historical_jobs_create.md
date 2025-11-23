## gc workforcemanagement agents me adherence historical jobs create

Request an agent historical adherence report

### Synopsis

Request an agent historical adherence report

```
gc workforcemanagement agents me adherence historical jobs create [flags]
```

### Options

```
  -d, --directory string   Directory path with files containing request bodies
      --expand strings     Which fields, if any, to expand with. wfm:AgentHistoricalAdherenceConformance:view permission is required for conformance, and wfm:agentSchedule:view permission is required for scheduledActivities. Valid values: exceptionInfo, actuals, scheduledActivities, conformance
  -f, --file string        File name containing the JSON body
  -h, --help               help for create
  -b, --printrequestbody   Print the request body format of the API.
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

* [gc workforcemanagement agents me adherence historical jobs](gc_workforcemanagement_agents_me_adherence_historical_jobs.html)	 - /api/v2/workforcemanagement/agents/me/adherence/historical/jobs


