package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesupdaterulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesupdaterulerequestDud struct { 
    


    


    

}

// Usersrulesupdaterulerequest - Update users rule request
type Usersrulesupdaterulerequest struct { 
    // Name - The name of the rule
    Name string `json:"name"`


    // Description - The description of the rule
    Description string `json:"description"`


    // Criteria - The criteria of the rule
    Criteria []Usersrulescriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Usersrulesupdaterulerequest) String() string {
    
    
     o.Criteria = []Usersrulescriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesupdaterulerequest) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesupdaterulerequest

    if UsersrulesupdaterulerequestMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesupdaterulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Criteria []Usersrulescriteria `json:"criteria"`
        *Alias
    }{

        


        


        
        Criteria: []Usersrulescriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

