## gc users profile get

Get user profile

### Synopsis

Get user profile

```
gc users profile get [userId] [flags]
```

### Options

```
      --expand strings                     Which fields, if any, to expand Valid values: routingStatus, presence, integrationPresence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, workPlanBidRanks
  -h, --help                               help for get
      --integrationPresenceSource string   Gets an integration presence for a user instead of their default. Valid values: MicrosoftTeams, ZoomPhone, EightByEight
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

* [gc users profile](gc_users_profile.html)	 - /api/v2/users/{userId}/profile

