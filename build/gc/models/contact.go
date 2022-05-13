package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactDud struct { 
    


    Display string `json:"display"`


    


    


    


    


    

}

// Contact
type Contact struct { 
    // Address - Email address or phone number for this contact type
    Address string `json:"address"`


    


    // MediaType
    MediaType string `json:"mediaType"`


    // VarType
    VarType string `json:"type"`


    // Extension - Use internal extension instead of address. Mutually exclusive with the address field.
    Extension string `json:"extension"`


    // CountryCode
    CountryCode string `json:"countryCode"`


    // Integration - Integration tag value if this number is associated with an external integration.
    Integration string `json:"integration"`

}

// String returns a JSON representation of the model
func (o *Contact) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contact) MarshalJSON() ([]byte, error) {
    type Alias Contact

    if ContactMarshalled {
        return []byte("{}"), nil
    }
    ContactMarshalled = true

    return json.Marshal(&struct {
        
        Address string `json:"address"`
        
        MediaType string `json:"mediaType"`
        
        VarType string `json:"type"`
        
        Extension string `json:"extension"`
        
        CountryCode string `json:"countryCode"`
        
        Integration string `json:"integration"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

