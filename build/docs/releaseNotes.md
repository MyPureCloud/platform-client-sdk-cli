Platform API version: 4631


# Major Changes (12 changes)

**PATCH /api/v2/routing/queues/{queueId}/members/{memberId}** (1 change)

* Response 200 was removed

**PATCH /api/v2/routing/queues/{queueId}/users/{memberId}** (1 change)

* Response 200 was removed

**Relationship** (1 change)

* Property name was removed

**WfmHistoricalAdherenceQuery** (1 change)

* Property teamIds was removed

**EntityListing** (5 changes)

* Property pageSize was removed
* Property pageNumber was removed
* Property total was removed
* Property pageCount was removed
* Property entities was changed from DataTableImportJob[] to object[]

**AnalyticsResolution** (1 change)

* Property getnNextContactAvoided was removed

**AnalyticsSession** (1 change)

* Property journeyActionMapVersion was changed from string to integer

**BuQueryAgentSchedulesRequest** (1 change)

* Property teamIds was removed


# Minor Changes (59 changes)

**/api/v2/authorization/subjects/{subjectId}/bulkreplace** (2 changes)

* Path was added
* Operation POST was added

**AuditQueryExecutionStatusResponse** (1 change)

* Enum value Integrations was added to property serviceName

**AuditLogMessage** (2 changes)

* Enum value Integrations was added to property serviceName
* Enum value Integration was added to property entityType

**AuditQueryRequest** (1 change)

* Enum value Integrations was added to property serviceName

**AuditRealtimeQueryRequest** (1 change)

* Enum value Integrations was added to property serviceName

**AuditQueryEntity** (1 change)

* Enum value Integration was added to property name

**AuditQueryService** (1 change)

* Enum value Integrations was added to property name

**ReportingExportJobResponse** (3 changes)

* Enum value BOT_PERFORMANCE_SUMMARY_VIEW was added to property viewType
* Enum value BOT_PERFORMANCE_DETAIL_VIEW was added to property viewType
* Enum value SCHEDULED_EXPORTS_VIEW was added to property viewType

**ViewFilter** (4 changes)

* Optional property customerSentimentScore was added
* Optional property customerSentimentTrend was added
* Optional property flowTransferTargets was added
* Optional property developmentName was added

**EventMessage** (1 change)

* Enum value INVALID_AGENT was added to property code

**Trustee** (1 change)

* Optional property usesDefaultRole was added

**AnalyticsConversationWithoutAttributes** (1 change)

* Optional property externalTag was added

**AnalyticsFlow** (1 change)

* Optional property recognitionFailureReason was added

**AnalyticsMediaEndpointStat** (1 change)

* Optional property eventTime was added

**AnalyticsResolution** (2 changes)

* Optional property eventTime was added
* Optional property nNextContactAvoided was added

**AnalyticsSession** (4 changes)

* Optional property agentBullseyeRing was added
* Optional property authenticated was added
* Values are no longer constrained by enum members
* Optional property routingRing was added

**PostTextResponse** (1 change)

* Optional property genesysBotConnector was added

**ConversationDetailQueryPredicate** (1 change)

* Enum value externalTag was added to property dimension

**SegmentDetailQueryPredicate** (1 change)

* Enum value authenticated was added to property dimension

**Dependency** (2 changes)

* Enum value LEXV2BOT was added to property type
* Enum value LEXV2BOTALIAS was added to property type

**DependencyObject** (2 changes)

* Enum value LEXV2BOT was added to property type
* Enum value LEXV2BOTALIAS was added to property type

**ReportingExportJobRequest** (3 changes)

* Enum value BOT_PERFORMANCE_SUMMARY_VIEW was added to property viewType
* Enum value BOT_PERFORMANCE_DETAIL_VIEW was added to property viewType
* Enum value SCHEDULED_EXPORTS_VIEW was added to property viewType

**ReportingExportMetadataJobResponse** (3 changes)

* Enum value BOT_PERFORMANCE_SUMMARY_VIEW was added to property viewType
* Enum value BOT_PERFORMANCE_DETAIL_VIEW was added to property viewType
* Enum value SCHEDULED_EXPORTS_VIEW was added to property viewType

**FlowAggregateQueryPredicate** (5 changes)

* Enum value agentBullseyeRing was added to property dimension
* Enum value authenticated was added to property dimension
* Enum value externalTag was added to property dimension
* Enum value recognitionFailureReason was added to property dimension
* Enum value routingRing was added to property dimension

**FlowAggregationQuery** (5 changes)

* Enum value agentBullseyeRing was added to property groupBy
* Enum value authenticated was added to property groupBy
* Enum value externalTag was added to property groupBy
* Enum value recognitionFailureReason was added to property groupBy
* Enum value routingRing was added to property groupBy

**ConversationAggregateQueryPredicate** (4 changes)

* Enum value agentBullseyeRing was added to property dimension
* Enum value authenticated was added to property dimension
* Enum value externalTag was added to property dimension
* Enum value routingRing was added to property dimension

**ConversationAggregationQuery** (4 changes)

* Enum value agentBullseyeRing was added to property groupBy
* Enum value authenticated was added to property groupBy
* Enum value externalTag was added to property groupBy
* Enum value routingRing was added to property groupBy

**AnalyticsConversation** (1 change)

* Optional property externalTag was added


# Point Changes (3 changes)

**POST /api/v2/outbound/dnclists/{dncListId}/phonenumbers** (1 change)

* Summary was changed

**GET /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media/{mediaId}** (1 change)

* Description was changed

**POST /api/v2/conversations/messages/{conversationId}/communications/{communicationId}/messages/media** (1 change)

* Description was changed
