package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinimumstaffingrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinimumstaffingrequestDud struct { 
    


    


    


    

}

// Minimumstaffingrequest
type Minimumstaffingrequest struct { 
    // Enabled - Whether the setting is turned on or off
    Enabled bool `json:"enabled"`


    // MinimumValue - Default minimum staff value to be applied to all planning groups
    MinimumValue float64 `json:"minimumValue"`


    // PlanningGroupOverrides - List of planning groups with their minimum staff value overrides and the days to which the overrides apply. Setting the enclosed list to empty will clear out all existing overrides
    PlanningGroupOverrides Listwrapperplanninggroupminimumsrequest `json:"planningGroupOverrides"`


    // ApplicableIntervals - The intervals to which the minimum staff values will apply
    ApplicableIntervals string `json:"applicableIntervals"`

}

// String returns a JSON representation of the model
func (o *Minimumstaffingrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minimumstaffingrequest) MarshalJSON() ([]byte, error) {
    type Alias Minimumstaffingrequest

    if MinimumstaffingrequestMarshalled {
        return []byte("{}"), nil
    }
    MinimumstaffingrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        MinimumValue float64 `json:"minimumValue"`
        
        PlanningGroupOverrides Listwrapperplanninggroupminimumsrequest `json:"planningGroupOverrides"`
        
        ApplicableIntervals string `json:"applicableIntervals"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

