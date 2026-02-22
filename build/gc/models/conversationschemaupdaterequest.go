package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationschemaupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationschemaupdaterequestDud struct { 
    


    


    

}

// Conversationschemaupdaterequest - Schema update request.
type Conversationschemaupdaterequest struct { 
    // Version - The schema's version, a positive integer.
    Version int `json:"version"`


    // Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
    Enabled bool `json:"enabled"`


    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Conversationjsonschemarequest `json:"jsonSchema"`

}

// String returns a JSON representation of the model
func (o *Conversationschemaupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationschemaupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Conversationschemaupdaterequest

    if ConversationschemaupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    ConversationschemaupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        JsonSchema Conversationjsonschemarequest `json:"jsonSchema"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

