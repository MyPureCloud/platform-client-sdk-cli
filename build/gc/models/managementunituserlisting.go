package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunituserlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunituserlistingDud struct { 
    


    

}

// Managementunituserlisting
type Managementunituserlisting struct { 
    // ManagementUnit - The management unit associated with the users
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // Users - Users in the management unit
    Users []Userreference `json:"users"`

}

// String returns a JSON representation of the model
func (o *Managementunituserlisting) String() string {
    
     o.Users = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunituserlisting) MarshalJSON() ([]byte, error) {
    type Alias Managementunituserlisting

    if ManagementunituserlistingMarshalled {
        return []byte("{}"), nil
    }
    ManagementunituserlistingMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        Users []Userreference `json:"users"`
        *Alias
    }{

        


        
        Users: []Userreference{{}},
        

        Alias: (*Alias)(u),
    })
}

