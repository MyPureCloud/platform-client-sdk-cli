package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappidentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Whatsappidentityresolutionconfig
type Whatsappidentityresolutionconfig struct { 
    


    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Whatsappidentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Whatsappidentityresolutionconfig

    if WhatsappidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    WhatsappidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

