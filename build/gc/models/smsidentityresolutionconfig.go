package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsidentityresolutionconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsidentityresolutionconfigDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Smsidentityresolutionconfig
type Smsidentityresolutionconfig struct { 
    


    // Division - The division to use when performing identity resolution.
    Division Writablestarrabledivision `json:"division"`


    // ResolveIdentities - Whether the channel should resolve identities
    ResolveIdentities bool `json:"resolveIdentities"`


    

}

// String returns a JSON representation of the model
func (o *Smsidentityresolutionconfig) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsidentityresolutionconfig) MarshalJSON() ([]byte, error) {
    type Alias Smsidentityresolutionconfig

    if SmsidentityresolutionconfigMarshalled {
        return []byte("{}"), nil
    }
    SmsidentityresolutionconfigMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        ResolveIdentities bool `json:"resolveIdentities"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

