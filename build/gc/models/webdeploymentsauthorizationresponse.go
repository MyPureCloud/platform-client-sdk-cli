package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentsauthorizationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentsauthorizationresponseDud struct { 
    


    

}

// Webdeploymentsauthorizationresponse
type Webdeploymentsauthorizationresponse struct { 
    // RefreshToken - Refresh token used to issue a new JWT.
    RefreshToken string `json:"refreshToken"`


    // Jwt
    Jwt string `json:"jwt"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsauthorizationresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentsauthorizationresponse) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentsauthorizationresponse

    if WebdeploymentsauthorizationresponseMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentsauthorizationresponseMarshalled = true

    return json.Marshal(&struct {
        
        RefreshToken string `json:"refreshToken"`
        
        Jwt string `json:"jwt"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

