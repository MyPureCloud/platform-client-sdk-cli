package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustgroupDud struct { 
    Id string `json:"id"`


    


    


    DateModified time.Time `json:"dateModified"`


    MemberCount int `json:"memberCount"`


    State string `json:"state"`


    Version int `json:"version"`


    


    


    


    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Orguser `json:"createdBy"`

}

// Trustgroup
type Trustgroup struct { 
    


    // Name - The group name.
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    


    


    


    


    // VarType - Type of group.
    VarType string `json:"type"`


    // Images
    Images []Userimage `json:"images"`


    // Addresses
    Addresses []Groupcontact `json:"addresses"`


    // RulesVisible - Are membership rules visible to the person requesting to view the group
    RulesVisible bool `json:"rulesVisible"`


    // Visibility - Who can view this group
    Visibility string `json:"visibility"`


    // RolesEnabled - Allow roles to be assigned to this group
    RolesEnabled bool `json:"rolesEnabled"`


    // IncludeOwners - Allow owners to be included as members of the group
    IncludeOwners bool `json:"includeOwners"`


    // Owners - Owners of the group
    Owners []User `json:"owners"`


    


    

}

// String returns a JSON representation of the model
func (o *Trustgroup) String() string {
    
    
    
     o.Images = []Userimage{{}} 
     o.Addresses = []Groupcontact{{}} 
    
    
    
    
     o.Owners = []User{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustgroup) MarshalJSON() ([]byte, error) {
    type Alias Trustgroup

    if TrustgroupMarshalled {
        return []byte("{}"), nil
    }
    TrustgroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Images []Userimage `json:"images"`
        
        Addresses []Groupcontact `json:"addresses"`
        
        RulesVisible bool `json:"rulesVisible"`
        
        Visibility string `json:"visibility"`
        
        RolesEnabled bool `json:"rolesEnabled"`
        
        IncludeOwners bool `json:"includeOwners"`
        
        Owners []User `json:"owners"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Images: []Userimage{{}},
        


        
        Addresses: []Groupcontact{{}},
        


        


        


        


        


        
        Owners: []User{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

