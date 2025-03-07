## gc integrations actions draft get

Retrieve a Draft

### Synopsis

Retrieve a Draft

```
gc integrations actions draft get [actionId] [flags]
```

### Options

```
      --expand string          Indicates a field in the response which should be expanded. Valid values: contract
      --flatten string         Indicates the response should be reformatted, based on Architect`s flattening format. Valid values: true, false Valid values: true, false
  -h, --help                   help for get
      --includeConfig string   Return config in response. Valid values: true, false Valid values: true, false
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

* [gc integrations actions draft](gc_integrations_actions_draft.html)	 - /api/v2/integrations/actions/{actionId}/draft


