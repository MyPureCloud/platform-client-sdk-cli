## gc outbound contactlists contacts deletecontacts

Delete contacts from a contact list.

### Synopsis

Delete contacts from a contact list.

```
gc outbound contactlists contacts deletecontacts [contactListId] [flags]
```

### Options

```
      --contactIds strings   ContactIds to delete. - REQUIRED
  -h, --help                 help for deletecontacts
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

* [gc outbound contactlists contacts](gc_outbound_contactlists_contacts.html)	 - /api/v2/outbound/contactlists/{contactListId}/contacts


