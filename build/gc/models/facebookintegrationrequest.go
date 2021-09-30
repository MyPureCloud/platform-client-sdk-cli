package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookintegrationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookintegrationrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Facebookintegrationrequest
type Facebookintegrationrequest struct { 
    


    // Name - The name of the Facebook Integration
    Name string `json:"name"`


    // PageAccessToken - The long-lived Page Access Token of a facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When a pageAccessToken is provided, pageId and userAccessToken are not required.
    PageAccessToken string `json:"pageAccessToken"`


    // UserAccessToken - The short-lived User Access Token of the facebook user logged into the facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  When userAccessToken is provided, pageId is mandatory.  When userAccessToken/pageId combination is provided, pageAccessToken is not required.
    UserAccessToken string `json:"userAccessToken"`


    // PageId - The page Id of a facebook page. The pageId is required when userAccessToken is provided.
    PageId string `json:"pageId"`


    // AppId - The app Id of a facebook app. The appId is required when a customer wants to use their own approved facebook app.
    AppId string `json:"appId"`


    // AppSecret - The app Secret of a facebook app. The appSecret is required when appId is provided.
    AppSecret string `json:"appSecret"`


    

}

// String returns a JSON representation of the model
func (o *Facebookintegrationrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookintegrationrequest) MarshalJSON() ([]byte, error) {
    type Alias Facebookintegrationrequest

    if FacebookintegrationrequestMarshalled {
        return []byte("{}"), nil
    }
    FacebookintegrationrequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        PageAccessToken string `json:"pageAccessToken"`
        
        UserAccessToken string `json:"userAccessToken"`
        
        PageId string `json:"pageId"`
        
        AppId string `json:"appId"`
        
        AppSecret string `json:"appSecret"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

