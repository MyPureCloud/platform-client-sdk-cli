package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Twitterintegrationrequest
type Twitterintegrationrequest struct { 
    


    // Name - The name of the Twitter Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // SignupCode - The authorization code returned from the signup flow to request access tokens later on
    SignupCode string `json:"signupCode"`


    // AppId - The appId of the Twitter app to register the integration on
    AppId string `json:"appId"`


    // CodeChallenge - The codeChallenge used during the signup flow
    CodeChallenge string `json:"codeChallenge"`


    // RedirectUri - The redirectUri used during the signup flow
    RedirectUri string `json:"redirectUri"`


    

}

// String returns a JSON representation of the model
func (o *Twitterintegrationrequest) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Twitterintegrationrequest

    if TwitterintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    TwitterintegrationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        SignupCode string `json:"signupCode"`
        
        AppId string `json:"appId"`
        
        CodeChallenge string `json:"codeChallenge"`
        
        RedirectUri string `json:"redirectUri"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

