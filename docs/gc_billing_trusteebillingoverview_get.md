## gc billing trusteebillingoverview get

Get the billing overview for an organization that is managed by a partner.

### Synopsis

Get the billing overview for an organization that is managed by a partner.

```
gc billing trusteebillingoverview get [trustorOrgId] [flags]
```

### Options

```
      --billingPeriodIndex string   0 for active period (overview data may change until period closes). 1 for prior completed billing period. 2 for two billing cycles prior, and so on. (default "0")
  -h, --help                        help for get
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

* [gc billing trusteebillingoverview](gc_billing_trusteebillingoverview.html)	 - /api/v2/billing/trusteebillingoverview


