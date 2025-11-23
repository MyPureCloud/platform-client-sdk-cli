package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulescriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulescriteriaDud struct { 
    


    


    

}

// Usersrulescriteria
type Usersrulescriteria struct { 
    // Id - The internal ID for this criteria
    Id string `json:"id"`


    // Operator - The operator for this criteria
    Operator string `json:"operator"`


    // Group - The user selection groups for this criteria
    Group []Usersrulesgroupitem `json:"group"`

}

// String returns a JSON representation of the model
func (o *Usersrulescriteria) String() string {
    
    
     o.Group = []Usersrulesgroupitem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulescriteria) MarshalJSON() ([]byte, error) {
    type Alias Usersrulescriteria

    if UsersrulescriteriaMarshalled {
        return []byte("{}"), nil
    }
    UsersrulescriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Operator string `json:"operator"`
        
        Group []Usersrulesgroupitem `json:"group"`
        *Alias
    }{

        


        


        
        Group: []Usersrulesgroupitem{{}},
        

        Alias: (*Alias)(u),
    })
}

