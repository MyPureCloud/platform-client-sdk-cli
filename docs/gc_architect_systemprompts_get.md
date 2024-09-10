## gc architect systemprompts get

Get a system prompt

### Synopsis

Get a system prompt

```
gc architect systemprompts get [promptId] [flags]
```

### Options

```
  -h, --help                      help for get
      --includeMediaUris string   Include the media URIs for each resource Valid values: true, false
      --includeResources string   Include the resources for each system prompt Valid values: true, false
      --language strings          Filter the resources down to the provided languages
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

* [gc architect systemprompts](gc_architect_systemprompts.html)	 - /api/v2/architect/systemprompts


