## gc externalcontacts reversewhitepageslookup search

Look up contacts and externalOrganizations based on an attribute. Maximum of 25 values returned.

### Synopsis

Look up contacts and externalOrganizations based on an attribute. Maximum of 25 values returned.

```
gc externalcontacts reversewhitepageslookup search [flags]
```

### Options

```
      --divisionId string   Specifies which division to lookup contacts/externalOrganizations in, for the given lookup value
      --expand strings      which field, if any, to expand Valid values: contacts.externalOrganization, externalDataSources, division
  -h, --help                help for search
      --lookupVal string    User supplied value to lookup contacts/externalOrganizations (supports email addresses, e164 phone numbers, Twitter screen names) - REQUIRED
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

* [gc externalcontacts reversewhitepageslookup](gc_externalcontacts_reversewhitepageslookup.html)	 - /api/v2/externalcontacts/reversewhitepageslookup


