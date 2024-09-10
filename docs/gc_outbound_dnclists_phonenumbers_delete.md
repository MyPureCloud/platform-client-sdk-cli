## gc outbound dnclists phonenumbers delete

Deletes all or expired phone numbers from a DNC list.

### Synopsis

Deletes all or expired phone numbers from a DNC list.

```
gc outbound dnclists phonenumbers delete [dncListId] [flags]
```

### Options

```
      --expiredOnly string   Set to true to only remove DNC entries that are expired Valid values: true, false
  -h, --help                 help for delete
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

* [gc outbound dnclists phonenumbers](gc_outbound_dnclists_phonenumbers.html)	 - /api/v2/outbound/dnclists/{dncListId}/phonenumbers


