package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuplanninggroupheadcountforecastuploadschemaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuplanninggroupheadcountforecastuploadschemaDud struct { 
    


    


    

}

// Buplanninggroupheadcountforecastuploadschema
type Buplanninggroupheadcountforecastuploadschema struct { 
    // RequiredPerInterval - Required headcount per interval, referenced against the reference start date
    RequiredPerInterval []float64 `json:"requiredPerInterval"`


    // RequiredWithoutShrinkagePerInterval - Required headcount per interval without accounting for shrinkage, referenced against the reference start date
    RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`


    // PlanningGroupId - The ID of the planning group to which this portion of the headcount forecast applies
    PlanningGroupId string `json:"planningGroupId"`

}

// String returns a JSON representation of the model
func (o *Buplanninggroupheadcountforecastuploadschema) String() string {
     o.RequiredPerInterval = []float64{0.0} 
     o.RequiredWithoutShrinkagePerInterval = []float64{0.0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buplanninggroupheadcountforecastuploadschema) MarshalJSON() ([]byte, error) {
    type Alias Buplanninggroupheadcountforecastuploadschema

    if BuplanninggroupheadcountforecastuploadschemaMarshalled {
        return []byte("{}"), nil
    }
    BuplanninggroupheadcountforecastuploadschemaMarshalled = true

    return json.Marshal(&struct {
        
        RequiredPerInterval []float64 `json:"requiredPerInterval"`
        
        RequiredWithoutShrinkagePerInterval []float64 `json:"requiredWithoutShrinkagePerInterval"`
        
        PlanningGroupId string `json:"planningGroupId"`
        *Alias
    }{

        
        RequiredPerInterval: []float64{0.0},
        


        
        RequiredWithoutShrinkagePerInterval: []float64{0.0},
        


        

        Alias: (*Alias)(u),
    })
}

