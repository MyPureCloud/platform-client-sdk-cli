package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthproviderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthproviderDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Oauthprovider
type Oauthprovider struct { 
    


    // Name
    Name string `json:"name"`


    // Disabled
    Disabled bool `json:"disabled"`


    

}

// String returns a JSON representation of the model
func (o *Oauthprovider) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthprovider) MarshalJSON() ([]byte, error) {
    type Alias Oauthprovider

    if OauthproviderMarshalled {
        return []byte("{}"), nil
    }
    OauthproviderMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Disabled bool `json:"disabled"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

