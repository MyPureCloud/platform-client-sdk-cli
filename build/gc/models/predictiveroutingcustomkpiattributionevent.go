package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictiveroutingcustomkpiattributioneventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictiveroutingcustomkpiattributioneventDud struct { 
    


    


    


    


    


    


    

}

// Predictiveroutingcustomkpiattributionevent
type Predictiveroutingcustomkpiattributionevent struct { 
    // EventId - A unique UUID v4 used to identify each attribution event uniquely
    EventId string `json:"eventId"`


    // EventDateTime - The date and time at which the attributed business result was achieved, in UTC, formatted as an ISO-8601 string. For example, 2023-01-07T08:30:00.123Z. Use the time the business result actually occurred or was confirmed, not a later ingestion, synchronisation, reporting, or batch-processing time. When conversationId is supplied, this timestamp must fall within the conversation attribution window, or the event is rejected. This means eventDateTime must be no earlier than the conversation start time, and no later than the conversation end time plus the grace period for the conversation's media type (1 hour for voice, message and email).
    EventDateTime time.Time `json:"eventDateTime"`


    // ExternalContactId - The id of the External Contact of the customer who was involved in the KPI attribution. If this field is populated, we will retrieve the relevant conversation internally - thus this field is not compatible with conversationId. You must chose to use either externalContactId or conversationId but not both on the same event.
    ExternalContactId string `json:"externalContactId"`


    // ConversationId - The id of the conversation. This field must be omitted if externalContactId is already populated on the event.
    ConversationId string `json:"conversationId"`


    // AgentId - The id of the agent who achieved the KPI attribution that is getting the attribution.
    AgentId string `json:"agentId"`


    // KpiId - The id of the relevant revenue KPI within the Genesys Platform - this id can be retrieved from the Revenue KPI Admin UI.
    KpiId string `json:"kpiId"`


    // AssociatedValue - For Sales Conversion, Retention and Churn KPI type, the numeric value expected is \"1\". For SalesValue KPI type, a numeric value representing the aggregated value of all KPI attributions achieved by an agent during a conversation, we do not allow multiple event uploads for the same conversationId and agentId pair.
    AssociatedValue float64 `json:"associatedValue"`

}

// String returns a JSON representation of the model
func (o *Predictiveroutingcustomkpiattributionevent) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictiveroutingcustomkpiattributionevent) MarshalJSON() ([]byte, error) {
    type Alias Predictiveroutingcustomkpiattributionevent

    if PredictiveroutingcustomkpiattributioneventMarshalled {
        return []byte("{}"), nil
    }
    PredictiveroutingcustomkpiattributioneventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ExternalContactId string `json:"externalContactId"`
        
        ConversationId string `json:"conversationId"`
        
        AgentId string `json:"agentId"`
        
        KpiId string `json:"kpiId"`
        
        AssociatedValue float64 `json:"associatedValue"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

