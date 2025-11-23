package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulescreaterulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulescreaterulerequestDud struct { 
    


    


    


    

}

// Usersrulescreaterulerequest - Create users rule request
type Usersrulescreaterulerequest struct { 
    // Name - The name of the rule
    Name string `json:"name"`


    // Description - The description of the rule
    Description string `json:"description"`


    // VarType - The type of the rule
    VarType string `json:"type"`


    // Criteria - The criteria of the rule
    Criteria []Usersrulescriteria `json:"criteria"`

}

// String returns a JSON representation of the model
func (o *Usersrulescreaterulerequest) String() string {
    
    
    
     o.Criteria = []Usersrulescriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulescreaterulerequest) MarshalJSON() ([]byte, error) {
    type Alias Usersrulescreaterulerequest

    if UsersrulescreaterulerequestMarshalled {
        return []byte("{}"), nil
    }
    UsersrulescreaterulerequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        VarType string `json:"type"`
        
        Criteria []Usersrulescriteria `json:"criteria"`
        *Alias
    }{

        


        


        


        
        Criteria: []Usersrulescriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

