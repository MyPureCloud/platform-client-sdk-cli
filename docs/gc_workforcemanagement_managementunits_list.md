## gc workforcemanagement managementunits list

Get management units

### Synopsis

Get management units

```
gc workforcemanagement managementunits list [flags]
```

### Options

```
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --divisionId string        If specified, the list of management units belonging to the specified division will be returned
      --expand string            Deprecated, expand settings on the single MU route Valid values: details
      --feature string           If specified, the list of management units for which the user is authorized to use the requested feature will be returned Valid values: AgentSchedule, AgentTimeOffRequest, AgentWorkPlanBid, AlternativeShift, Coaching, Learning, ActivityCodes, ActivityPlans, Agents, BuActivityCodes, BusinessUnits, ContinuousForecast, HistoricalAdherence, HistoricalShrinkage, IntradayMonitoring, BuIntradayMonitoring, ManagementUnits, RealTimeAdherence, Schedules, BuSchedules, ServiceGoalTemplates, PlanningGroups, ShiftTrading, ShortTermForecasts, BuShortTermForecasts, StaffingGroups, TimeOffPlans, TimeOffRequests, TimeOffLimits, WorkPlanBids, WorkPlanBidGroups, WorkPlanRotations, WorkPlans
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --pageNumber int           Deprecated, paging is not supported
      --pageSize int             Deprecated, paging is not supported
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

* [gc workforcemanagement managementunits](gc_workforcemanagement_managementunits.html)	 - /api/v2/workforcemanagement/managementunits


