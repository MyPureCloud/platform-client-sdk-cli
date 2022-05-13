package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SmsaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SmsaddressDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    Validated bool `json:"validated"`


    SelfUri string `json:"selfUri"`

}

// Smsaddress
type Smsaddress struct { 
    


    // Name
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


    


    

}

// String returns a JSON representation of the model
func (o *Smsaddress) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Smsaddress) MarshalJSON() ([]byte, error) {
    type Alias Smsaddress

    if SmsaddressMarshalled {
        return []byte("{}"), nil
    }
    SmsaddressMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Street string `json:"street"`
        
        City string `json:"city"`
        
        Region string `json:"region"`
        
        PostalCode string `json:"postalCode"`
        
        CountryCode string `json:"countryCode"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

