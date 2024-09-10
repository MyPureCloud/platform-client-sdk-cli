## gc flows versions intents utterances health get

Get health metrics associated with a specific utterance of an intent.

### Synopsis

Get health metrics associated with a specific utterance of an intent.

```
gc flows versions intents utterances health get [flowId] [versionId] [intentId] [utteranceId] [flags]
```

### Options

```
  -h, --help              help for get
      --language string   Language to filter for - REQUIRED Valid values: en-us, en-gb, en-au, en-za, en-nz, en-ie, fr-ca, fr-fr, es-us, es-es, es-mx, de-de, it-it, pt-br, pt-pt, nl-nl
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

* [gc flows versions intents utterances health](gc_flows_versions_intents_utterances_health.html)	 - /api/v2/flows/{flowId}/versions/{versionId}/intents/{intentId}/utterances/{utteranceId}/health


