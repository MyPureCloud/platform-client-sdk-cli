package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MinimumstaffingresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MinimumstaffingresponseDud struct { 
    


    


    


    


    

}

// Minimumstaffingresponse
type Minimumstaffingresponse struct { 
    // Enabled - Whether the setting is turned on or off
    Enabled bool `json:"enabled"`


    // MinimumValue - Default minimum staff value to be applied to all planning groups
    MinimumValue float64 `json:"minimumValue"`


    // PlanningGroupOverrides - List of planning groups with their minimum staff value overrides and the days to which the overrides apply
    PlanningGroupOverrides []Planninggroupminimumsresponse `json:"planningGroupOverrides"`


    // ApplicableIntervals - The intervals to which the minimum staff values will apply
    ApplicableIntervals string `json:"applicableIntervals"`


    // Metadata - Metadata for the business unit's minimum staffing settings
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Minimumstaffingresponse) String() string {
    
    
     o.PlanningGroupOverrides = []Planninggroupminimumsresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Minimumstaffingresponse) MarshalJSON() ([]byte, error) {
    type Alias Minimumstaffingresponse

    if MinimumstaffingresponseMarshalled {
        return []byte("{}"), nil
    }
    MinimumstaffingresponseMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        MinimumValue float64 `json:"minimumValue"`
        
        PlanningGroupOverrides []Planninggroupminimumsresponse `json:"planningGroupOverrides"`
        
        ApplicableIntervals string `json:"applicableIntervals"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        
        PlanningGroupOverrides: []Planninggroupminimumsresponse{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

