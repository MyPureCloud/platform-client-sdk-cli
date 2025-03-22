package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TwitterintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TwitterintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Twitterintegrationupdaterequest
type Twitterintegrationupdaterequest struct { 
    


    // Name - The name of the Twitter messaging integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // SignupCode - The authorization code returned from the signup flow to request access tokens later on
    SignupCode string `json:"signupCode"`


    // CodeChallenge - The codeChallenge used during the signup flow
    CodeChallenge string `json:"codeChallenge"`


    // RedirectUri - The redirectUri used during the signup flow
    RedirectUri string `json:"redirectUri"`


    

}

// String returns a JSON representation of the model
func (o *Twitterintegrationupdaterequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Twitterintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Twitterintegrationupdaterequest

    if TwitterintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    TwitterintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        SignupCode string `json:"signupCode"`
        
        CodeChallenge string `json:"codeChallenge"`
        
        RedirectUri string `json:"redirectUri"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

