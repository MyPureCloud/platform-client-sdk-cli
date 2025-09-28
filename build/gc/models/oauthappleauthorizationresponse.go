package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthappleauthorizationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthappleauthorizationresponseDud struct { 
    


    

}

// Oauthappleauthorizationresponse
type Oauthappleauthorizationresponse struct { 
    // RedirectUrl - The redirected url
    RedirectUrl string `json:"redirectUrl"`


    // VarError - The error object
    VarError Oauthappleauthorizationresponseerror `json:"error"`

}

// String returns a JSON representation of the model
func (o *Oauthappleauthorizationresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthappleauthorizationresponse) MarshalJSON() ([]byte, error) {
    type Alias Oauthappleauthorizationresponse

    if OauthappleauthorizationresponseMarshalled {
        return []byte("{}"), nil
    }
    OauthappleauthorizationresponseMarshalled = true

    return json.Marshal(&struct {
        
        RedirectUrl string `json:"redirectUrl"`
        
        VarError Oauthappleauthorizationresponseerror `json:"error"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

