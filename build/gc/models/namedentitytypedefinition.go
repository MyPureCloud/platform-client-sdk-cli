package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitytypedefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitytypedefinitionDud struct { 
    


    


    

}

// Namedentitytypedefinition
type Namedentitytypedefinition struct { 
    // Name - The name of the entity type.
    Name string `json:"name"`


    // Description - Description of the of the named entity type.
    Description string `json:"description"`


    // Mechanism - The mechanism enabling detection of the named entity type.
    Mechanism Namedentitytypemechanism `json:"mechanism"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypedefinition) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitytypedefinition) MarshalJSON() ([]byte, error) {
    type Alias Namedentitytypedefinition

    if NamedentitytypedefinitionMarshalled {
        return []byte("{}"), nil
    }
    NamedentitytypedefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Mechanism Namedentitytypemechanism `json:"mechanism"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

