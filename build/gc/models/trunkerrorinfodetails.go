package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkerrorinfodetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkerrorinfodetailsDud struct { 
    


    


    

}

// Trunkerrorinfodetails
type Trunkerrorinfodetails struct { 
    // Code
    Code string `json:"code"`


    // Message
    Message string `json:"message"`


    // Hostname
    Hostname string `json:"hostname"`

}

// String returns a JSON representation of the model
func (o *Trunkerrorinfodetails) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunkerrorinfodetails) MarshalJSON() ([]byte, error) {
    type Alias Trunkerrorinfodetails

    if TrunkerrorinfodetailsMarshalled {
        return []byte("{}"), nil
    }
    TrunkerrorinfodetailsMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Hostname string `json:"hostname"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

