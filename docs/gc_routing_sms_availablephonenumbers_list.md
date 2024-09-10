## gc routing sms availablephonenumbers list

Get a list of available phone numbers for SMS provisioning.

### Synopsis

Get a list of available phone numbers for SMS provisioning.

```
gc routing sms availablephonenumbers list [flags]
```

### Options

```
      --addressRequirement string   This indicates whether the phone number requires to have an Address registered. Valid values: none, any, local, foreign
      --areaCode string             Area code that can be used to restrict the numbers returned
      --city string                 City that can be used to restrict the numbers returned
      --countryCode string          The ISO 3166-1 alpha-2 country code of the county for which available phone numbers should be returned - REQUIRED
  -h, --help                        help for list
      --pattern *                   A pattern to match phone numbers. Valid characters are * and [0-9a-zA-Z]. The `*` character will match any single digit.
      --phoneNumberType string      Type of available phone numbers searched - REQUIRED Valid values: local, mobile, tollfree
      --region string               Region/province/state that can be used to restrict the numbers returned
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

* [gc routing sms availablephonenumbers](gc_routing_sms_availablephonenumbers.html)	 - /api/v2/routing/sms/availablephonenumbers


