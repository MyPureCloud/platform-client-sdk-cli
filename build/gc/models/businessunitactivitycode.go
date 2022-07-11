package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BusinessunitactivitycodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BusinessunitactivitycodeDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Businessunitactivitycode
type Businessunitactivitycode struct { 
    


    // Name
    Name string `json:"name"`


    // Active - Whether this activity code is active or has been deleted
    Active bool `json:"active"`


    // DefaultCode - Whether this is a default activity code
    DefaultCode bool `json:"defaultCode"`


    // Category - The category of the activity code
    Category string `json:"category"`


    // LengthInMinutes - The default length of the activity in minutes
    LengthInMinutes int `json:"lengthInMinutes"`


    // CountsAsPaidTime - Whether an agent is paid while performing this activity
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // CountsAsWorkTime - Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
    CountsAsWorkTime bool `json:"countsAsWorkTime"`


    // AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
    AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`


    // CountsTowardShrinkage - Whether or not this activity code counts toward shrinkage calculations
    CountsTowardShrinkage bool `json:"countsTowardShrinkage"`


    // PlannedShrinkage - Whether this activity code is considered planned or unplanned shrinkage
    PlannedShrinkage bool `json:"plannedShrinkage"`


    // Interruptible - Whether this activity code is considered interruptible
    Interruptible bool `json:"interruptible"`


    // SecondaryPresences - The secondary presences of this activity code
    SecondaryPresences []Secondarypresence `json:"secondaryPresences"`


    // Metadata - Version metadata of this activity code
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Businessunitactivitycode) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.SecondaryPresences = []Secondarypresence{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Businessunitactivitycode) MarshalJSON() ([]byte, error) {
    type Alias Businessunitactivitycode

    if BusinessunitactivitycodeMarshalled {
        return []byte("{}"), nil
    }
    BusinessunitactivitycodeMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Active bool `json:"active"`
        
        DefaultCode bool `json:"defaultCode"`
        
        Category string `json:"category"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        CountsAsWorkTime bool `json:"countsAsWorkTime"`
        
        AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`
        
        CountsTowardShrinkage bool `json:"countsTowardShrinkage"`
        
        PlannedShrinkage bool `json:"plannedShrinkage"`
        
        Interruptible bool `json:"interruptible"`
        
        SecondaryPresences []Secondarypresence `json:"secondaryPresences"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        SecondaryPresences: []Secondarypresence{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

