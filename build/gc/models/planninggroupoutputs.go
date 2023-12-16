package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupoutputsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupoutputsDud struct { 
    


    


    


    


    

}

// Planninggroupoutputs
type Planninggroupoutputs struct { 
    // PlanningGroupId - The ID for for the associated planning group result
    PlanningGroupId string `json:"planningGroupId"`


    // ServiceLevelPerInterval - List of Service Level percentage (0.0-100.0) results per interval
    ServiceLevelPerInterval []float64 `json:"serviceLevelPerInterval"`


    // OccupancyPerInterval - List of Occupancy percentage (0.0-100.0) results per interval
    OccupancyPerInterval []float64 `json:"occupancyPerInterval"`


    // AverageSpeedOfAnswerSecondsPerInterval - List of Average Speed of Answer (in seconds) results per interval
    AverageSpeedOfAnswerSecondsPerInterval []float64 `json:"averageSpeedOfAnswerSecondsPerInterval"`


    // AbandonRatePerInterval - List of Abandon rate percentage (0.0-100.0) results per interval
    AbandonRatePerInterval []float64 `json:"abandonRatePerInterval"`

}

// String returns a JSON representation of the model
func (o *Planninggroupoutputs) String() string {
    
     o.ServiceLevelPerInterval = []float64{0.0} 
     o.OccupancyPerInterval = []float64{0.0} 
     o.AverageSpeedOfAnswerSecondsPerInterval = []float64{0.0} 
     o.AbandonRatePerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroupoutputs) MarshalJSON() ([]byte, error) {
    type Alias Planninggroupoutputs

    if PlanninggroupoutputsMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupoutputsMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        ServiceLevelPerInterval []float64 `json:"serviceLevelPerInterval"`
        
        OccupancyPerInterval []float64 `json:"occupancyPerInterval"`
        
        AverageSpeedOfAnswerSecondsPerInterval []float64 `json:"averageSpeedOfAnswerSecondsPerInterval"`
        
        AbandonRatePerInterval []float64 `json:"abandonRatePerInterval"`
        *Alias
    }{

        


        
        ServiceLevelPerInterval: []float64{0.0},
        


        
        OccupancyPerInterval: []float64{0.0},
        


        
        AverageSpeedOfAnswerSecondsPerInterval: []float64{0.0},
        


        
        AbandonRatePerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

