package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateactivitycoderequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateactivitycoderequestitemDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Bulkupdateactivitycoderequestitem
type Bulkupdateactivitycoderequestitem struct { 
    // Name - The name of the activity code
    Name string `json:"name"`


    // Category - The activity code's category. Attempting to change the category of a default activity code will return an error
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


    // PlanningGroupIds - The planning group IDs associated with this activity code
    PlanningGroupIds Listwrapperstring `json:"planningGroupIds"`


    // Style - The style configuration for the activity code
    Style Valuewrapperactivitycodestyle `json:"style"`


    // Metadata - Version metadata for the associated business unit's list of activity codes
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // Id - The ID of the activity code
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateactivitycoderequestitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateactivitycoderequestitem) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateactivitycoderequestitem

    if BulkupdateactivitycoderequestitemMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateactivitycoderequestitemMarshalled = true

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
        
        PlanningGroupIds Listwrapperstring `json:"planningGroupIds"`
        
        Style Valuewrapperactivitycodestyle `json:"style"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

