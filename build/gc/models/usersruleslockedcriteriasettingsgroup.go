package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersruleslockedcriteriasettingsgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersruleslockedcriteriasettingsgroupDud struct { 
    


    


    

}

// Usersruleslockedcriteriasettingsgroup
type Usersruleslockedcriteriasettingsgroup struct { 
    // Operators - The allowed operators for this criteria
    Operators []string `json:"operators"`


    // Container - The container that the ids belong to
    Container string `json:"container"`


    // MaxIdCount - Maximum number of ids that can be specified in this container
    MaxIdCount int `json:"maxIdCount"`

}

// String returns a JSON representation of the model
func (o *Usersruleslockedcriteriasettingsgroup) String() string {
     o.Operators = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersruleslockedcriteriasettingsgroup) MarshalJSON() ([]byte, error) {
    type Alias Usersruleslockedcriteriasettingsgroup

    if UsersruleslockedcriteriasettingsgroupMarshalled {
        return []byte("{}"), nil
    }
    UsersruleslockedcriteriasettingsgroupMarshalled = true

    return json.Marshal(&struct {
        
        Operators []string `json:"operators"`
        
        Container string `json:"container"`
        
        MaxIdCount int `json:"maxIdCount"`
        *Alias
    }{

        
        Operators: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

