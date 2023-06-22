package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InstagramintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InstagramintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Instagramintegrationupdaterequest
type Instagramintegrationupdaterequest struct { 
    


    // Name - The name of the Instagram Integration
    Name string `json:"name"`


    // SupportedContent - Defines the SupportedContent profile configured for an integration
    SupportedContent Supportedcontentreference `json:"supportedContent"`


    // MessagingSetting - Defines the message settings to be applied for this integration
    MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`


    // PageAccessToken - The long-lived Page Access Token of Instagram page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
    PageAccessToken string `json:"pageAccessToken"`


    // UserAccessToken - The short-lived User Access Token of the Instagram user logged into the Facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
    UserAccessToken string `json:"userAccessToken"`


    

}

// String returns a JSON representation of the model
func (o *Instagramintegrationupdaterequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Instagramintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Instagramintegrationupdaterequest

    if InstagramintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    InstagramintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        SupportedContent Supportedcontentreference `json:"supportedContent"`
        
        MessagingSetting Messagingsettingrequestreference `json:"messagingSetting"`
        
        PageAccessToken string `json:"pageAccessToken"`
        
        UserAccessToken string `json:"userAccessToken"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

