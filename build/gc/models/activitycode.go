package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodeDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    


    


    


    


    


    


    


    


    

}

// Activitycode - Activity code data
type Activitycode struct { 
    


    


    // Name - The name of the activity code. Default activity codes will be created with an empty name
    Name string `json:"name"`


    // IsActive - Whether this activity code is active or has been deleted
    IsActive bool `json:"isActive"`


    // IsDefault - Whether this is a default activity code
    IsDefault bool `json:"isDefault"`


    // Category - The activity code's category.
    Category string `json:"category"`


    // LengthInMinutes - The default length of the activity in minutes
    LengthInMinutes int `json:"lengthInMinutes"`


    // CountsAsPaidTime - Whether an agent is paid while performing this activity
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // CountsAsWorkTime - Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
    CountsAsWorkTime bool `json:"countsAsWorkTime"`


    // AgentTimeOffSelectable - Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
    AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`


    // Metadata - Version metadata for the associated management unit's list of activity codes
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Activitycode) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycode) MarshalJSON() ([]byte, error) {
    type Alias Activitycode

    if ActivitycodeMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodeMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        Name string `json:"name"`
        
        IsActive bool `json:"isActive"`
        
        IsDefault bool `json:"isDefault"`
        
        Category string `json:"category"`
        
        LengthInMinutes int `json:"lengthInMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        CountsAsWorkTime bool `json:"countsAsWorkTime"`
        
        AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

