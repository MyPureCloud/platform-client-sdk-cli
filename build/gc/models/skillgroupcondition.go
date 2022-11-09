package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillgroupconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillgroupconditionDud struct { 
    


    


    

}

// Skillgroupcondition
type Skillgroupcondition struct { 
    // RoutingSkillConditions - Routing skill conditions that will be used for building the query
    RoutingSkillConditions []Skillgrouproutingcondition `json:"routingSkillConditions"`


    // LanguageSkillConditions - Routing skill conditions that will be used for building the query
    LanguageSkillConditions []Skillgrouplanguagecondition `json:"languageSkillConditions"`


    // Operation - Operator that will be applied to the conditions
    Operation string `json:"operation"`

}

// String returns a JSON representation of the model
func (o *Skillgroupcondition) String() string {
     o.RoutingSkillConditions = []Skillgrouproutingcondition{{}} 
     o.LanguageSkillConditions = []Skillgrouplanguagecondition{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillgroupcondition) MarshalJSON() ([]byte, error) {
    type Alias Skillgroupcondition

    if SkillgroupconditionMarshalled {
        return []byte("{}"), nil
    }
    SkillgroupconditionMarshalled = true

    return json.Marshal(&struct {
        
        RoutingSkillConditions []Skillgrouproutingcondition `json:"routingSkillConditions"`
        
        LanguageSkillConditions []Skillgrouplanguagecondition `json:"languageSkillConditions"`
        
        Operation string `json:"operation"`
        *Alias
    }{

        
        RoutingSkillConditions: []Skillgrouproutingcondition{{}},
        


        
        LanguageSkillConditions: []Skillgrouplanguagecondition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

