package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuplanninggroupheadcountforecastresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuplanninggroupheadcountforecastresultDud struct { 
    


    


    

}

// Buplanninggroupheadcountforecastresult
type Buplanninggroupheadcountforecastresult struct { 
    // RequiredPerInterval - Required headcount per interval, referenced against the reference start date
    RequiredPerInterval []float64 `json:"requiredPerInterval"`


    // RequiredWithoutShrinkagePerInterval - Required headcount per interval without accounting for shrinkage, referenced against the reference start date
    RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`


    // PlanningGroup - The planning group to which this portion of the headcount forecast applies
    PlanningGroup Planninggroupreference `json:"planningGroup"`

}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecastresult) String() string {
     o.RequiredPerInterval = []float64{0.0} 
     o.RequiredWithoutShrinkagePerInterval = []float64{0.0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buplanninggroupheadcountforecastresult) MarshalJSON() ([]byte, error) {
    type Alias Buplanninggroupheadcountforecastresult

    if BuplanninggroupheadcountforecastresultMarshalled {
        return []byte("{}"), nil
    }
    BuplanninggroupheadcountforecastresultMarshalled = true

    return json.Marshal(&struct {
        
        RequiredPerInterval []float64 `json:"requiredPerInterval"`
        
        RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`
        
        PlanningGroup Planninggroupreference `json:"planningGroup"`
        *Alias
    }{

        
        RequiredPerInterval: []float64{0.0},
        


        
        RequiredWithoutShrinkagePerInterval: []float64{0.0},
        


        

        Alias: (*Alias)(u),
    })
}

