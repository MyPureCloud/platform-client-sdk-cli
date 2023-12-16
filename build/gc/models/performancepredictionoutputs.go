package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PerformancepredictionoutputsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PerformancepredictionoutputsDud struct { 
    


    


    

}

// Performancepredictionoutputs
type Performancepredictionoutputs struct { 
    // CalculationStartDate - Date as an ISO-8601 string, corresponding to the beginning of the performance prediction results
    CalculationStartDate time.Time `json:"calculationStartDate"`


    // CalculationIntervalLengthMinutes - Interval length of the response metrics
    CalculationIntervalLengthMinutes int `json:"calculationIntervalLengthMinutes"`


    // PlanningGroupResults - List of planning group level performance prediction results
    PlanningGroupResults []Planninggroupoutputs `json:"planningGroupResults"`

}

// String returns a JSON representation of the model
func (o *Performancepredictionoutputs) String() string {
    
    
     o.PlanningGroupResults = []Planninggroupoutputs{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Performancepredictionoutputs) MarshalJSON() ([]byte, error) {
    type Alias Performancepredictionoutputs

    if PerformancepredictionoutputsMarshalled {
        return []byte("{}"), nil
    }
    PerformancepredictionoutputsMarshalled = true

    return json.Marshal(&struct {
        
        CalculationStartDate time.Time `json:"calculationStartDate"`
        
        CalculationIntervalLengthMinutes int `json:"calculationIntervalLengthMinutes"`
        
        PlanningGroupResults []Planninggroupoutputs `json:"planningGroupResults"`
        *Alias
    }{

        


        


        
        PlanningGroupResults: []Planninggroupoutputs{{}},
        

        Alias: (*Alias)(u),
    })
}

