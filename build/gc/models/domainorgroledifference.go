package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainorgroledifferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainorgroledifferenceDud struct { 
    


    


    


    


    

}

// Domainorgroledifference
type Domainorgroledifference struct { 
    // RemovedPermissionPolicies
    RemovedPermissionPolicies []Domainpermissionpolicy `json:"removedPermissionPolicies"`


    // AddedPermissionPolicies
    AddedPermissionPolicies []Domainpermissionpolicy `json:"addedPermissionPolicies"`


    // SamePermissionPolicies
    SamePermissionPolicies []Domainpermissionpolicy `json:"samePermissionPolicies"`


    // UserOrgRole
    UserOrgRole Domainorganizationrole `json:"userOrgRole"`


    // RoleFromDefault
    RoleFromDefault Domainorganizationrole `json:"roleFromDefault"`

}

// String returns a JSON representation of the model
func (o *Domainorgroledifference) String() string {
     o.RemovedPermissionPolicies = []Domainpermissionpolicy{{}} 
     o.AddedPermissionPolicies = []Domainpermissionpolicy{{}} 
     o.SamePermissionPolicies = []Domainpermissionpolicy{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainorgroledifference) MarshalJSON() ([]byte, error) {
    type Alias Domainorgroledifference

    if DomainorgroledifferenceMarshalled {
        return []byte("{}"), nil
    }
    DomainorgroledifferenceMarshalled = true

    return json.Marshal(&struct {
        
        RemovedPermissionPolicies []Domainpermissionpolicy `json:"removedPermissionPolicies"`
        
        AddedPermissionPolicies []Domainpermissionpolicy `json:"addedPermissionPolicies"`
        
        SamePermissionPolicies []Domainpermissionpolicy `json:"samePermissionPolicies"`
        
        UserOrgRole Domainorganizationrole `json:"userOrgRole"`
        
        RoleFromDefault Domainorganizationrole `json:"roleFromDefault"`
        *Alias
    }{

        
        RemovedPermissionPolicies: []Domainpermissionpolicy{{}},
        


        
        AddedPermissionPolicies: []Domainpermissionpolicy{{}},
        


        
        SamePermissionPolicies: []Domainpermissionpolicy{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

