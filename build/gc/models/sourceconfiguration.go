package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SourceconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SourceconfigurationDud struct { 
    


    


    

}

// Sourceconfiguration
type Sourceconfiguration struct { 
    // SourceId - Identifies the external platform that is the source of the conversation.
    SourceId string `json:"sourceId"`


    // InteractionId - The customer's unique external identifier associated with the conversation that comes from the external platform.
    InteractionId string `json:"interactionId"`


    // TagId - The customer's external identifier or tag associated with the conversation. If set, it will be used to tag the conversation.
    TagId string `json:"tagId"`

}

// String returns a JSON representation of the model
func (o *Sourceconfiguration) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sourceconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Sourceconfiguration

    if SourceconfigurationMarshalled {
        return []byte("{}"), nil
    }
    SourceconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        SourceId string `json:"sourceId"`
        
        InteractionId string `json:"interactionId"`
        
        TagId string `json:"tagId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

