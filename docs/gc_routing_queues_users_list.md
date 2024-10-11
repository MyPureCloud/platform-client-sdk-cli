## gc routing queues users list

DEPRECATED: use GET /routing/queues/{queueId}/members.  Get the members of this queue.

### Synopsis

DEPRECATED: use GET /routing/queues/{queueId}/members.  Get the members of this queue.

```
gc routing queues users list [queueId] [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --expand strings           Which fields, if any, to expand. Valid values: routingStatus, presence, integrationPresence, conversationSummary, outOfOffice, geolocation, station, authorization, lasttokenissued, authorization.unusedRoles, team, workPlanBidRanks, externalContactsSettings, profileSkills, certifications, locations, groups, skills, languages, languagePreference, employerInfo, biography, dateLastLogin
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --joined string            Filter by joined status Valid values: true, false
      --languages strings        Filter by language
      --name string              Filter by queue member name
      --pageNumber int            (default 1)
      --pageSize int             Max value is 100 (default 25)
      --presence strings         Filter by presence
      --profileSkills strings    Filter by profile skill
      --routingStatus strings    Filter by routing status
      --skills strings           Filter by skill
      --sortOrder string         Note: results are sorted by name. Valid values: asc, desc
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

* [gc routing queues users](gc_routing_queues_users.html)	 - /api/v2/routing/queues/{queueId}/users


