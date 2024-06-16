package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FieldconfigsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FieldconfigsDud struct { 
    


    


    

}

// Fieldconfigs
type Fieldconfigs struct { 
    // Org
    Org Fieldconfig `json:"org"`


    // Person
    Person Fieldconfig `json:"person"`


    // Group
    Group Fieldconfig `json:"group"`

}

// String returns a JSON representation of the model
func (o *Fieldconfigs) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Fieldconfigs) MarshalJSON() ([]byte, error) {
    type Alias Fieldconfigs

    if FieldconfigsMarshalled {
        return []byte("{}"), nil
    }
    FieldconfigsMarshalled = true

    return json.Marshal(&struct {
        
        Org Fieldconfig `json:"org"`
        
        Person Fieldconfig `json:"person"`
        
        Group Fieldconfig `json:"group"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

