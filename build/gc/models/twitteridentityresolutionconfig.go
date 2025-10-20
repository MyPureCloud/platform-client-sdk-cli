package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitteridentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitteridentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Twitteridentityresolutionconfig
type Twitteridentityresolutionconfig struct { 
    


    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Twitteridentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitteridentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Twitteridentityresolutionconfig

    if TwitteridentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    TwitteridentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

