## gc flows versions health

/api/v2/flows/{flowId}/versions/{versionId}/health

### Synopsis

/api/v2/flows/{flowId}/versions/{versionId}/health

### Options

```
  -h, --help   help for health
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

* [gc flows versions](gc_flows_versions.html)	 - /api/v2/flows/{flowId}/versions
* [gc flows versions health get](gc_flows_versions_health_get.html)	 - Get overall health scores for all intents present in the NLU domain version associated with the bot flow version.


