package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsersrulesgroupitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsersrulesgroupitemDud struct { 
    


    


    


    

}

// Usersrulesgroupitem
type Usersrulesgroupitem struct { 
    // Id - The internal ID for this group
    Id string `json:"id"`


    // Operator - The operator for this group
    Operator string `json:"operator"`


    // Container - The container that the ids belong to
    Container string `json:"container"`


    // Values - The ids and contextIds to select for this group
    Values []Usersrulesvalue `json:"values"`

}

// String returns a JSON representation of the model
func (o *Usersrulesgroupitem) String() string {
    
    
    
     o.Values = []Usersrulesvalue{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usersrulesgroupitem) MarshalJSON() ([]byte, error) {
    type Alias Usersrulesgroupitem

    if UsersrulesgroupitemMarshalled {
        return []byte("{}"), nil
    }
    UsersrulesgroupitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Operator string `json:"operator"`
        
        Container string `json:"container"`
        
        Values []Usersrulesvalue `json:"values"`
        *Alias
    }{

        


        


        


        
        Values: []Usersrulesvalue{{}},
        

        Alias: (*Alias)(u),
    })
}

