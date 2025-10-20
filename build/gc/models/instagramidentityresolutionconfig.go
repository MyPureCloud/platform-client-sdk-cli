package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramidentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Instagramidentityresolutionconfig
type Instagramidentityresolutionconfig struct { 
    


    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Instagramidentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Instagramidentityresolutionconfig

    if InstagramidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    InstagramidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

