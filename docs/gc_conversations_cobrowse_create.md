## gc conversations cobrowse create

Creates a cobrowse session. Requires conversation:cobrowse:add (for web messaging) or conversation:cobrowsevoice:add permission.

### Synopsis

Creates a cobrowse session. Requires conversation:cobrowse:add (for web messaging) or conversation:cobrowsevoice:add permission.

```
gc conversations cobrowse create [conversationId] [flags]
```

### Options

```
  -h, --help   help for create
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

* [gc conversations cobrowse](gc_conversations_cobrowse.html)	 - /api/v2/conversations/{conversationId}/cobrowse


