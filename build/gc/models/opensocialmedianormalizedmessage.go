package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmedianormalizedmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmedianormalizedmessageDud struct { 
    Id string `json:"id"`


    


    


    VarType string `json:"type"`


    


    


    SelfUri string `json:"selfUri"`

}

// Opensocialmedianormalizedmessage - Open Social Messaging rich media message structure
type Opensocialmedianormalizedmessage struct { 
    


    // Channel - Channel-specific information that describes the message and the message channel/provider.
    Channel Opensocialmediachannel `json:"channel"`


    // Text - Message text.
    Text string `json:"text"`


    


    // Content - List of content elements.
    Content []Opensocialmediamessagecontent `json:"content"`


    // Metadata - Additional metadata about this message.
    Metadata map[string]string `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Opensocialmedianormalizedmessage) String() string {
    
    
     o.Content = []Opensocialmediamessagecontent{{}} 
     o.Metadata = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmedianormalizedmessage) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmedianormalizedmessage

    if OpensocialmedianormalizedmessageMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmedianormalizedmessageMarshalled = true

    return json.Marshal(&struct {
        
        Channel Opensocialmediachannel `json:"channel"`
        
        Text string `json:"text"`
        
        Content []Opensocialmediamessagecontent `json:"content"`
        
        Metadata map[string]string `json:"metadata"`
        *Alias
    }{

        


        


        


        


        
        Content: []Opensocialmediamessagecontent{{}},
        


        
        Metadata: map[string]string{"": ""},
        


        

        Alias: (*Alias)(u),
    })
}

