package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FacebookintegrationupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FacebookintegrationupdaterequestDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Facebookintegrationupdaterequest
type Facebookintegrationupdaterequest struct { 
    


    // Name - The name of the Facebook Integration
    Name string `json:"name"`


    // PageAccessToken - The long-lived Page Access Token of a facebook page.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
    PageAccessToken string `json:"pageAccessToken"`


    // UserAccessToken - The short-lived User Access Token of the facebook user logged into the facebook app.  See https://developers.facebook.com/docs/facebook-login/access-tokens.  Either pageAccessToken or userAccessToken should be provided.
    UserAccessToken string `json:"userAccessToken"`


    

}

// String returns a JSON representation of the model
func (o *Facebookintegrationupdaterequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Facebookintegrationupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Facebookintegrationupdaterequest

    if FacebookintegrationupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    FacebookintegrationupdaterequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        PageAccessToken string `json:"pageAccessToken"`
        
        UserAccessToken string `json:"userAccessToken"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

