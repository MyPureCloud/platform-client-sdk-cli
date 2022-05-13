package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupupdateDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Groupupdate
type Groupupdate struct { 
    


    // Name - The group name.
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // State - State of the group.
    State string `json:"state"`


    // Version - Current version for this resource.
    Version int `json:"version"`


    // Images
    Images []Userimage `json:"images"`


    // Addresses
    Addresses []Groupcontact `json:"addresses"`


    // RulesVisible - Are membership rules visible to the person requesting to view the group
    RulesVisible bool `json:"rulesVisible"`


    // Visibility - Who can view this group
    Visibility string `json:"visibility"`


    // OwnerIds - Owners of the group
    OwnerIds []string `json:"ownerIds"`


    

}

// String returns a JSON representation of the model
func (o *Groupupdate) String() string {
    
    
    
    
     o.Images = []Userimage{{}} 
     o.Addresses = []Groupcontact{{}} 
    
    
     o.OwnerIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupupdate) MarshalJSON() ([]byte, error) {
    type Alias Groupupdate

    if GroupupdateMarshalled {
        return []byte("{}"), nil
    }
    GroupupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        State string `json:"state"`
        
        Version int `json:"version"`
        
        Images []Userimage `json:"images"`
        
        Addresses []Groupcontact `json:"addresses"`
        
        RulesVisible bool `json:"rulesVisible"`
        
        Visibility string `json:"visibility"`
        
        OwnerIds []string `json:"ownerIds"`
        *Alias
    }{

        


        


        


        


        


        
        Images: []Userimage{{}},
        


        
        Addresses: []Groupcontact{{}},
        


        


        


        
        OwnerIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

