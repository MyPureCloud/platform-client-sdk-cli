## gc flows versions health get

Get overall health scores for all intents present in the NLU domain version associated with the bot flow version.

### Synopsis

Get overall health scores for all intents present in the NLU domain version associated with the bot flow version.

```
gc flows versions health get [flowId] [versionId] [flags]
```

### Options

```
  -h, --help              help for get
      --language string   Language to filter for Valid values: en-us, en-gb, en-au, en-za, en-nz, en-ie, fr-ca, fr-fr, es-us, es-es, es-mx, de-de, it-it, pt-br, pt-pt, nl-nl
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

* [gc flows versions health](gc_flows_versions_health.html)	 - /api/v2/flows/{flowId}/versions/{versionId}/health


