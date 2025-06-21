## gc users externalid getuserbyexternalid

Get the user associated with external identifier.

### Synopsis

Get the user associated with external identifier.

```
gc users externalid getuserbyexternalid [authorityName] [externalKey] [flags]
```

### Options

```
      --expand strings   Which fields, if any, to expand Valid values: routingStatus, presence, integrationPresence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, workPlanBidRanks, externalContactsSettings, groups, profileSkills, certifications, locations, skills, languages, languagePreference, employerInfo, biography, dateLastLogin, dateWelcomeSent
  -h, --help             help for getuserbyexternalid
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

* [gc users externalid](gc_users_externalid.html)	 - /api/v2/users/{userId}/externalid /api/v2/users/{userId}/externalid/{authorityName} /api/v2/users/externalid/{authorityName}


