package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PrefixMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PrefixDud struct { 
    


    


    


    

}

// Prefix
type Prefix struct { 
    // CountryCode - The ITU-T E.164 country code (numeric, max 4 digits, required)
    CountryCode string `json:"countryCode"`


    // Number - The DID (Direct Inward Dialing) number (numeric, max 20 digits)
    Number string `json:"number"`


    // VarType - Prefix type: allow or block
    VarType string `json:"type"`


    // Action - The action to perform: ADD or DELETE
    Action string `json:"action"`

}

// String returns a JSON representation of the model
func (o *Prefix) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Prefix) MarshalJSON() ([]byte, error) {
    type Alias Prefix

    if PrefixMarshalled {
        return []byte("{}"), nil
    }
    PrefixMarshalled = true

    return json.Marshal(&struct {
        
        CountryCode string `json:"countryCode"`
        
        Number string `json:"number"`
        
        VarType string `json:"type"`
        
        Action string `json:"action"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

