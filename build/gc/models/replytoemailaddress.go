package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReplytoemailaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReplytoemailaddressDud struct { 
    


    

}

// Replytoemailaddress
type Replytoemailaddress struct { 
    // Domain - The InboundDomain used for the email address.
    Domain Domainentityref `json:"domain"`


    // Route - The InboundRoute used for the email address.
    Route Domainentityref `json:"route"`

}

// String returns a JSON representation of the model
func (o *Replytoemailaddress) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Replytoemailaddress) MarshalJSON() ([]byte, error) {
    type Alias Replytoemailaddress

    if ReplytoemailaddressMarshalled {
        return []byte("{}"), nil
    }
    ReplytoemailaddressMarshalled = true

    return json.Marshal(&struct { 
        Domain Domainentityref `json:"domain"`
        
        Route Domainentityref `json:"route"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

