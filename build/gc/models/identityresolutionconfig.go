package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IdentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IdentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Identityresolutionconfig
type Identityresolutionconfig struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Identityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Identityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Identityresolutionconfig

    if IdentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    IdentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

