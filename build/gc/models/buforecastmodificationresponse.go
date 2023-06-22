package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecastmodificationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecastmodificationresponseDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Buforecastmodificationresponse
type Buforecastmodificationresponse struct { 
    // VarType - The type of the modification
    VarType string `json:"type"`


    // StartIntervalIndex - The number of intervals past referenceStartDate representing the first interval to which this modification applies
    StartIntervalIndex int `json:"startIntervalIndex"`


    // EndIntervalIndex - The number of intervals past referenceStartDate representing the last interval to which this modification applies
    EndIntervalIndex int `json:"endIntervalIndex"`


    // Metric - The metric to which this modification applies
    Metric string `json:"metric"`


    // LegacyMetric - The legacy metric to which this modification applies if applicable
    LegacyMetric string `json:"legacyMetric"`


    // Value - The value of the modification
    Value float64 `json:"value"`


    // Values - The list of modification values. Only applicable for grid-type modifications
    Values []Wfmforecastmodificationintervaloffsetvalue `json:"values"`


    // DisplayGranularity - The client side display granularity of the modification, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    DisplayGranularity string `json:"displayGranularity"`


    // Granularity - The actual granularity of the modification as stored behind the scenes, expressed in the ISO-8601 duration format. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Granularity string `json:"granularity"`


    // Enabled - Whether the modification is enabled for the forecast
    Enabled bool `json:"enabled"`


    // PlanningGroupIds - The IDs of the planning groups to which this forecast modification applies
    PlanningGroupIds []string `json:"planningGroupIds"`

}

// String returns a JSON representation of the model
func (o *Buforecastmodificationresponse) String() string {
    
    
    
    
    
    
     o.Values = []Wfmforecastmodificationintervaloffsetvalue{{}} 
    
    
    
     o.PlanningGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecastmodificationresponse) MarshalJSON() ([]byte, error) {
    type Alias Buforecastmodificationresponse

    if BuforecastmodificationresponseMarshalled {
        return []byte("{}"), nil
    }
    BuforecastmodificationresponseMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        StartIntervalIndex int `json:"startIntervalIndex"`
        
        EndIntervalIndex int `json:"endIntervalIndex"`
        
        Metric string `json:"metric"`
        
        LegacyMetric string `json:"legacyMetric"`
        
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

