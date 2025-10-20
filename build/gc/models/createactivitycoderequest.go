package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateactivitycoderequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateactivitycoderequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Createactivitycoderequest
type Createactivitycoderequest struct { 
    // Name - The name of the activity code
    Name string `json:"name"`


    // Category - The activity code's category
    Category string `json:"category"`


    // LengthInMinutes - The default length of the activity in minutes
    LengthInMinutes int `json:"lengthInMinutes"`


    // CountsAsPaidTime - Whether an agent is paid while performing this activity
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // CountsAsWorkTime - Indicates whether or not the activity should be counted as work time
    CountsAsWorkTime bool `json:"countsAsWorkTime"`


    // AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request
    AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`


    // CountsTowardShrinkage - Whether or not this activity code counts toward shrinkage calculations
    CountsTowardShrinkage bool `json:"countsTowardShrinkage"`


    // PlannedShrinkage - Whether this activity code is considered planned or unplanned shrinkage
    PlannedShrinkage bool `json:"plannedShrinkage"`


    // Interruptible - Whether this activity code is considered interruptible
    Interruptible bool `json:"interruptible"`


    // SecondaryPresences - The secondary presences of this activity code
    SecondaryPresences []Secondarypresence `json:"secondaryPresences"`


    // PlanningGroupIds - The planning group IDs associated with this activity code
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Createactivitycoderequest) String() string {
    
    
    
    
    
    
    
    
    
     o.SecondaryPresences = []Secondarypresence{{}} 
     o.PlanningGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createactivitycoderequest) MarshalJSON() ([]byte, error) {
    type Alias Createactivitycoderequest

    if CreateactivitycoderequestMarshalled {
        return []byte("{}"), nil
    }
    CreateactivitycoderequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Category string `json:"category"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        CountsAsWorkTime bool `json:"countsAsWorkTime"`
        
        AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`
        
        CountsTowardShrinkage bool `json:"countsTowardShrinkage"`
        
        PlannedShrinkage bool `json:"plannedShrinkage"`
        
        Interruptible bool `json:"interruptible"`
        
        SecondaryPresences []Secondarypresence `json:"secondaryPresences"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        SecondaryPresences: []Secondarypresence{{}},
        


        
        PlanningGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

