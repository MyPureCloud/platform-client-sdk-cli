package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesqueryrulerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesqueryrulerequestDud struct { 
    


    


    

}

// Usersrulesqueryrulerequest - Users query rule request
type Usersrulesqueryrulerequest struct { 
    // VarType - The type of the rule
    VarType string `json:"type"`


    // Criteria - The criteria of the rule
    Criteria []Usersrulescriteria `json:"criteria"`


    // UserIds - The user ids to query
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Usersrulesqueryrulerequest) String() string {
    
     o.Criteria = []Usersrulescriteria{{}} 
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesqueryrulerequest) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesqueryrulerequest

    if UsersrulesqueryrulerequestMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesqueryrulerequestMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Criteria []Usersrulescriteria `json:"criteria"`
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        


        
        Criteria: []Usersrulescriteria{{}},
        


        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

