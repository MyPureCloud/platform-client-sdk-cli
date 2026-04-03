## gc users get

Get user.

### Synopsis

Get user.

```
gc users get [userId] [flags]
```

### Options

```
      --expand strings                         Which fields, if any, to expand. Note, expand parameters are resolved with a best effort approach and not guaranteed to be returned. If requested expand information is absolutely required, it`s recommended to use specific API requests instead. Valid values: routingStatus, presence, integrationPresence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, workPlanBidRanks, externalContactsSettings, groups, customAttributes, profileSkills, certifications, locations, skills, languages, languagePreference, employerInfo, biography, dateLastLogin, dateWelcomeSent
  -h, --help                                   help for get
      --integrationPresenceSource string       Gets an integration presence for a user instead of their default. Valid values: MicrosoftTeams, ZoomPhone, EightByEight
      --state string                           Search for a user with this state Valid values: active, deleted
      --userCustomAttributeSchemaIds strings   Gets custom user attribute values for given schemas set for user. This parameter will only be used when customAttributes is provided as an expand. The maximum number of schemaIds that can be requested is 100
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

* [gc users](gc_users.html)	 - /api/v2/users


