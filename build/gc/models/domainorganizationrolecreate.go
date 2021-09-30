package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainorganizationrolecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainorganizationrolecreateDud struct { 
    Id string `json:"id"`


    


    


    


    


    UnusedPermissions []string `json:"unusedPermissions"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Domainorganizationrolecreate
type Domainorganizationrolecreate struct { 
    


    // Name - The role name
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


    // VarDefault
    VarDefault bool `json:"default"`


    // Base
    Base bool `json:"base"`


    

}

// String returns a JSON representation of the model
func (o *Domainorganizationrolecreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Permissions = []string{""} 
    
    
    
    
    
     o.PermissionPolicies = []Domainpermissionpolicy{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainorganizationrolecreate) MarshalJSON() ([]byte, error) {
    type Alias Domainorganizationrolecreate

    if DomainorganizationrolecreateMarshalled {
        return []byte("{}"), nil
    }
    DomainorganizationrolecreateMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DefaultRoleId string `json:"defaultRoleId"`
        
        Permissions []string `json:"permissions"`
        
        
        
        PermissionPolicies []Domainpermissionpolicy `json:"permissionPolicies"`
        
        UserCount int `json:"userCount"`
        
        RoleNeedsUpdate bool `json:"roleNeedsUpdate"`
        
        VarDefault bool `json:"default"`
        
        Base bool `json:"base"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Permissions: []string{""},
        

        

        

        

        
        PermissionPolicies: []Domainpermissionpolicy{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

