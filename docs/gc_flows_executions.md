## gc flows executions

/api/v2/flows/executions

### Synopsis

/api/v2/flows/executions

### Options

```
  -h, --help   help for executions
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

* [gc flows](gc_flows.html)	 - /api/v2/flows
* [gc flows executions create](gc_flows_executions_create.html)	 - Launch an instance of a flow definition, for flow types that support it such as the `workflow` type.
* [gc flows executions get](gc_flows_executions_get.html)	 - Get a flow execution`s details. Flow execution details are available for several days after the flow is started.


