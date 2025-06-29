package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmoduleautoassignresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmoduleautoassignresponseDud struct { 
    


    

}

// Learningmoduleautoassignresponse - Auto assign response
type Learningmoduleautoassignresponse struct { 
    // Rule - The rule reference
    Rule Usersrulesrulereference `json:"rule"`


    // Enabled - Whether the rule is enabled for the module
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Learningmoduleautoassignresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmoduleautoassignresponse) MarshalJSON() ([]byte, error) {
    type Alias Learningmoduleautoassignresponse

    if LearningmoduleautoassignresponseMarshalled {
        return []byte("{}"), nil
    }
    LearningmoduleautoassignresponseMarshalled = true

    return json.Marshal(&struct {
        
        Rule Usersrulesrulereference `json:"rule"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

