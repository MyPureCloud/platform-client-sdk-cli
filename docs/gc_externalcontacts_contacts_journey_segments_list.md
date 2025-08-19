## gc externalcontacts contacts journey segments list

Retrieve segment assignments by external contact ID.

### Synopsis

Retrieve segment assignments by external contact ID.

```
gc externalcontacts contacts journey segments list [contactId] [flags]
```

### Options

```
  -h, --help                   help for list
      --includeMerged string   Indicates whether to return segment assignments from all external contacts in the merge-set of the given one. Valid values: true, false
      --limit string           Number of entities to return. Default of 25, maximum of 500.
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

* [gc externalcontacts contacts journey segments](gc_externalcontacts_contacts_journey_segments.html)	 - /api/v2/externalcontacts/contacts/{contactId}/journey/segments


