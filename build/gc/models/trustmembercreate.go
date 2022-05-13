package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustmembercreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustmembercreateDud struct { 
    


    


    

}

// Trustmembercreate
type Trustmembercreate struct { 
    // Id - Trustee User or Group Id
    Id string `json:"id"`


    // RoleIds - The list of roles to be granted to this user or group. Roles will be granted in all divisions.
    RoleIds []string `json:"roleIds"`


    // RoleDivisions - The list of trustor organization roles granting this user or group access paired with the divisions for those roles.
    RoleDivisions Roledivisiongrants `json:"roleDivisions"`

}

// String returns a JSON representation of the model
func (o *Trustmembercreate) String() string {
    
     o.RoleIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustmembercreate) MarshalJSON() ([]byte, error) {
    type Alias Trustmembercreate

    if TrustmembercreateMarshalled {
        return []byte("{}"), nil
    }
    TrustmembercreateMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        RoleIds []string `json:"roleIds"`
        
        RoleDivisions Roledivisiongrants `json:"roleDivisions"`
        *Alias
    }{

        


        
        RoleIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

