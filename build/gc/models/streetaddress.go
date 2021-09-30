package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StreetaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StreetaddressDud struct { 
    


    


    


    


    


    


    


    

}

// Streetaddress
type Streetaddress struct { 
    // Country - 2 Letter Country code, like US or GB
    Country string `json:"country"`


    // A1 - State or Province
    A1 string `json:"A1"`


    // A3 - City or township
    A3 string `json:"A3"`


    // RD - Number and street
    RD string `json:"RD"`


    // HNO - House Number
    HNO string `json:"HNO"`


    // LOC - extra location info like suite 300
    LOC string `json:"LOC"`


    // NAM - Name of the customer
    NAM string `json:"NAM"`


    // PC - Postal code
    PC string `json:"PC"`

}

// String returns a JSON representation of the model
func (o *Streetaddress) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Streetaddress) MarshalJSON() ([]byte, error) {
    type Alias Streetaddress

    if StreetaddressMarshalled {
        return []byte("{}"), nil
    }
    StreetaddressMarshalled = true

    return json.Marshal(&struct { 
        Country string `json:"country"`
        
        A1 string `json:"A1"`
        
        A3 string `json:"A3"`
        
        RD string `json:"RD"`
        
        HNO string `json:"HNO"`
        
        LOC string `json:"LOC"`
        
        NAM string `json:"NAM"`
        
        PC string `json:"PC"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

