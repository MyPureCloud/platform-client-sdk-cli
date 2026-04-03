package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyschemaupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyschemaupdaterequestDud struct { 
    


    


    

}

// Journeyschemaupdaterequest - Update an external event schema
type Journeyschemaupdaterequest struct { 
    // Version - The schema's version, a positive integer.
    Version int `json:"version"`


    // Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
    Enabled bool `json:"enabled"`


    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Journeyjsonschemarequest `json:"jsonSchema"`

}

// String returns a JSON representation of the model
func (o *Journeyschemaupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyschemaupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Journeyschemaupdaterequest

    if JourneyschemaupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    JourneyschemaupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        JsonSchema Journeyjsonschemarequest `json:"jsonSchema"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

