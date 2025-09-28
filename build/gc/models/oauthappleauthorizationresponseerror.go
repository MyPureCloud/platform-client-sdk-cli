package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthappleauthorizationresponseerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthappleauthorizationresponseerrorDud struct { 
    


    


    

}

// Oauthappleauthorizationresponseerror
type Oauthappleauthorizationresponseerror struct { 
    // Code - The error code
    Code string `json:"code"`


    // Message - The error message
    Message string `json:"message"`


    // Details - The error details
    Details interface{} `json:"details"`

}

// String returns a JSON representation of the model
func (o *Oauthappleauthorizationresponseerror) String() string {
    
    
     o.Details = Interface{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthappleauthorizationresponseerror) MarshalJSON() ([]byte, error) {
    type Alias Oauthappleauthorizationresponseerror

    if OauthappleauthorizationresponseerrorMarshalled {
        return []byte("{}"), nil
    }
    OauthappleauthorizationresponseerrorMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Details interface{} `json:"details"`
        *Alias
    }{

        


        


        
        Details: Interface{},
        

        Alias: (*Alias)(u),
    })
}

