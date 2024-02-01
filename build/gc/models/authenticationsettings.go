package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticationsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticationsettingsDud struct { 
    


    


    

}

// Authenticationsettings - Settings for authenticated webdeployments.
type Authenticationsettings struct { 
    // Enabled - Indicate if these auth is required for this deployment. If, for example, this flag is set to true then webmessaging sessions can not send messages unless the end-user is authenticated.
    Enabled bool `json:"enabled"`


    // IntegrationId - The integration identifier which contains the auth settings required on the deployment.
    IntegrationId string `json:"integrationId"`


    // AllowSessionUpgrade - Allow end-users to upgrade an anonymous session to authenticated conversation.
    AllowSessionUpgrade bool `json:"allowSessionUpgrade"`

}

// String returns a JSON representation of the model
func (o *Authenticationsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticationsettings) MarshalJSON() ([]byte, error) {
    type Alias Authenticationsettings

    if AuthenticationsettingsMarshalled {
        return []byte("{}"), nil
    }
    AuthenticationsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        IntegrationId string `json:"integrationId"`
        
        AllowSessionUpgrade bool `json:"allowSessionUpgrade"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

