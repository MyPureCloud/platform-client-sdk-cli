package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappavailablephonenumberdetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappavailablephonenumberdetailsDud struct { 
    Name string `json:"name"`


    PhoneNumber string `json:"phoneNumber"`


    Status string `json:"status"`

}

// Whatsappavailablephonenumberdetails
type Whatsappavailablephonenumberdetails struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Whatsappavailablephonenumberdetails) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappavailablephonenumberdetails) MarshalJSON() ([]byte, error) {
    type Alias Whatsappavailablephonenumberdetails

    if WhatsappavailablephonenumberdetailsMarshalled {
        return []byte("{}"), nil
    }
    WhatsappavailablephonenumberdetailsMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

