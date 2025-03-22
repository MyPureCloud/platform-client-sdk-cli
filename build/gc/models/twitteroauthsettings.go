package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitteroauthsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitteroauthsettingsDud struct { 
    


    

}

// Twitteroauthsettings
type Twitteroauthsettings struct { 
    // ClientId - The client id of the twitter app the requesting org will use to signup
    ClientId string `json:"clientId"`


    // Scopes - The scopes/permissions requested during the signup process during the signup process to allow their future integrations to direct message
    Scopes []string `json:"scopes"`

}

// String returns a JSON representation of the model
func (o *Twitteroauthsettings) String() string {
    
     o.Scopes = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitteroauthsettings) MarshalJSON() ([]byte, error) {
    type Alias Twitteroauthsettings

    if TwitteroauthsettingsMarshalled {
        return []byte("{}"), nil
    }
    TwitteroauthsettingsMarshalled = true

    return json.Marshal(&struct {
        
        ClientId string `json:"clientId"`
        
        Scopes []string `json:"scopes"`
        *Alias
    }{

        


        
        Scopes: []string{""},
        

        Alias: (*Alias)(u),
    })
}

