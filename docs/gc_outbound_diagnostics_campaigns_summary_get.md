## gc outbound diagnostics campaigns summary get

Get diagnostic summary for a single campaign

### Synopsis

Get diagnostic summary for a single campaign

```
gc outbound diagnostics campaigns summary get [campaignId] [flags]
```

### Options

```
      --end string     End datetime (ISO 8601 or Unix epoch) - REQUIRED
  -h, --help           help for get
      --start string   Start datetime (ISO 8601 or Unix epoch) - REQUIRED
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

* [gc outbound diagnostics campaigns summary](gc_outbound_diagnostics_campaigns_summary.html)	 - /api/v2/outbound/diagnostics/campaigns/{campaignId}/summary


