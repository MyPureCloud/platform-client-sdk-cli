## gc outbound campaigns linedistribution get

Get line distribution information for campaigns using same Edge Group or Site as given campaign

### Synopsis

Get line distribution information for campaigns using same Edge Group or Site as given campaign

```
gc outbound campaigns linedistribution get [campaignId] [flags]
```

### Options

```
      --edgeGroupId s Edge Group. Campaign   Edge group to be used in line distribution calculations instead of current Campaigns Edge Group. Campaigns Site and Edge Group are mutually exclusive.
  -h, --help                                 help for get
      --includeOnlyActiveCampaigns string    If true will return only active Campaigns Valid values: true, false
      --outboundLineCount int                The number of outbound lines to be used in line distribution calculations, instead of current Campaign`s Outbound Lines Count
      --relativeWeight int                   Relative weight to be used in line distribution calculations instead of current Campaign`s relative weight
      --siteId s Site.  Campaign             Site to be used in line distribution calculations instead of current Campaigns Site.  Campaigns Site and Edge Group are mutually exclusive.
      --useWeight string                     Enable usage of weight, this value overrides current Campaign`s setting in line distribution calculations Valid values: true, false
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

* [gc outbound campaigns linedistribution](gc_outbound_campaigns_linedistribution.html)	 - /api/v2/outbound/campaigns/{campaignId}/linedistribution


