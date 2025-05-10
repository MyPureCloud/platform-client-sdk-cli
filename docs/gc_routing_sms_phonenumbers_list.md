## gc routing sms phonenumbers list

Get a list of provisioned phone numbers.

### Synopsis

Get a list of provisioned phone numbers.

```
gc routing sms phonenumbers list [flags]
```

### Options

```
  -a, --autopaginate                Automatically paginate through the results stripping page information
      --countryCode strings         Filter on country code
      --expand strings              Which fields, if any, to expand Valid values: identityresolution, supportedContent
      --filtercondition string      Filter list command output based on a given condition or regular expression
  -h, --help                        help for list
      --integrationId string        Filter on the Genesys Cloud integration id to which the phone number belongs to
      --language string             A language tag (which is sometimes referred to as a locale identifier) to use to localize country field and sort operations
      --pageNumber string           Page number (default "1")
      --pageSize string             Page size (default "25")
      --phoneNumber 0-9             Filter on phone number address. Allowable characters are the digits 0-9 and the wild card character `\*`. If just digits are present, a contains search is done on the address pattern. For example, `317` could be matched anywhere in the address. An `\*` will match multiple digits. For example, to match a specific area code within the US a pattern like `1317*` could be used.
      --phoneNumberStatus strings   Filter on phone number status Valid values: active, invalid, initiated, porting, pending, pending-cancellation
      --phoneNumberType strings     Filter on phone number type Valid values: local, mobile, tollfree, shortcode, alphanumeric
      --sortBy string               Optional field to sort results Valid values: phoneNumber, countryCode, country, dateCreated, dateModified, phoneNumberStatus, phoneNumberType, purchaseDate, supportsMms, supportsSms, supportsVoice
      --sortOrder string            Sort order Valid values: ascending, descending
  -s, --stream                      Paginate and stream data as it is being processed leaving page information intact
      --supportedContentId string   Filter based on the supported content ID
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

* [gc routing sms phonenumbers](gc_routing_sms_phonenumbers.html)	 - /api/v2/routing/sms/phonenumbers


