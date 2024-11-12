## gc users development activities list

Get list of Development Activities

### Synopsis

Get list of Development Activities

```
gc users development activities list [flags]
```

### Options

```
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --completionInterval string   Specifies the range of completion dates to be used for filtering. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --interval string             Specifies the dateDue range to be queried. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --moduleId string             Specifies the ID of the learning module.
      --overdue string              Specifies if non-overdue, overdue, or all activities are returned. If not specified, all activities are returned Valid values: True, False, Any
      --pageNumber string           Page number (default "1")
      --pageSize string             Page size (default "25")
      --pass string                 Specifies if only the failed (pass is False) or passed (pass is True) activities are returned. If pass is Any or if the pass parameter is not supplied, all activities are returned Valid values: True, False, Any
      --relationship strings        Specifies how the current user relation should be interpreted, and filters the activities returned to only the activities that have the specified relationship. If a value besides Attendee is specified, it will only return Coaching Appointments. If not specified, no filtering is applied. Valid values: Creator, Facilitator, Attendee
      --sortOrder string            Specifies result set sort order sorted by the date due; if not specified, default sort order is descending (Desc) Valid values: Asc, Desc
      --statuses strings            Specifies the activity statuses to filter by Valid values: Planned, InProgress, Completed, InvalidSchedule, NotCompleted
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
      --types strings               Specifies the activity types. Informational, AssessedContent and Assessment are deprecated Valid values: Informational, Coaching, AssessedContent, Assessment, External, Native
      --userId strings              Specifies the list of user IDs to be queried, up to 100 user IDs. It searches for any relationship for the userId.
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

* [gc users development activities](gc_users_development_activities.html)	 - /api/v2/users/development/activities


