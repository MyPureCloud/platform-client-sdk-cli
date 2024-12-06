## gc quality evaluations query list

Queries Evaluations and returns a paged list

### Synopsis

Queries Evaluations and returns a paged list

```
gc quality evaluations query list [flags]
```

### Options

```
      --agentHasRead string              agent has the evaluation Valid values: true, false
      --agentTeamId string               team id of the agent
      --agentUserId string               user id of the agent
      --assigneeUserId string            assignee user id
  -a, --autopaginate                     Automatically paginate through the results stripping page information
      --conversationId string            conversationId specified
      --endTime string                   end time of the evaluation query
      --evaluationState strings          
      --evaluatorUserId string           evaluator user id
      --expand strings                   variable name requested by expand list
      --expandAnswerTotalScores string   get the total scores for evaluations. NOTE: The answers will only be populated if this parameter is set to true in the request. Valid values: true, false
      --filtercondition string           Filter list command output based on a given condition or regular expression
      --formContextId string             shared id between form versions
  -h, --help                             help for list
      --includeDeletedUsers delete       Allow returning an agent or evaluator user with a delete status. Defaults to false. Valid values: true, false
      --isReleased string                the evaluation has been released Valid values: true, false
      --maximum string                   the maximum number of results to return
      --pageNumber string                The page number requested (default "1")
      --pageSize string                  The total page size requested (default "25")
      --previousPage string              Previous page token
      --queueId string                   queue id
      --sortOrder string                 NOTE: Does not work when conversationId is supplied.
      --startTime string                 start time of the evaluation query
  -s, --stream                           Paginate and stream data as it is being processed leaving page information intact
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

* [gc quality evaluations query](gc_quality_evaluations_query.html)	 - /api/v2/quality/evaluations/query


