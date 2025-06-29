package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningmoduleautoassignrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningmoduleautoassignrequestDud struct { 
    


    

}

// Learningmoduleautoassignrequest - Auto assign request
type Learningmoduleautoassignrequest struct { 
    // RuleId - The id of the rule
    RuleId string `json:"ruleId"`


    // Enabled - Whether the rule is enabled for the module
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Learningmoduleautoassignrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningmoduleautoassignrequest) MarshalJSON() ([]byte, error) {
    type Alias Learningmoduleautoassignrequest

    if LearningmoduleautoassignrequestMarshalled {
        return []byte("{}"), nil
    }
    LearningmoduleautoassignrequestMarshalled = true

    return json.Marshal(&struct {
        
        RuleId string `json:"ruleId"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

