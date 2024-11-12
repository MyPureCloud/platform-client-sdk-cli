## gc externalcontacts contacts notes get

Fetch a note for an external contact

### Synopsis

Fetch a note for an external contact

```
gc externalcontacts contacts notes get [contactId] [noteId] [flags]
```

### Options

```
      --expand strings   which fields, if any, to expand Valid values: author, externalDataSources, division
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

* [gc externalcontacts contacts notes](gc_externalcontacts_contacts_notes.html)	 - /api/v2/externalcontacts/contacts/{contactId}/notes


