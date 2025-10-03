package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessrulesschemacreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessrulesschemacreaterequestDud struct { 
    

}

// Businessrulesschemacreaterequest - Request to create a new Business Rules Schema
type Businessrulesschemacreaterequest struct { 
    // JsonSchema - A JSON schema defining the extension to the built-in entity type.
    JsonSchema Jsonschemawithdefinitions `json:"jsonSchema"`

}

// String returns a JSON representation of the model
func (o *Businessrulesschemacreaterequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessrulesschemacreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Businessrulesschemacreaterequest

    if BusinessrulesschemacreaterequestMarshalled {
        return []byte("{}"), nil
    }
    BusinessrulesschemacreaterequestMarshalled = true

    return json.Marshal(&struct {
        
        JsonSchema Jsonschemawithdefinitions `json:"jsonSchema"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

