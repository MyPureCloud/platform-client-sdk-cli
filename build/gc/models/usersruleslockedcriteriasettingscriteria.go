package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersruleslockedcriteriasettingscriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersruleslockedcriteriasettingscriteriaDud struct { 
    


    

}

// Usersruleslockedcriteriasettingscriteria
type Usersruleslockedcriteriasettingscriteria struct { 
    // Operator - The operator for this criteria
    Operator string `json:"operator"`


    // Group - The user selection groups for this criteria
    Group []Usersruleslockedcriteriasettingsgroup `json:"group"`

}

// String returns a JSON representation of the model
func (o *Usersruleslockedcriteriasettingscriteria) String() string {
    
     o.Group = []Usersruleslockedcriteriasettingsgroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersruleslockedcriteriasettingscriteria) MarshalJSON() ([]byte, error) {
    type Alias Usersruleslockedcriteriasettingscriteria

    if UsersruleslockedcriteriasettingscriteriaMarshalled {
        return []byte("{}"), nil
    }
    UsersruleslockedcriteriasettingscriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Operator string `json:"operator"`
        
        Group []Usersruleslockedcriteriasettingsgroup `json:"group"`
        *Alias
    }{

        


        
        Group: []Usersruleslockedcriteriasettingsgroup{{}},
        

        Alias: (*Alias)(u),
    })
}

