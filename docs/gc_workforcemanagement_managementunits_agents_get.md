## gc workforcemanagement managementunits agents get

Get data for agent in the management unit

### Synopsis

Get data for agent in the management unit

```
gc workforcemanagement managementunits agents get [managementUnitId] [agentId] [flags]
```

### Options

```
      --excludeCapabilities string   Excludes all capabilities of the agent such as queues, languages, and skills Valid values: true, false
      --expand strings                Valid values: workPlanOverrides
  -h, --help                         help for get
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

* [gc workforcemanagement managementunits agents](gc_workforcemanagement_managementunits_agents.html)	 - /api/v2/workforcemanagement/managementunits/{managementUnitId}/agents


