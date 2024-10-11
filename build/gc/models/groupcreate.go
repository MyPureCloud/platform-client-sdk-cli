package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupcreateDud struct { 
    Id string `json:"id"`


    


    


    DateModified time.Time `json:"dateModified"`


    MemberCount int `json:"memberCount"`


    State string `json:"state"`


    Version int `json:"version"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Groupcreate
type Groupcreate struct { 
    


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


    // CallsEnabled - Allow calls to be placed to this group.
    CallsEnabled bool `json:"callsEnabled"`


    // OwnerIds - Owners of the group
    OwnerIds []string `json:"ownerIds"`


    

}

// String returns a JSON representation of the model
func (o *Groupcreate) String() string {
    
    
    
     o.Images = []Userimage{{}} 
     o.Addresses = []Groupcontact{{}} 
    
    
    
    
    
     o.OwnerIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupcreate) MarshalJSON() ([]byte, error) {
    type Alias Groupcreate

    if GroupcreateMarshalled {
        return []byte("{}"), nil
    }
    GroupcreateMarshalled = true

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
        
        CallsEnabled bool `json:"callsEnabled"`
        
        OwnerIds []string `json:"ownerIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Images: []Userimage{{}},
        


        
        Addresses: []Groupcontact{{}},
        


        


        


        


        


        


        
        OwnerIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

