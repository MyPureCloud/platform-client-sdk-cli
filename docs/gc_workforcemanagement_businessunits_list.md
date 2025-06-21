## gc workforcemanagement businessunits list

Get business units

### Synopsis

Get business units

```
gc workforcemanagement businessunits list [flags]
```

### Options

```
      --divisionId string   If specified, the list of business units belonging to the specified division will be returned
      --feature string      If specified, the list of business units for which the user is authorized to use the requested feature will be returned Valid values: AgentHistoricalAdherence, AgentHistoricalAdherenceConformance, AgentSchedule, AgentTimeOffRequest, AgentWorkPlanBid, AlternativeShift, Coaching, Learning, AgentUnavailableTimes, ActivityCodes, ActivityPlans, UnavailableTimes, Agents, BuActivityCodes, BusinessUnits, CapacityPlan, ContinuousForecast, HistoricalAdherence, HistoricalShrinkage, IntradayMonitoring, BuIntradayMonitoring, ManagementUnits, RealTimeAdherence, Schedules, BuSchedules, ServiceGoalTemplates, PlanningGroups, LongTermStaffing, ShiftTrading, ShortTermForecasts, BuShortTermForecasts, StaffingGroups, TimeOffPlans, TimeOffRequests, TimeOffLimits, WorkPlanBids, WorkPlanBidGroups, WorkPlanRotations, WorkPlans
  -h, --help                help for list
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

* [gc workforcemanagement businessunits](gc_workforcemanagement_businessunits.html)	 - /api/v2/workforcemanagement/businessunits


