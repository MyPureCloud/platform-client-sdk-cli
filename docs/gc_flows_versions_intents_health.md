## gc flows versions intents health

/api/v2/flows/{flowId}/versions/{versionId}/intents/{intentId}/health

### Synopsis

/api/v2/flows/{flowId}/versions/{versionId}/intents/{intentId}/health

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

* [gc flows versions intents](gc_flows_versions_intents.html)	 - /api/v2/flows/{flowId}/versions/{versionId}/intents
* [gc flows versions intents health get](gc_flows_versions_intents_health_get.html)	 - Get health scores and other health metrics for a specific intent. This includes the health metrics for each utterance in an intent.


