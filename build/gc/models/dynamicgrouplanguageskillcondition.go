package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DynamicgrouplanguageskillconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DynamicgrouplanguageskillconditionDud struct { 
    


    


    


    

}

// Dynamicgrouplanguageskillcondition
type Dynamicgrouplanguageskillcondition struct { 
    // LanguageSkill - The language skill to be used in the skill condition query
    LanguageSkill string `json:"languageSkill"`


    // Comparator - Comparator that will be applied to the proficiency
    Comparator string `json:"comparator"`


    // Proficiency - The skill proficiency that will be used for the language skill. Integer range 0-5
    Proficiency int `json:"proficiency"`


    // ChildConditions - Nested conditions to be applied to this skill condition
    ChildConditions []Dynamicgroupskillcondition `json:"childConditions"`

}

// String returns a JSON representation of the model
func (o *Dynamicgrouplanguageskillcondition) String() string {
    
    
    
     o.ChildConditions = []Dynamicgroupskillcondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dynamicgrouplanguageskillcondition) MarshalJSON() ([]byte, error) {
    type Alias Dynamicgrouplanguageskillcondition

    if DynamicgrouplanguageskillconditionMarshalled {
        return []byte("{}"), nil
    }
    DynamicgrouplanguageskillconditionMarshalled = true

    return json.Marshal(&struct {
        
        LanguageSkill string `json:"languageSkill"`
        
        Comparator string `json:"comparator"`
        
        Proficiency int `json:"proficiency"`
        
        ChildConditions []Dynamicgroupskillcondition `json:"childConditions"`
        *Alias
    }{

        


        


        


        
        ChildConditions: []Dynamicgroupskillcondition{{}},
        

        Alias: (*Alias)(u),
    })
}

