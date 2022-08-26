package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredgroupsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredgroupsDud struct { 
    

}

// Unansweredgroups
type Unansweredgroups struct { 
    // Entities
    Entities []Unansweredgroup `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Unansweredgroups) String() string {
     o.Entities = []Unansweredgroup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredgroups) MarshalJSON() ([]byte, error) {
    type Alias Unansweredgroups

    if UnansweredgroupsMarshalled {
        return []byte("{}"), nil
    }
    UnansweredgroupsMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Unansweredgroup `json:"entities"`
        *Alias
    }{

        
        Entities: []Unansweredgroup{{}},
        

        Alias: (*Alias)(u),
    })
}

