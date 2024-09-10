## gc outbound dnclists get

Get dialer DNC list

### Synopsis

Get dialer DNC list

```
gc outbound dnclists get [dncListId] [flags]
```

### Options

```
  -h, --help                         help for get
      --includeImportStatus string   Import status Valid values: true, false
      --includeSize string           Include size Valid values: true, false
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

* [gc outbound dnclists](gc_outbound_dnclists.html)	 - /api/v2/outbound/dnclists


