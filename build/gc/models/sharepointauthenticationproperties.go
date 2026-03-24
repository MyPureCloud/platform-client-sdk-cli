package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SharepointauthenticationpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SharepointauthenticationpropertiesDud struct { 
    


    


    


    

}

// Sharepointauthenticationproperties
type Sharepointauthenticationproperties struct { 
    // AuthenticationUrl - The authentication URL for the connection.
    AuthenticationUrl string `json:"authenticationUrl"`


    // TenantId - The tenant ID for the connection.
    TenantId string `json:"tenantId"`


    // ClientId - The client ID for the connection.
    ClientId string `json:"clientId"`


    // RedirectUrl - The redirect URL for the connection.
    RedirectUrl string `json:"redirectUrl"`

}

// String returns a JSON representation of the model
func (o *Sharepointauthenticationproperties) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sharepointauthenticationproperties) MarshalJSON() ([]byte, error) {
    type Alias Sharepointauthenticationproperties

    if SharepointauthenticationpropertiesMarshalled {
        return []byte("{}"), nil
    }
    SharepointauthenticationpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        AuthenticationUrl string `json:"authenticationUrl"`
        
        TenantId string `json:"tenantId"`
        
        ClientId string `json:"clientId"`
        
        RedirectUrl string `json:"redirectUrl"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

