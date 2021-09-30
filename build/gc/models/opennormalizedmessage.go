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


    

}

// String returns a JSON representation of the model
func (o *Opennormalizedmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Content = []Openmessagecontent{{}} 
    
    
    
    

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
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Content: []Openmessagecontent{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

