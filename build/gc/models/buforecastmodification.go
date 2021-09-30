package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecastmodificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecastmodificationDud struct { 
    


    


    


    


    LegacyMetric string `json:"legacyMetric"`


    


    


    


    


    


    

}

// Buforecastmodification
type Buforecastmodification struct { 
    // VarType - The type of the modification
    VarType string `json:"type"`


    // StartIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the first interval to which to apply this modification. Must be null if values is populated
    StartIntervalIndex int `json:"startIntervalIndex"`


    // EndIntervalIndex - The number of 15 minute intervals past referenceStartDate representing the last interval to which to apply this modification.  Must be null if values is populated
    EndIntervalIndex int `json:"endIntervalIndex"`


    // Metric - The metric to which this modification applies
    Metric string `json:"metric"`


    


    // Value - The value of the modification.  Must be null if \"values\" is populated
    Value float64 `json:"value"`


    // Values - The list of values to update.  Only applicable for grid-type modifications. Must be null if \"value\" is populated
    Values []Wfmforecastmodificationintervaloffsetvalue `json:"values"`


    // DisplayGranularity - The client side display granularity of the modification, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    DisplayGranularity string `json:"displayGranularity"`


    // Granularity - The actual granularity of the modification as stored behind the scenes, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Granularity string `json:"granularity"`


    // Enabled - Whether the modification is enabled for the forecast
    Enabled bool `json:"enabled"`


    // PlanningGroupIds - The IDs of the planning groups to which this forecast modification applies.  Leave empty to apply to all
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Buforecastmodification) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Values = []Wfmforecastmodificationintervaloffsetvalue{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.PlanningGroupIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecastmodification) MarshalJSON() ([]byte, error) {
    type Alias Buforecastmodification

    if BuforecastmodificationMarshalled {
        return []byte("{}"), nil
    }
    BuforecastmodificationMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        StartIntervalIndex int `json:"startIntervalIndex"`
        
        EndIntervalIndex int `json:"endIntervalIndex"`
        
        Metric string `json:"metric"`
        
        
        
        Value float64 `json:"value"`
        
        Values []Wfmforecastmodificationintervaloffsetvalue `json:"values"`
        
        DisplayGranularity string `json:"displayGranularity"`
        
        Granularity string `json:"granularity"`
        
        Enabled bool `json:"enabled"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Values: []Wfmforecastmodificationintervaloffsetvalue{{}},
        

        

        

        

        

        

        

        

        
        PlanningGroupIds: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

