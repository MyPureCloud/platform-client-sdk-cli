package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsaddressprovisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsaddressprovisionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Smsaddressprovision
type Smsaddressprovision struct { 
    


    // Name - Name associated with this address
    Name string `json:"name"`


    // Street - The number and street address where this address is located.
    Street string `json:"street"`


    // City - The city in which this address is in
    City string `json:"city"`


    // Region - The state or region this address is in
    Region string `json:"region"`


    // PostalCode - The postal code this address is in
    PostalCode string `json:"postalCode"`


    // CountryCode - The ISO country code of this address
    CountryCode string `json:"countryCode"`


    // AutoCorrectAddress - This is used when the address is created. If the value is not set or true, then the system will, if necessary, auto-correct the address you provide. Set this value to false if the system should not auto-correct the address.
    AutoCorrectAddress bool `json:"autoCorrectAddress"`


    

}

// String returns a JSON representation of the model
func (o *Smsaddressprovision) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsaddressprovision) MarshalJSON() ([]byte, error) {
    type Alias Smsaddressprovision

    if SmsaddressprovisionMarshalled {
        return []byte("{}"), nil
    }
    SmsaddressprovisionMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Street string `json:"street"`
        
        City string `json:"city"`
        
        Region string `json:"region"`
        
        PostalCode string `json:"postalCode"`
        
        CountryCode string `json:"countryCode"`
        
        AutoCorrectAddress bool `json:"autoCorrectAddress"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

