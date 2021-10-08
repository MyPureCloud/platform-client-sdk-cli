package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpennormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpennormalizedmessageDud struct { 
    Id string `json:"id"`


    


    


    


    


    Status string `json:"status"`


    Reasons []Reason `json:"reasons"`


    IsFinalReceipt bool `json:"isFinalReceipt"`


    Direction string `json:"direction"`


    

}

// Opennormalizedmessage - Open Messaging rich media message structure
type Opennormalizedmessage struct { 
    


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
func (o *Opennormalizedmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Content = []Openmessagecontent{{}} 
    
    
    
    
    
    
    
    
    
    
    
     o.Metadata = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opennormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Opennormalizedmessage

    if OpennormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    OpennormalizedmessageMarshalled = true

    return json.Marshal(&struct { 
        
        
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

