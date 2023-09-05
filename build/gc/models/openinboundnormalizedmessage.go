package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeninboundnormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeninboundnormalizedmessageDud struct { 
    


    


    


    

}

// Openinboundnormalizedmessage - Open Messaging rich media message structure
type Openinboundnormalizedmessage struct { 
    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openinboundmessagemessagingchannel `json:"channel"`


    // Text - Message text.
    Text string `json:"text"`


    // Content - List of content elements.
    Content []Openinboundmessagecontent `json:"content"`


    // Metadata - Additional metadata about this message to capture non-channel specific data.
    Metadata map[string]string `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Openinboundnormalizedmessage) String() string {
    
    
     o.Content = []Openinboundmessagecontent{{}} 
     o.Metadata = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openinboundnormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Openinboundnormalizedmessage

    if OpeninboundnormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    OpeninboundnormalizedmessageMarshalled = true

    return json.Marshal(&struct {
        
        Channel Openinboundmessagemessagingchannel `json:"channel"`
        
        Text string `json:"text"`
        
        Content []Openinboundmessagecontent `json:"content"`
        
        Metadata map[string]string `json:"metadata"`
        *Alias
    }{

        


        


        
        Content: []Openinboundmessagecontent{{}},
        


        
        Metadata: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

