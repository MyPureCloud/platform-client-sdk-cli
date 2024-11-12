## gc coaching appointments me list

Get my appointments for a given date range

### Synopsis

Get my appointments for a given date range

```
gc coaching appointments me list [flags]
```

### Options

```
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --completionInterval string   Appointment completion start and end to filter by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --facilitatorIds strings      The facilitator IDs for which to retrieve appointments
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --interval string             Interval to filter data by. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --intervalCondition string    Filter condition for interval Valid values: StartsIn, Overlaps
      --overdue string              Overdue status to filter by Valid values: True, False, Any
      --pageNumber string           Page number (default "1")
      --pageSize string             Page size (default "25")
      --relationships strings       Relationships to filter by Valid values: Creator, Facilitator, Attendee
      --sortOrder string            Sort (by due date) either Asc or Desc Valid values: Desc, Asc
      --statuses strings            Appointment Statuses to filter by Valid values: Scheduled, InProgress, Completed, InvalidSchedule
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
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

* [gc coaching appointments me](gc_coaching_appointments_me.html)	 - /api/v2/coaching/appointments/me


