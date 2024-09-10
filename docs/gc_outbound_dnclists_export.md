## gc outbound dnclists export

/api/v2/outbound/dnclists/{dncListId}/export

### Synopsis

/api/v2/outbound/dnclists/{dncListId}/export

### Options

```
  -h, --help   help for export
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
* [gc outbound dnclists export create](gc_outbound_dnclists_export_create.html)	 - Initiate the export of a dnc list.
* [gc outbound dnclists export get](gc_outbound_dnclists_export_get.html)	 - Get the URI of a DNC list export.


