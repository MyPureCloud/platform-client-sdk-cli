package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionsDud struct { 
    

}

// Actions
type Actions struct { 
    // SkillsToRemove
    SkillsToRemove []Skillstoremove `json:"skillsToRemove"`

}

// String returns a JSON representation of the model
func (o *Actions) String() string {
     o.SkillsToRemove = []Skillstoremove{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actions) MarshalJSON() ([]byte, error) {
    type Alias Actions

    if ActionsMarshalled {
        return []byte("{}"), nil
    }
    ActionsMarshalled = true

    return json.Marshal(&struct {
        
        SkillsToRemove []Skillstoremove `json:"skillsToRemove"`
        *Alias
    }{

        
        SkillsToRemove: []Skillstoremove{{}},
        

        Alias: (*Alias)(u),
    })
}

