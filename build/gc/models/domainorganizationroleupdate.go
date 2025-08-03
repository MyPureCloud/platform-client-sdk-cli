package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainorganizationroleupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainorganizationroleupdateDud struct { 
    


    


    


    


    


    UnusedPermissions []string `json:"unusedPermissions"`


    


    


    


    


    


    DateLicenseLastUpdated time.Time `json:"dateLicenseLastUpdated"`


    


    


    SelfUri string `json:"selfUri"`

}

// Domainorganizationroleupdate
type Domainorganizationroleupdate struct { 
    // Id - role id
    Id string `json:"id"`


    // Name - The name of the role
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


    // BaseLicense
    BaseLicense string `json:"baseLicense"`


    // AddonLicenses
    AddonLicenses []string `json:"addonLicenses"`


    


    // VarDefault
    VarDefault bool `json:"default"`


    // Base
    Base bool `json:"base"`


    

}

// String returns a JSON representation of the model
func (o *Domainorganizationroleupdate) String() string {
    
    
    
    
     o.Permissions = []string{""} 
     o.PermissionPolicies = []Domainpermissionpolicy{{}} 
    
    
    
     o.AddonLicenses = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainorganizationroleupdate) MarshalJSON() ([]byte, error) {
    type Alias Domainorganizationroleupdate

    if DomainorganizationroleupdateMarshalled {
        return []byte("{}"), nil
    }
    DomainorganizationroleupdateMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DefaultRoleId string `json:"defaultRoleId"`
        
        Permissions []string `json:"permissions"`
        
        PermissionPolicies []Domainpermissionpolicy `json:"permissionPolicies"`
        
        UserCount int `json:"userCount"`
        
        RoleNeedsUpdate bool `json:"roleNeedsUpdate"`
        
        BaseLicense string `json:"baseLicense"`
        
        AddonLicenses []string `json:"addonLicenses"`
        
        VarDefault bool `json:"default"`
        
        Base bool `json:"base"`
        *Alias
    }{

        


        


        


        


        
        Permissions: []string{""},
        


        


        
        PermissionPolicies: []Domainpermissionpolicy{{}},
        


        


        


        


        
        AddonLicenses: []string{""},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

