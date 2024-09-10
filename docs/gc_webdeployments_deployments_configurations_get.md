## gc webdeployments deployments configurations get

Get active configuration for a given deployment

### Synopsis

Get active configuration for a given deployment

```
gc webdeployments deployments configurations get [deploymentId] [flags]
```

### Options

```
      --expand strings   Expand instructions for the return value Valid values: supportedContent
  -h, --help             help for get
      --varType string   Get active configuration on a deployment
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

* [gc webdeployments deployments configurations](gc_webdeployments_deployments_configurations.html)	 - /api/v2/webdeployments/deployments/{deploymentId}/configurations


