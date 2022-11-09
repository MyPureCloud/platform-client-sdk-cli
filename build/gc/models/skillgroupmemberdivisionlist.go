package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupmemberdivisionlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupmemberdivisionlistDud struct { 
    

}

// Skillgroupmemberdivisionlist
type Skillgroupmemberdivisionlist struct { 
    // Entities
    Entities []Division `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Skillgroupmemberdivisionlist) String() string {
     o.Entities = []Division{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupmemberdivisionlist) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupmemberdivisionlist

    if SkillgroupmemberdivisionlistMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupmemberdivisionlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Division `json:"entities"`
        *Alias
    }{

        
        Entities: []Division{{}},
        

        Alias: (*Alias)(u),
    })
}

