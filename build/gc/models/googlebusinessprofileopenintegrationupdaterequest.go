package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GooglebusinessprofileopenintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GooglebusinessprofileopenintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Googlebusinessprofileopenintegrationupdaterequest
type Googlebusinessprofileopenintegrationupdaterequest struct { 
    


    // Name - The name of the Google Business Profile Open Integration.
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // Token - Google OAuth 2 access token reference. The actual token cannot be accessed via Genesys API, only referenced by this property by its ID. When the token is not referenced by any integration, it is deleted after 24 hours.
    Token Googleauthtokenreference `json:"token"`


    // Account - Google Business Profile account accessible with the authorization token
    Account Googlebusinessprofileaccountreference `json:"account"`


    

}

// String returns a JSON representation of the model
func (o *Googlebusinessprofileopenintegrationupdaterequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Googlebusinessprofileopenintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Googlebusinessprofileopenintegrationupdaterequest

    if GooglebusinessprofileopenintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    GooglebusinessprofileopenintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        Token Googleauthtokenreference `json:"token"`
        
        Account Googlebusinessprofileaccountreference `json:"account"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

