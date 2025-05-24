package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationpushproviderintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationpushproviderintegrationDud struct { 
    


    

}

// Conversationpushproviderintegration - A Push provider integration.
type Conversationpushproviderintegration struct { 
    // Id - Genesys Cloud Integration ID
    Id string `json:"id"`


    // Provider - Type of the integration
    Provider string `json:"provider"`

}

// String returns a JSON representation of the model
func (o *Conversationpushproviderintegration) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationpushproviderintegration) MarshalJSON() ([]byte, error) {
    type Alias Conversationpushproviderintegration

    if ConversationpushproviderintegrationMarshalled {
        return []byte("{}"), nil
    }
    ConversationpushproviderintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Provider string `json:"provider"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

