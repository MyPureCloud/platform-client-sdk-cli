package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallmessageDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Callmessage
type Callmessage struct { 
    


    // Name
    Name string `json:"name"`


    // Message - raw SIP message
    Message string `json:"message"`


    

}

// String returns a JSON representation of the model
func (o *Callmessage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callmessage) MarshalJSON() ([]byte, error) {
    type Alias Callmessage

    if CallmessageMarshalled {
        return []byte("{}"), nil
    }
    CallmessageMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

