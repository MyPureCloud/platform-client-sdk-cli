## gc journey sessions events list

Retrieve all events for a given session.

### Synopsis

Retrieve all events for a given session.

```
gc journey sessions events list [sessionId] [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --eventType string         A comma separated list of journey event types to include in the results. Valid values: com.genesys.journey.OutcomeAchievedEvent, com.genesys.journey.SegmentAssignmentEvent, com.genesys.journey.WebActionEvent, com.genesys.journey.WebEvent, com.genesys.journey.AppEvent
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
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

* [gc journey sessions events](gc_journey_sessions_events.html)	 - /api/v2/journey/sessions/{sessionId}/events


