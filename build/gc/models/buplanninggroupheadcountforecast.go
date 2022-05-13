package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuplanninggroupheadcountforecastMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuplanninggroupheadcountforecastDud struct { 
    


    


    

}

// Buplanninggroupheadcountforecast
type Buplanninggroupheadcountforecast struct { 
    // PlanningGroup - The planning group to which this portion of the headcount forecast applies
    PlanningGroup Planninggroupreference `json:"planningGroup"`


    // RequiredPerInterval - Required headcount per interval, referenced against the reference start date
    RequiredPerInterval []float64 `json:"requiredPerInterval"`


    // RequiredWithoutShrinkagePerInterval - Required headcount per interval without accounting for shrinkage, referenced against the reference start date
    RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`

}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecast) String() string {
    
     o.RequiredPerInterval = []float64{0.0} 
     o.RequiredWithoutShrinkagePerInterval = []float64{0.0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buplanninggroupheadcountforecast) MarshalJSON() ([]byte, error) {
    type Alias Buplanninggroupheadcountforecast

    if BuplanninggroupheadcountforecastMarshalled {
        return []byte("{}"), nil
    }
    BuplanninggroupheadcountforecastMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup Planninggroupreference `json:"planningGroup"`
        
        RequiredPerInterval []float64 `json:"requiredPerInterval"`
        
        RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`
        *Alias
    }{

        


        
        RequiredPerInterval: []float64{0.0},
        


        
        RequiredWithoutShrinkagePerInterval: []float64{0.0},
        

        Alias: (*Alias)(u),
    })
}

