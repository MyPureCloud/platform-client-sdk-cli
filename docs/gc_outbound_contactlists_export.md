## gc outbound contactlists export

/api/v2/outbound/contactlists/{contactListId}/export

### Synopsis

/api/v2/outbound/contactlists/{contactListId}/export

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

* [gc outbound contactlists](gc_outbound_contactlists.html)	 - /api/v2/outbound/contactlists
* [gc outbound contactlists export create](gc_outbound_contactlists_export_create.html)	 - Initiate the export of a contact list.
* [gc outbound contactlists export get](gc_outbound_contactlists_export_get.html)	 - Get the URI of a contact list export.


