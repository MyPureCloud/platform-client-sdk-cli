package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainorganizationroleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainorganizationroleDud struct { 
    


    


    


    


    


    UnusedPermissions []string `json:"unusedPermissions"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainorganizationrole
type Domainorganizationrole struct { 
    // Id - role id
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // DefaultRoleId
    DefaultRoleId string `json:"defaultRoleId"`


    // Permissions
    Permissions []string `json:"permissions"`


    


    // PermissionPolicies
    PermissionPolicies []Domainpermissionpolicy `json:"permissionPolicies"`


    // UserCount
    UserCount int `json:"userCount"`


    // RoleNeedsUpdate - Optional unless patch operation.
    RoleNeedsUpdate bool `json:"roleNeedsUpdate"`


    // Base
    Base bool `json:"base"`


    // VarDefault
    VarDefault bool `json:"default"`


    

}

// String returns a JSON representation of the model
func (o *Domainorganizationrole) String() string {
    
    
    
    
     o.Permissions = []string{""} 
     o.PermissionPolicies = []Domainpermissionpolicy{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainorganizationrole) MarshalJSON() ([]byte, error) {
    type Alias Domainorganizationrole

    if DomainorganizationroleMarshalled {
        return []byte("{}"), nil
    }
    DomainorganizationroleMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DefaultRoleId string `json:"defaultRoleId"`
        
        Permissions []string `json:"permissions"`
        
        PermissionPolicies []Domainpermissionpolicy `json:"permissionPolicies"`
        
        UserCount int `json:"userCount"`
        
        RoleNeedsUpdate bool `json:"roleNeedsUpdate"`
        
        Base bool `json:"base"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        


        


        
        Permissions: []string{""},
        


        


        
        PermissionPolicies: []Domainpermissionpolicy{{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

