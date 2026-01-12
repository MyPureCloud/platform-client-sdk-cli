## gc

gc is a CLI for interacting with Genesys Cloud

### Synopsis

gc is a CLI for interacting with Genesys Cloud

### Options

```
      --accesstoken string    accessToken override
      --clientid string       clientId override
      --clientsecret string   clientSecret override
      --environment string    environment override. E.g. mypurecloud.com.au or ap-southeast-2
  -h, --help                  help for gc
  -i, --indicateprogress      Trace progress indicators to stderr
      --inputformat string    Data input format. Supported formats: YAML, JSON
      --outputformat string   Data output format. Supported formats: YAML, JSON
  -p, --profile string        Name of the profile to use for configuring the cli (default "DEFAULT")
      --transform string      Provide a Go template file for transforming output data
      --transformstr string   Provide a Go template string for transforming output data
```

### SEE ALSO

* [gc alerting](gc_alerting.html)	 - /api/v2/alerting
* [gc alternativeformats](gc_alternativeformats.html)	 - Used to specify the desired input and output formats
* [gc analytics](gc_analytics.html)	 - /api/v2/analytics
* [gc architect](gc_architect.html)	 - /api/v2/architect
* [gc assistants](gc_assistants.html)	 - /api/v2/assistants
* [gc audits](gc_audits.html)	 - /api/v2/audits
* [gc authorization](gc_authorization.html)	 - /api/v2/authorization
* [gc autopagination](gc_autopagination.html)	 - autopagination
* [gc billing](gc_billing.html)	 - /api/v2/billing
* [gc businessrules](gc_businessrules.html)	 - /api/v2/businessrules
* [gc carrierservices](gc_carrierservices.html)	 - /api/v2/carrierservices
* [gc certificate](gc_certificate.html)	 - /api/v2/certificate
* [gc chats](gc_chats.html)	 - /api/v2/chats
* [gc coaching](gc_coaching.html)	 - /api/v2/coaching
* [gc completion](gc_completion.html)	 - Generate completion script
* [gc contentmanagement](gc_contentmanagement.html)	 - /api/v2/contentmanagement
* [gc conversations](gc_conversations.html)	 - /api/v2/conversations
* [gc dataextensions](gc_dataextensions.html)	 - /api/v2/dataextensions
* [gc dataprivacy](gc_dataprivacy.html)	 - /api/v2/dataprivacy
* [gc date](gc_date.html)	 - /api/v2/date
* [gc diagnostics](gc_diagnostics.html)	 - /api/v2/diagnostics
* [gc documentation](gc_documentation.html)	 - /api/v2/documentation
* [gc downloads](gc_downloads.html)	 - /api/v2/downloads
* [gc emails](gc_emails.html)	 - /api/v2/emails
* [gc employeeengagement](gc_employeeengagement.html)	 - /api/v2/employeeengagement
* [gc employeeperformance](gc_employeeperformance.html)	 - /api/v2/employeeperformance
* [gc events](gc_events.html)	 - /api/v2/events
* [gc experimental](gc_experimental.html)	 - Manages the experimental features for the CLI
* [gc externalcontacts](gc_externalcontacts.html)	 - /api/v2/externalcontacts
* [gc fax](gc_fax.html)	 - /api/v2/fax
* [gc fieldconfig](gc_fieldconfig.html)	 - /api/v2/fieldconfig
* [gc flows](gc_flows.html)	 - /api/v2/flows
* [gc gamification](gc_gamification.html)	 - /api/v2/gamification
* [gc gateway](gc_gateway.html)	 - Manages the gateway for the CLI
* [gc gdpr](gc_gdpr.html)	 - /api/v2/gdpr
* [gc geolocations](gc_geolocations.html)	 - /api/v2/geolocations
* [gc greetings](gc_greetings.html)	 - /api/v2/greetings
* [gc groups](gc_groups.html)	 - /api/v2/groups
* [gc guides](gc_guides.html)	 - /api/v2/guides
* [gc identityproviders](gc_identityproviders.html)	 - /api/v2/identityproviders
* [gc infrastructureascode](gc_infrastructureascode.html)	 - /api/v2/infrastructureascode
* [gc integrations](gc_integrations.html)	 - /api/v2/integrations
* [gc intents](gc_intents.html)	 - /api/v2/intents
* [gc ipranges](gc_ipranges.html)	 - /api/v2/ipranges
* [gc journey](gc_journey.html)	 - /api/v2/journey
* [gc knowledge](gc_knowledge.html)	 - /api/v2/knowledge
* [gc languages](gc_languages.html)	 - /api/v2/languages
* [gc languageunderstanding](gc_languageunderstanding.html)	 - /api/v2/languageunderstanding
* [gc learning](gc_learning.html)	 - /api/v2/learning
* [gc license](gc_license.html)	 - /api/v2/license
* [gc locations](gc_locations.html)	 - /api/v2/locations
* [gc logging](gc_logging.html)	 - Manages the logging for the CLI
* [gc messaging](gc_messaging.html)	 - /api/v2/messaging
* [gc mobiledevices](gc_mobiledevices.html)	 - /api/v2/mobiledevices
* [gc notifications](gc_notifications.html)	 - /api/v2/notifications
* [gc oauth](gc_oauth.html)	 - /api/v2/oauth
* [gc organizations](gc_organizations.html)	 - /api/v2/organizations
* [gc orgauthorization](gc_orgauthorization.html)	 - /api/v2/orgauthorization
* [gc orphanrecordings](gc_orphanrecordings.html)	 - /api/v2/orphanrecordings
* [gc outbound](gc_outbound.html)	 - /api/v2/outbound
* [gc presence](gc_presence.html)	 - /api/v2/presence
* [gc presencedefinitions](gc_presencedefinitions.html)	 - /api/v2/presencedefinitions
* [gc processautomation](gc_processautomation.html)	 - /api/v2/processautomation
* [gc profile](gc_profile.html)	 - /api/v2/profiles
* [gc profiles](gc_profiles.html)	 - Manages the application profiles
* [gc proxy](gc_proxy.html)	 - Manages the proxy for the CLI
* [gc quality](gc_quality.html)	 - /api/v2/quality
* [gc recording](gc_recording.html)	 - /api/v2/recording
* [gc recordings](gc_recordings.html)	 - /api/v2/recordings
* [gc responsemanagement](gc_responsemanagement.html)	 - /api/v2/responsemanagement
* [gc routing](gc_routing.html)	 - /api/v2/routing
* [gc scim](gc_scim.html)	 - /api/v2/scim
* [gc screenrecording](gc_screenrecording.html)	 - /api/v2/screenrecording
* [gc scripts](gc_scripts.html)	 - /api/v2/scripts
* [gc search](gc_search.html)	 - /api/v2/search
* [gc settings](gc_settings.html)	 - /api/v2/settings
* [gc socialmedia](gc_socialmedia.html)	 - /api/v2/socialmedia
* [gc speechandtextanalytics](gc_speechandtextanalytics.html)	 - /api/v2/speechandtextanalytics
* [gc stations](gc_stations.html)	 - /api/v2/stations
* [gc systempresences](gc_systempresences.html)	 - /api/v2/systempresences
* [gc taskmanagement](gc_taskmanagement.html)	 - /api/v2/taskmanagement
* [gc teams](gc_teams.html)	 - /api/v2/teams
* [gc telephony](gc_telephony.html)	 - /api/v2/telephony
* [gc textbots](gc_textbots.html)	 - /api/v2/textbots
* [gc timezones](gc_timezones.html)	 - /api/v2/timezones
* [gc tokens](gc_tokens.html)	 - /api/v2/tokens
* [gc uploads](gc_uploads.html)	 - /api/v2/uploads
* [gc usage](gc_usage.html)	 - /api/v2/usage
* [gc userrecordings](gc_userrecordings.html)	 - /api/v2/userrecordings
* [gc users](gc_users.html)	 - /api/v2/users
* [gc version](gc_version.html)	 - Print the version number of gc
* [gc voicemail](gc_voicemail.html)	 - /api/v2/voicemail
* [gc webchat](gc_webchat.html)	 - /api/v2/webchat
* [gc webdeployments](gc_webdeployments.html)	 - /api/v2/webdeployments
* [gc webmessaging](gc_webmessaging.html)	 - /api/v2/webmessaging
* [gc widgets](gc_widgets.html)	 - /api/v2/widgets
* [gc workforcemanagement](gc_workforcemanagement.html)	 - /api/v2/workforcemanagement


