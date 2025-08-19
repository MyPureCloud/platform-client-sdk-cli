## gc externalcontacts contacts journey segments

/api/v2/externalcontacts/contacts/{contactId}/journey/segments

### Synopsis

/api/v2/externalcontacts/contacts/{contactId}/journey/segments

### Options

```
  -h, --help   help for segments
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

* [gc externalcontacts contacts journey](gc_externalcontacts_contacts_journey.html)	 - /api/v2/externalcontacts/contacts/{contactId}/journey
* [gc externalcontacts contacts journey segments create](gc_externalcontacts_contacts_journey_segments_create.html)	 - Assign/Unassign up to 10 segments to/from an external contact or, if a segment is already assigned, update the expiry date of the segment assignment. Any unprocessed segment assignments are returned in the body for the client to retry, in the event of a partial success.
* [gc externalcontacts contacts journey segments list](gc_externalcontacts_contacts_journey_segments_list.html)	 - Retrieve segment assignments by external contact ID.


