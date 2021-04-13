Platform API version: 4600


Adding configurable log path and ability to switch logging on and off
Adding the ability to add a description to commands

# Major Changes (21 changes)

**GET /api/v2/knowledge/knowledgebases/{knowledgeBaseId}/languages/{languageCode}/documents** (1 change)

* Parameter documentIds was added

**ViewFilter** (1 change)

* Property developmentKeyType was removed

**BusinessUnitListItem** (1 change)

* Property division was changed from Division to DivisionReference

**EntityListing** (1 change)

* Property entities was changed from object[] to DataTableImportJob[]

**BusinessUnit** (1 change)

* Property division was changed from Division to DivisionReference

**ManagementUnit** (1 change)

* Property division was changed from Division to DivisionReference

**AnalyticsSession** (3 changes)

* Property bullseyeRing was removed
* Property agentBullseyeRing was removed
* Property routingRule was removed

**FlowAggregateQueryPredicate** (3 changes)

* Enum value agentBullseyeRing was removed from property dimension
* Enum value bullseyeRing was removed from property dimension
* Enum value routingRule was removed from property dimension

**FlowAggregationQuery** (3 changes)

* Enum value agentBullseyeRing was removed from property groupBy
* Enum value bullseyeRing was removed from property groupBy
* Enum value routingRule was removed from property groupBy

**ConversationAggregateQueryPredicate** (3 changes)

* Enum value agentBullseyeRing was removed from property dimension
* Enum value bullseyeRing was removed from property dimension
* Enum value routingRule was removed from property dimension

**ConversationAggregationQuery** (3 changes)

* Enum value agentBullseyeRing was removed from property groupBy
* Enum value bullseyeRing was removed from property groupBy
* Enum value routingRule was removed from property groupBy


# Minor Changes (8 changes)

**GDPRSubject** (1 change)

* Optional property externalId was added

**DivisionReference** (1 change)

* Model was added

**EntityListing** (4 changes)

* Optional property pageSize was added
* Optional property pageNumber was added
* Optional property total was added
* Optional property pageCount was added

**DevelopmentActivity** (2 changes)

* Enum value AssessedContent was added to property type
* Enum value Questionnaire was added to property type


# Point Changes (1 change)

**DELETE /api/v2/tokens/me** (1 change)

* Summary was changed
