## gc flows latestconfiguration get

Get the latest configuration for flow

### Synopsis

Get the latest configuration for flow

```
gc flows latestconfiguration get [flowId] [flags]
```

### Options

```
      --deleted string   Deleted flows Valid values: true, false
  -h, --help             help for get
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

* [gc flows latestconfiguration](gc_flows_latestconfiguration.html)	 - /api/v2/flows/{flowId}/latestconfiguration


