## gc knowledge knowledgebases list

Get knowledge bases

### Synopsis

Get knowledge bases

```
gc knowledge knowledgebases list [flags]
```

### Options

```
      --after string             The cursor that points to the end of the set of entities that has been returned.
  -a, --autopaginate             Automatically paginate through the results stripping page information
      --before string            The cursor that points to the start of the set of entities that has been returned.
      --coreLanguage string      Filter by core language. Valid values: en-US, en-UK, en-AU, en-CA, en-HK, en-IN, en-IE, en-NZ, en-PH, en-SG, en-ZA, de-DE, de-AT, de-CH, es-AR, es-CO, es-MX, es-US, es-ES, fr-FR, fr-BE, fr-CA, fr-CH, pt-BR, pt-PT, nl-NL, nl-BE, it-IT, ca-ES, tr-TR, sv-SE, fi-FI, nb-NO, da-DK, ja-JP, ar-AE, zh-CN, zh-TW, zh-HK, ko-KR, pl-PL, hi-IN, th-TH, hu-HU, vi-VN, uk-UA, cs-CZ
      --filtercondition string   Filter list command output based on a given condition or regular expression
  -h, --help                     help for list
      --limit string             Number of entities to return. Maximum of 100. Deprecated in favour of pageSize
      --name string              Filter by Name.
      --pageSize string          Number of entities to return. Maximum of 100.
      --published string         Filter by published status. Valid values: true, false
      --sortBy string            Sort by. Valid values: Name, Date
      --sortOrder string         Sort Order. Valid values: ASC, ascending, DESC, descending
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

* [gc knowledge knowledgebases](gc_knowledge_knowledgebases.html)	 - /api/v2/knowledge/knowledgebases


