package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ErrordetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ErrordetailsDud struct { 
    


    


    


    


    


    


    


    

}

// Errordetails
type Errordetails struct { 
    // Status
    Status int `json:"status"`


    // Message
    Message string `json:"message"`


    // MessageWithParams
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams
    MessageParams map[string]string `json:"messageParams"`


    // Code
    Code string `json:"code"`


    // ContextId
    ContextId string `json:"contextId"`


    // Nested
    Nested *Errordetails `json:"nested"`


    // Details
    Details string `json:"details"`

}

// String returns a JSON representation of the model
func (o *Errordetails) String() string {
    
    
    
     o.MessageParams = map[string]string{"": ""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Errordetails) MarshalJSON() ([]byte, error) {
    type Alias Errordetails

    if ErrordetailsMarshalled {
        return []byte("{}"), nil
    }
    ErrordetailsMarshalled = true

    return json.Marshal(&struct {
        
        Status int `json:"status"`
        
        Message string `json:"message"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]string `json:"messageParams"`
        
        Code string `json:"code"`
        
        ContextId string `json:"contextId"`
        
        Nested *Errordetails `json:"nested"`
        
        Details string `json:"details"`
        *Alias
    }{

        


        


        


        
        MessageParams: map[string]string{"": ""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

