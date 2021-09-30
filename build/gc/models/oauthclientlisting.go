package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthclientlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthclientlistingDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Oauthclientlisting
type Oauthclientlisting struct { 
    


    // Name - The name of the OAuth client.
    Name string `json:"name"`


    // AccessTokenValiditySeconds - The number of seconds, between 5mins and 48hrs, until tokens created with this client expire. If this field is omitted, a default of 24 hours will be applied.
    AccessTokenValiditySeconds int `json:"accessTokenValiditySeconds"`


    // Description
    Description string `json:"description"`


    // RegisteredRedirectUri - List of allowed callbacks for this client. For example: https://myap.example.com/auth/callback
    RegisteredRedirectUri []string `json:"registeredRedirectUri"`


    // Secret - System created secret assigned to this client. Secrets are required for code authorization and client credential grants.
    Secret string `json:"secret"`


    // RoleIds - Deprecated. Use roleDivisions instead.
    RoleIds []string `json:"roleIds"`


    // DateCreated - Date this client was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this client was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User that created this client
    CreatedBy Domainentityref `json:"createdBy"`


    // ModifiedBy - User that last modified this client
    ModifiedBy Domainentityref `json:"modifiedBy"`


    // Scope - The scope requested by this client. Scopes only apply to clients not using the client_credential grant
    Scope []string `json:"scope"`


    // RoleDivisions - Set of roles and their corresponding divisions associated with this client. Roles and divisions only apply to clients using the client_credential grant
    RoleDivisions []Roledivision `json:"roleDivisions"`


    // State - The state of the OAuth client. Active: The OAuth client can be used to create access tokens. This is the default state. Disabled: Access tokens created by the client are invalid and new ones cannot be created. Inactive: Access tokens cannot be created with this OAuth client and it will be deleted.
    State string `json:"state"`


    // DateToDelete - The time at which this client will be deleted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateToDelete time.Time `json:"dateToDelete"`


    

}

// String returns a JSON representation of the model
func (o *Oauthclientlisting) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RegisteredRedirectUri = []string{""} 
    
    
    
    
    
    
    
     o.RoleIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Scope = []string{""} 
    
    
    
     o.RoleDivisions = []Roledivision{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthclientlisting) MarshalJSON() ([]byte, error) {
    type Alias Oauthclientlisting

    if OauthclientlistingMarshalled {
        return []byte("{}"), nil
    }
    OauthclientlistingMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        AccessTokenValiditySeconds int `json:"accessTokenValiditySeconds"`
        
        Description string `json:"description"`
        
        RegisteredRedirectUri []string `json:"registeredRedirectUri"`
        
        Secret string `json:"secret"`
        
        RoleIds []string `json:"roleIds"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        ModifiedBy Domainentityref `json:"modifiedBy"`
        
        Scope []string `json:"scope"`
        
        RoleDivisions []Roledivision `json:"roleDivisions"`
        
        State string `json:"state"`
        
        DateToDelete time.Time `json:"dateToDelete"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        RegisteredRedirectUri: []string{""},
        

        

        

        

        
        RoleIds: []string{""},
        

        

        

        

        

        

        

        

        

        

        
        Scope: []string{""},
        

        

        
        RoleDivisions: []Roledivision{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

