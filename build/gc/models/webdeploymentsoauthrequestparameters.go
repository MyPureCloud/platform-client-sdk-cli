package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebdeploymentsoauthrequestparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebdeploymentsoauthrequestparametersDud struct { 
    


    


    


    


    


    

}

// Webdeploymentsoauthrequestparameters
type Webdeploymentsoauthrequestparameters struct { 
    // Code - The authorization code to be sent to the authentication server during the token request.  Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
    Code string `json:"code"`


    // RedirectUri - Redirect URI sent in the \"Authentication Request\"Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
    RedirectUri string `json:"redirectUri"`


    // Nonce - Required if provided in the \"Authentication Request\". Otherwise should be empty.String value used to associate a Client session with an ID Token, and to mitigate replay attacks. The value is passed through unmodified from the Authentication Request to the ID Token. Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
    Nonce string `json:"nonce"`


    // MaxAge - Required if provided in the  \"Authentication Request\". Otherwise should be empty.Specifies the allowable elapsed time in seconds since the last time the End-User was actively authenticated.Refer to https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest
    MaxAge int `json:"maxAge"`


    // CodeVerifier - Required if authorizing using Proof Key for Code Exchange (PKCE). Otherwise should be empty.Random URL-safe string with a minimum length of 43 characters generated at start of authorization flow to mitigate the threat of having the authorization code intercepted. Refer to https://datatracker.ietf.org/doc/html/rfc7636
    CodeVerifier string `json:"codeVerifier"`


    // Iss - Optional parameter. Set it if authorization server discovery metadata authorization_response_iss_parameter_supported is enabled. Refer to https://datatracker.ietf.org/doc/html/rfc9207
    Iss string `json:"iss"`

}

// String returns a JSON representation of the model
func (o *Webdeploymentsoauthrequestparameters) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webdeploymentsoauthrequestparameters) MarshalJSON() ([]byte, error) {
    type Alias Webdeploymentsoauthrequestparameters

    if WebdeploymentsoauthrequestparametersMarshalled {
        return []byte("{}"), nil
    }
    WebdeploymentsoauthrequestparametersMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        RedirectUri string `json:"redirectUri"`
        
        Nonce string `json:"nonce"`
        
        MaxAge int `json:"maxAge"`
        
        CodeVerifier string `json:"codeVerifier"`
        
        Iss string `json:"iss"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

