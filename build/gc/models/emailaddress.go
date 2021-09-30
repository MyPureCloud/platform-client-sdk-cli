package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailaddressDud struct { 
    


    

}

// Emailaddress
type Emailaddress struct { 
    // Email
    Email string `json:"email"`


    // Name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Emailaddress) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailaddress) MarshalJSON() ([]byte, error) {
    type Alias Emailaddress

    if EmailaddressMarshalled {
        return []byte("{}"), nil
    }
    EmailaddressMarshalled = true

    return json.Marshal(&struct { 
        Email string `json:"email"`
        
        Name string `json:"name"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

