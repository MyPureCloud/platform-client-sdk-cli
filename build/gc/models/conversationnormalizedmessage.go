package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationnormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationnormalizedmessageDud struct { 
    Id string `json:"id"`


    Channel Conversationmessagingchannel `json:"channel"`


    


    


    


    


    Status string `json:"status"`


    Reasons []Conversationreason `json:"reasons"`


    


    IsFinalReceipt bool `json:"isFinalReceipt"`


    Direction string `json:"direction"`


    


    


    


    

}

// Conversationnormalizedmessage - General rich media message structure with normalized feature support across many messaging channels.
type Conversationnormalizedmessage struct { 
    


    


    // VarType - Message type.
    VarType string `json:"type"`


    // Text - Message text.
    Text string `json:"text"`


    // Content - List of content elements.
    Content []Conversationmessagecontent `json:"content"`


    // Events - List of event elements.
    Events []Conversationmessageevent `json:"events"`


    


    


    // OriginatingEntity - Specifies if this message was sent by a human agent or bot. The platform may use this to apply appropriate provider policies.
    OriginatingEntity string `json:"originatingEntity"`


    


    


    // RelatedMessages - A list of messages related to this one.
    RelatedMessages []Conversationnormalizedmessage `json:"relatedMessages"`


    // Metadata - Additional metadata about this message.
    Metadata map[string]string `json:"metadata"`


    // Enrichment - Metadata enrichments provided by the platform.
    Enrichment Conversationenrichment `json:"enrichment"`


    // ByoSmsIntegrationId - The internal id representing the customer supplied sms integration message.
    ByoSmsIntegrationId string `json:"byoSmsIntegrationId"`

}

// String returns a JSON representation of the model
func (o *Conversationnormalizedmessage) String() string {
    
    
     o.Content = []Conversationmessagecontent{{}} 
     o.Events = []Conversationmessageevent{{}} 
    
     o.RelatedMessages = []Conversationnormalizedmessage{{}} 
     o.Metadata = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationnormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Conversationnormalizedmessage

    if ConversationnormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    ConversationnormalizedmessageMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Content []Conversationmessagecontent `json:"content"`
        
        Events []Conversationmessageevent `json:"events"`
        
        OriginatingEntity string `json:"originatingEntity"`
        
        RelatedMessages []Conversationnormalizedmessage `json:"relatedMessages"`
        
        Metadata map[string]string `json:"metadata"`
        
        Enrichment Conversationenrichment `json:"enrichment"`
        
        ByoSmsIntegrationId string `json:"byoSmsIntegrationId"`
        *Alias
    }{

        


        


        


        


        
        Content: []Conversationmessagecontent{{}},
        


        
        Events: []Conversationmessageevent{{}},
        


        


        


        


        


        


        
        RelatedMessages: []Conversationnormalizedmessage{{}},
        


        
        Metadata: map[string]string{"": ""},
        


        


        

        Alias: (*Alias)(u),
    })
}

