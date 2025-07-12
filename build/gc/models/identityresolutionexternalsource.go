package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentityresolutionexternalsourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentityresolutionexternalsourceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Identityresolutionexternalsource
type Identityresolutionexternalsource struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Identityresolutionexternalsource) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identityresolutionexternalsource) MarshalJSON() ([]byte, error) {
    type Alias Identityresolutionexternalsource

    if IdentityresolutionexternalsourceMarshalled {
        return []byte("{}"), nil
    }
    IdentityresolutionexternalsourceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

