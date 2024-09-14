package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpenmessagenormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpenmessagenormalizedmessageDud struct { 
    


    


    


    


    


    

}

// Openmessagenormalizedmessage - Open Messaging rich media message structure
type Openmessagenormalizedmessage struct { 
    // Id - Unique ID of the message generated by Messaging Platform.
    Id string `json:"id"`


    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Openmessagingchannel `json:"channel"`


    // VarType - Message type.
    VarType string `json:"type"`


    // Text - Message text.
    Text string `json:"text"`


    // Content - List of content elements.
    Content []Openmessagecontent `json:"content"`


    // Metadata - Additional metadata about this message.
    Metadata map[string]string `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Openmessagenormalizedmessage) String() string {
    
    
    
    
     o.Content = []Openmessagecontent{{}} 
     o.Metadata = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openmessagenormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Openmessagenormalizedmessage

    if OpenmessagenormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    OpenmessagenormalizedmessageMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Channel Openmessagingchannel `json:"channel"`
        
        VarType string `json:"type"`
        
        Text string `json:"text"`
        
        Content []Openmessagecontent `json:"content"`
        
        Metadata map[string]string `json:"metadata"`
        *Alias
    }{

        


        


        


        


        
        Content: []Openmessagecontent{{}},
        


        
        Metadata: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}
