package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenstructuredresponsenormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenstructuredresponsenormalizedmessageDud struct { 
    


    


    


    


    


    

}

// Openstructuredresponsenormalizedmessage - Open Messaging rich media response normalized message structure
type Openstructuredresponsenormalizedmessage struct { 
    // Id - Unique ID of the message. This ID is generated by Messaging Platform. Message receipts will have the same ID as the message they reference, as such should only be set when sending a message receipt.
    Id string `json:"id"`


    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openmessagingchannel `json:"channel"`


    // VarType - Message type.
    VarType string `json:"type"`


    // Direction - The direction of the message.
    Direction string `json:"direction"`


    // Content - List of content elements.
    Content []Openinboundstructuredresponsenormalizedmessagecontent `json:"content"`


    // Metadata - Additional metadata about this message.
    Metadata map[string]string `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Openstructuredresponsenormalizedmessage) String() string {
    
    
    
    
     o.Content = []Openinboundstructuredresponsenormalizedmessagecontent{{}} 
     o.Metadata = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openstructuredresponsenormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Openstructuredresponsenormalizedmessage

    if OpenstructuredresponsenormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    OpenstructuredresponsenormalizedmessageMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Channel Openmessagingchannel `json:"channel"`
        
        VarType string `json:"type"`
        
        Direction string `json:"direction"`
        
        Content []Openinboundstructuredresponsenormalizedmessagecontent `json:"content"`
        
        Metadata map[string]string `json:"metadata"`
        *Alias
    }{

        


        


        


        


        
        Content: []Openinboundstructuredresponsenormalizedmessagecontent{{}},
        


        
        Metadata: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

