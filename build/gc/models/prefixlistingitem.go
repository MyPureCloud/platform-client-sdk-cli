package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PrefixlistingitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PrefixlistingitemDud struct { 
    


    


    

}

// Prefixlistingitem
type Prefixlistingitem struct { 
    // CountryCode - The ITU-T E.164 country code (numeric, max 4 digits)
    CountryCode string `json:"countryCode"`


    // Number - The DID (Direct Inward Dialing) number (numeric, max 20 digits)
    Number string `json:"number"`


    // VarType - Prefix type: allow or block
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Prefixlistingitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Prefixlistingitem) MarshalJSON() ([]byte, error) {
    type Alias Prefixlistingitem

    if PrefixlistingitemMarshalled {
        return []byte("{}"), nil
    }
    PrefixlistingitemMarshalled = true

    return json.Marshal(&struct {
        
        CountryCode string `json:"countryCode"`
        
        Number string `json:"number"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

