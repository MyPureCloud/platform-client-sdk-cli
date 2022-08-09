package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationencryptionconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationencryptionconfigurationDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Conversationencryptionconfiguration
type Conversationencryptionconfiguration struct { 
    


    // Url - keyConfigurationType is always KmsSymmetric, and should be the arn to the key alias for the master key
    Url string `json:"url"`


    // KeyConfigurationType - Type should be 'KmsSymmetric' when create or update Key configurations, 'None' to disable configuration.
    KeyConfigurationType string `json:"keyConfigurationType"`


    // LastError - The error message related to the configuration
    LastError Errorbody `json:"lastError"`


    

}

// String returns a JSON representation of the model
func (o *Conversationencryptionconfiguration) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationencryptionconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Conversationencryptionconfiguration

    if ConversationencryptionconfigurationMarshalled {
        return []byte("{}"), nil
    }
    ConversationencryptionconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        
        KeyConfigurationType string `json:"keyConfigurationType"`
        
        LastError Errorbody `json:"lastError"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

