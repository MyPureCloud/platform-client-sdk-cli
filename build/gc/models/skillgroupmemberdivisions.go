package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupmemberdivisionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupmemberdivisionsDud struct { 
    


    

}

// Skillgroupmemberdivisions
type Skillgroupmemberdivisions struct { 
    // AddDivisionIds
    AddDivisionIds []string `json:"addDivisionIds"`


    // RemoveDivisionIds
    RemoveDivisionIds []string `json:"removeDivisionIds"`

}

// String returns a JSON representation of the model
func (o *Skillgroupmemberdivisions) String() string {
     o.AddDivisionIds = []string{""} 
     o.RemoveDivisionIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupmemberdivisions) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupmemberdivisions

    if SkillgroupmemberdivisionsMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupmemberdivisionsMarshalled = true

    return json.Marshal(&struct {
        
        AddDivisionIds []string `json:"addDivisionIds"`
        
        RemoveDivisionIds []string `json:"removeDivisionIds"`
        *Alias
    }{

        
        AddDivisionIds: []string{""},
        


        
        RemoveDivisionIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

