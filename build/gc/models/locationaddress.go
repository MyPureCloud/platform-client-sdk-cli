package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationaddressDud struct { 
    


    


    


    


    


    


    

}

// Locationaddress
type Locationaddress struct { 
    // City
    City string `json:"city"`


    // Country
    Country string `json:"country"`


    // CountryName
    CountryName string `json:"countryName"`


    // State
    State string `json:"state"`


    // Street1
    Street1 string `json:"street1"`


    // Street2
    Street2 string `json:"street2"`


    // Zipcode
    Zipcode string `json:"zipcode"`

}

// String returns a JSON representation of the model
func (o *Locationaddress) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationaddress) MarshalJSON() ([]byte, error) {
    type Alias Locationaddress

    if LocationaddressMarshalled {
        return []byte("{}"), nil
    }
    LocationaddressMarshalled = true

    return json.Marshal(&struct { 
        City string `json:"city"`
        
        Country string `json:"country"`
        
        CountryName string `json:"countryName"`
        
        State string `json:"state"`
        
        Street1 string `json:"street1"`
        
        Street2 string `json:"street2"`
        
        Zipcode string `json:"zipcode"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

