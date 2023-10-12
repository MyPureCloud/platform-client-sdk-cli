package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NuancebotcredentialsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NuancebotcredentialsDud struct { 
    


    


    


    

}

// Nuancebotcredentials - Model for a Nuance bot credentials
type Nuancebotcredentials struct { 
    // AppId - The application ID
    AppId string `json:"appId"`


    // ClientId - The credentials client ID
    ClientId string `json:"clientId"`


    // ClientSecret - The credentials client secret
    ClientSecret string `json:"clientSecret"`


    // ClientSecretProvided - True if the credentials secret is set (but not returned due to security reasons)
    ClientSecretProvided bool `json:"clientSecretProvided"`

}

// String returns a JSON representation of the model
func (o *Nuancebotcredentials) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nuancebotcredentials) MarshalJSON() ([]byte, error) {
    type Alias Nuancebotcredentials

    if NuancebotcredentialsMarshalled {
        return []byte("{}"), nil
    }
    NuancebotcredentialsMarshalled = true

    return json.Marshal(&struct {
        
        AppId string `json:"appId"`
        
        ClientId string `json:"clientId"`
        
        ClientSecret string `json:"clientSecret"`
        
        ClientSecretProvided bool `json:"clientSecretProvided"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

