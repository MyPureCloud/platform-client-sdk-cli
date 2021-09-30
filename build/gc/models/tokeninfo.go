package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TokeninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TokeninfoDud struct { 
    Organization Namedentity `json:"organization"`


    HomeOrganization Namedentity `json:"homeOrganization"`


    AuthorizedScope []string `json:"authorizedScope"`


    ClonedUser Tokeninfocloneduser `json:"clonedUser"`


    

}

// Tokeninfo
type Tokeninfo struct { 
    


    


    


    


    // OAuthClient
    OAuthClient Orgoauthclient `json:"OAuthClient"`

}

// String returns a JSON representation of the model
func (o *Tokeninfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tokeninfo) MarshalJSON() ([]byte, error) {
    type Alias Tokeninfo

    if TokeninfoMarshalled {
        return []byte("{}"), nil
    }
    TokeninfoMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        OAuthClient Orgoauthclient `json:"OAuthClient"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

