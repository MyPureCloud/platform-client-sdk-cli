## gc learning assignments me list

List of Learning Assignments assigned to current user

### Synopsis

List of Learning Assignments assigned to current user

```
gc learning assignments me list [flags]
```

### Options

```
  -a, --autopaginate                 Automatically paginate through the results stripping page information
      --completionInterval string    Specifies the range of completion dates to be used for filtering. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --expand strings               Specifies the expand option for returning additional information Valid values: ModuleSummary
      --filtercondition string       Filter list command output based on a given condition or regular expression
  -h, --help                         help for list
      --interval string              Specifies the range of dueDates to be queried. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
      --maxPercentageScore float32   The maximum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)
      --minPercentageScore float32   The minimum assessment score for an assignment (completed with assessment) to be included in the results (inclusive)
      --moduleId string              Specifies the ID of the learning module. Fetch assignments for learning module ID
      --overdue string               Specifies if only the non-overdue (overdue is False) or overdue (overdue is True) assignments are returned. If overdue is Any or if the overdue parameter is not supplied, all assignments are returned Valid values: True, False, Any
      --pageNumber int               Page number (default 1)
      --pageSize int                 Page size (default 25)
      --pass string                  Specifies if only the failed (pass is False) or passed (pass is True) assignments (completed with assessment)are returned. If pass is Any or if the pass parameter is not supplied, all assignments are returned Valid values: True, False, Any
      --sortBy string                Specifies which field to sort the results by, default sort is by recommendedCompletionDate Valid values: RecommendedCompletionDate, DateModified
      --sortOrder string             Specifies result set sort order; if not specified, default sort order is descending (Desc) Valid values: Asc, Desc
      --states strings               Specifies the assignment states to filter by Valid values: Assigned, InProgress, Completed, NotCompleted, InvalidSchedule
  -s, --stream                       Paginate and stream data as it is being processed leaving page information intact
      --types strings                Specifies the module types to filter by. Informational, AssessedContent and Assessment are deprecated Valid values: Informational, AssessedContent, Assessment, External, Native
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

* [gc learning assignments me](gc_learning_assignments_me.html)	 - /api/v2/learning/assignments/me


