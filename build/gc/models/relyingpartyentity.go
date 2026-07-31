package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RelyingpartyentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RelyingpartyentityDud struct { 
    


    

}

// Relyingpartyentity
type Relyingpartyentity struct { 
    // Id - The relying party identifier (typically the registrable domain).
    Id string `json:"id"`


    // Name - The human-readable name of the relying party.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Relyingpartyentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Relyingpartyentity) MarshalJSON() ([]byte, error) {
    type Alias Relyingpartyentity

    if RelyingpartyentityMarshalled {
        return []byte("{}"), nil
    }
    RelyingpartyentityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

