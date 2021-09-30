package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PermissiondetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PermissiondetailsDud struct { 
    


    


    


    

}

// Permissiondetails
type Permissiondetails struct { 
    // VarType - The type of permission requirement
    VarType string `json:"type"`


    // Permissions - List of required permissions
    Permissions []string `json:"permissions"`


    // AllowsCurrentUser - Whether the current user can subscribe, when division permissions are otherwise required
    AllowsCurrentUser bool `json:"allowsCurrentUser"`


    // Enforced - Whether or not this permission requirement is enforced
    Enforced bool `json:"enforced"`

}

// String returns a JSON representation of the model
func (o *Permissiondetails) String() string {
    
    
    
    
    
    
     o.Permissions = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Permissiondetails) MarshalJSON() ([]byte, error) {
    type Alias Permissiondetails

    if PermissiondetailsMarshalled {
        return []byte("{}"), nil
    }
    PermissiondetailsMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Permissions []string `json:"permissions"`
        
        AllowsCurrentUser bool `json:"allowsCurrentUser"`
        
        Enforced bool `json:"enforced"`
        
        *Alias
    }{
        

        

        

        
        Permissions: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

