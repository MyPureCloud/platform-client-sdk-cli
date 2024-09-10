## gc outbound campaigns progress

/api/v2/outbound/campaigns/progress /api/v2/outbound/campaigns/{campaignId}/progress

### Synopsis

/api/v2/outbound/campaigns/progress /api/v2/outbound/campaigns/{campaignId}/progress

### Options

```
  -h, --help   help for progress
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

* [gc outbound campaigns](gc_outbound_campaigns.html)	 - /api/v2/outbound/campaigns
* [gc outbound campaigns progress create](gc_outbound_campaigns_progress_create.html)	 - Get progress for a list of campaigns
* [gc outbound campaigns progress delete](gc_outbound_campaigns_progress_delete.html)	 - Reset campaign progress and recycle the campaign
* [gc outbound campaigns progress get](gc_outbound_campaigns_progress_get.html)	 - Get campaign progress


