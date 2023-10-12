package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicgroupqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicgroupqueryDud struct { 
    

}

// Dynamicgroupquery
type Dynamicgroupquery struct { 
    // SkillConditions - The skill conditions that need to be used for this dynamic group
    SkillConditions []Dynamicgroupskillcondition `json:"skillConditions"`

}

// String returns a JSON representation of the model
func (o *Dynamicgroupquery) String() string {
     o.SkillConditions = []Dynamicgroupskillcondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicgroupquery) MarshalJSON() ([]byte, error) {
    type Alias Dynamicgroupquery

    if DynamicgroupqueryMarshalled {
        return []byte("{}"), nil
    }
    DynamicgroupqueryMarshalled = true

    return json.Marshal(&struct {
        
        SkillConditions []Dynamicgroupskillcondition `json:"skillConditions"`
        *Alias
    }{

        
        SkillConditions: []Dynamicgroupskillcondition{{}},
        

        Alias: (*Alias)(u),
    })
}

