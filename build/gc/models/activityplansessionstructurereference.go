package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivityplansessionstructurereferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivityplansessionstructurereferenceDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Activityplansessionstructurereference
type Activityplansessionstructurereference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Users - The list of users to delete from this session
    Users []Userreference `json:"users"`


    

}

// String returns a JSON representation of the model
func (o *Activityplansessionstructurereference) String() string {
    
     o.Users = []Userreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activityplansessionstructurereference) MarshalJSON() ([]byte, error) {
    type Alias Activityplansessionstructurereference

    if ActivityplansessionstructurereferenceMarshalled {
        return []byte("{}"), nil
    }
    ActivityplansessionstructurereferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Users []Userreference `json:"users"`
        *Alias
    }{

        


        
        Users: []Userreference{{}},
        


        

        Alias: (*Alias)(u),
    })
}

