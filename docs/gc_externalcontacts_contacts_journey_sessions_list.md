## gc externalcontacts contacts journey sessions list

Retrieve all sessions for a given external contact.

### Synopsis

Retrieve all sessions for a given external contact.

```
gc externalcontacts contacts journey sessions list [contactId] [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --includeMerged string     Indicates whether to return sessions from all external contacts in the merge-set of the given one. Valid values: true, false
      --pageSize string          Number of entities to return. Maximum of 200.
  -s, --stream                   Paginate and stream data as it is being processed leaving page information intact
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

* [gc externalcontacts contacts journey sessions](gc_externalcontacts_contacts_journey_sessions.html)	 - /api/v2/externalcontacts/contacts/{contactId}/journey/sessions


