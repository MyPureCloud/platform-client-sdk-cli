## gc integrations actions draft schemas get

Retrieve schema for a Draft based on filename.

### Synopsis

Retrieve schema for a Draft based on filename.

```
gc integrations actions draft schemas get [actionId] [fileName] [flags]
```

### Options

```
      --flatten string   Indicates the response should be reformatted, based on Architect`s flattening format. Valid values: true, false
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

* [gc integrations actions draft schemas](gc_integrations_actions_draft_schemas.html)	 - /api/v2/integrations/actions/{actionId}/draft/schemas


