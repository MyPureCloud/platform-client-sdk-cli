package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicgroupskillconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicgroupskillconditionDud struct { 
    


    


    

}

// Dynamicgroupskillcondition
type Dynamicgroupskillcondition struct { 
    // RoutingSkillConditions - Routing skill conditions that will be used for building the query
    RoutingSkillConditions []Dynamicgrouproutingskillcondition `json:"routingSkillConditions"`


    // LanguageSkillConditions - Routing skill conditions that will be used for building the query
    LanguageSkillConditions []Dynamicgrouplanguageskillcondition `json:"languageSkillConditions"`


    // Operation - Operator that will be applied to the conditions
    Operation string `json:"operation"`

}

// String returns a JSON representation of the model
func (o *Dynamicgroupskillcondition) String() string {
     o.RoutingSkillConditions = []Dynamicgrouproutingskillcondition{{}} 
     o.LanguageSkillConditions = []Dynamicgrouplanguageskillcondition{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicgroupskillcondition) MarshalJSON() ([]byte, error) {
    type Alias Dynamicgroupskillcondition

    if DynamicgroupskillconditionMarshalled {
        return []byte("{}"), nil
    }
    DynamicgroupskillconditionMarshalled = true

    return json.Marshal(&struct {
        
        RoutingSkillConditions []Dynamicgrouproutingskillcondition `json:"routingSkillConditions"`
        
        LanguageSkillConditions []Dynamicgrouplanguageskillcondition `json:"languageSkillConditions"`
        
        Operation string `json:"operation"`
        *Alias
    }{

        
        RoutingSkillConditions: []Dynamicgrouproutingskillcondition{{}},
        


        
        LanguageSkillConditions: []Dynamicgrouplanguageskillcondition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

