package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EngineintegrationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EngineintegrationDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Engineintegration
type Engineintegration struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Engineintegration) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Engineintegration) MarshalJSON() ([]byte, error) {
    type Alias Engineintegration

    if EngineintegrationMarshalled {
        return []byte("{}"), nil
    }
    EngineintegrationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

