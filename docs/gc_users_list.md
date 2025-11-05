## gc users list

Get the list of available users.

### Synopsis

Get the list of available users.

```
gc users list [flags]
```

### Options

```
  -a, --autopaginate                       Automatically paginate through the results stripping page information
      --expand strings                     Which fields, if any, to expand. Note, expand parameters are resolved with a best effort approach and not guaranteed to be returned. If requested expand information is absolutely required, it`s recommended to use specific API requests instead. Valid values: routingStatus, presence, integrationPresence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, workPlanBidRanks, externalContactsSettings, groups, customAttributes, profileSkills, certifications, locations, skills, languages, languagePreference, employerInfo, biography, dateLastLogin, dateWelcomeSent
      --filtercondition string             Filter list command output based on a given condition or regular expression
  -h, --help                               help for list
      --id strings                         A list of user IDs to fetch by bulk
      --integrationPresenceSource string   Gets an integration presence for users instead of their defaults. This parameter will only be used when presence is provided as an expand. When using this parameter the maximum number of users that can be returned is 100. Valid values: MicrosoftTeams, ZoomPhone, EightByEight
      --jabberId strings                   A list of jabberIds to fetch by bulk (cannot be used with the id parameter)
      --pageNumber string                  Page number (default "1")
      --pageSize string                    Page size (default "25")
      --sortOrder string                   Ascending or descending sort order Valid values: ascending, descending
      --state string                       Only list users of this state Valid values: active, inactive, deleted, any
  -s, --stream                             Paginate and stream data as it is being processed leaving page information intact
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


