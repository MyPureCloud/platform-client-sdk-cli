package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthauthorizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthauthorizationDud struct { 
    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Oauthauthorization
type Oauthauthorization struct { 
    // Client
    Client Oauthclient `json:"client"`


    // Scope
    Scope []string `json:"scope"`


    // ResourceOwner
    ResourceOwner Domainentityref `json:"resourceOwner"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // Pending
    Pending bool `json:"pending"`


    

}

// String returns a JSON representation of the model
func (o *Oauthauthorization) String() string {
    
    
    
    
    
    
     o.Scope = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthauthorization) MarshalJSON() ([]byte, error) {
    type Alias Oauthauthorization

    if OauthauthorizationMarshalled {
        return []byte("{}"), nil
    }
    OauthauthorizationMarshalled = true

    return json.Marshal(&struct { 
        Client Oauthclient `json:"client"`
        
        Scope []string `json:"scope"`
        
        ResourceOwner Domainentityref `json:"resourceOwner"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        Pending bool `json:"pending"`
        
        
        
        *Alias
    }{
        

        

        

        
        Scope: []string{""},
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

