package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwittersignupoauthsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwittersignupoauthsettingsDud struct { 
    


    


    

}

// Twittersignupoauthsettings
type Twittersignupoauthsettings struct { 
    // ClientId - The client id of the twitter app the requesting org will use to signup
    ClientId string `json:"clientId"`


    // Scopes - The scopes/permissions requested during the signup process during the signup process to allow their future integrations to direct message
    Scopes []string `json:"scopes"`


    // AppId - The app id of the twitter app the requesting org will use to signup
    AppId string `json:"appId"`

}

// String returns a JSON representation of the model
func (o *Twittersignupoauthsettings) String() string {
    
     o.Scopes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twittersignupoauthsettings) MarshalJSON() ([]byte, error) {
    type Alias Twittersignupoauthsettings

    if TwittersignupoauthsettingsMarshalled {
        return []byte("{}"), nil
    }
    TwittersignupoauthsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ClientId string `json:"clientId"`
        
        Scopes []string `json:"scopes"`
        
        AppId string `json:"appId"`
        *Alias
    }{

        


        
        Scopes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

