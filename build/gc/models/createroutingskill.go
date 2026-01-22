package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateroutingskillMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateroutingskillDud struct { 
    

}

// Createroutingskill
type Createroutingskill struct { 
    // Name - The name of the skill.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Createroutingskill) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createroutingskill) MarshalJSON() ([]byte, error) {
    type Alias Createroutingskill

    if CreateroutingskillMarshalled {
        return []byte("{}"), nil
    }
    CreateroutingskillMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

