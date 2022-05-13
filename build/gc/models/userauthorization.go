package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserauthorizationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserauthorizationDud struct { 
    


    UnusedRoles []Domainrole `json:"unusedRoles"`


    Permissions []string `json:"permissions"`


    PermissionPolicies []Resourcepermissionpolicy `json:"permissionPolicies"`

}

// Userauthorization
type Userauthorization struct { 
    // Roles
    Roles []Domainrole `json:"roles"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Userauthorization) String() string {
     o.Roles = []Domainrole{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userauthorization) MarshalJSON() ([]byte, error) {
    type Alias Userauthorization

    if UserauthorizationMarshalled {
        return []byte("{}"), nil
    }
    UserauthorizationMarshalled = true

    return json.Marshal(&struct {
        
        Roles []Domainrole `json:"roles"`
        *Alias
    }{

        
        Roles: []Domainrole{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

