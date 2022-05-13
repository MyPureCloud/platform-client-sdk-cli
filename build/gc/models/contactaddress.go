package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactaddressDud struct { 
    


    


    


    


    


    

}

// Contactaddress
type Contactaddress struct { 
    // Address1
    Address1 string `json:"address1"`


    // Address2
    Address2 string `json:"address2"`


    // City
    City string `json:"city"`


    // State
    State string `json:"state"`


    // PostalCode
    PostalCode string `json:"postalCode"`


    // CountryCode
    CountryCode string `json:"countryCode"`

}

// String returns a JSON representation of the model
func (o *Contactaddress) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactaddress) MarshalJSON() ([]byte, error) {
    type Alias Contactaddress

    if ContactaddressMarshalled {
        return []byte("{}"), nil
    }
    ContactaddressMarshalled = true

    return json.Marshal(&struct {
        
        Address1 string `json:"address1"`
        
        Address2 string `json:"address2"`
        
        City string `json:"city"`
        
        State string `json:"state"`
        
        PostalCode string `json:"postalCode"`
        
        CountryCode string `json:"countryCode"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

