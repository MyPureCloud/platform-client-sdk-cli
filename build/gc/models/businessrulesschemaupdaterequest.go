package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessrulesschemaupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessrulesschemaupdaterequestDud struct { 
    


    


    

}

// Businessrulesschemaupdaterequest - Request to update an existing Business Rules Schema
type Businessrulesschemaupdaterequest struct { 
    // Version - The schema's version, a positive integer. Required for updates.
    Version int `json:"version"`


    // Enabled - The schema's enabled/disabled status. A disabled schema cannot be assigned to any other entities, but the data on those entities from the schema still exists.
    Enabled bool `json:"enabled"`


    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Jsonschemawithdefinitions `json:"jsonSchema"`

}

// String returns a JSON representation of the model
func (o *Businessrulesschemaupdaterequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessrulesschemaupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Businessrulesschemaupdaterequest

    if BusinessrulesschemaupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    BusinessrulesschemaupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        JsonSchema Jsonschemawithdefinitions `json:"jsonSchema"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

