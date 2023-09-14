package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealtherrorinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealtherrorinfoDud struct { 
    


    


    


    

}

// Flowhealtherrorinfo
type Flowhealtherrorinfo struct { 
    // Message
    Message string `json:"message"`


    // Code
    Code string `json:"code"`


    // MessageWithParams - Error message with params included.
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams - Map of variables and params for the error message.
    MessageParams map[string]interface{} `json:"messageParams"`

}

// String returns a JSON representation of the model
func (o *Flowhealtherrorinfo) String() string {
    
    
    
     o.MessageParams = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealtherrorinfo) MarshalJSON() ([]byte, error) {
    type Alias Flowhealtherrorinfo

    if FlowhealtherrorinfoMarshalled {
        return []byte("{}"), nil
    }
    FlowhealtherrorinfoMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]interface{} `json:"messageParams"`
        *Alias
    }{

        


        


        


        
        MessageParams: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

