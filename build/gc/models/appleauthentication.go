package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AppleauthenticationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AppleauthenticationDud struct { 
    


    


    


    

}

// Appleauthentication
type Appleauthentication struct { 
    // OauthClientId - The Apple Messages for Business OAuth client id.
    OauthClientId string `json:"oauthClientId"`


    // OauthClientSecret - The Apple Messages for Business OAuth client secret.
    OauthClientSecret string `json:"oauthClientSecret"`


    // OauthTokenUrl - The Apple Messages for Business token URL.
    OauthTokenUrl string `json:"oauthTokenUrl"`


    // OauthScope - The Apple Messages for Business OAuth scope.
    OauthScope string `json:"oauthScope"`

}

// String returns a JSON representation of the model
func (o *Appleauthentication) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Appleauthentication) MarshalJSON() ([]byte, error) {
    type Alias Appleauthentication

    if AppleauthenticationMarshalled {
        return []byte("{}"), nil
    }
    AppleauthenticationMarshalled = true

    return json.Marshal(&struct {
        
        OauthClientId string `json:"oauthClientId"`
        
        OauthClientSecret string `json:"oauthClientSecret"`
        
        OauthTokenUrl string `json:"oauthTokenUrl"`
        
        OauthScope string `json:"oauthScope"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

