package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuforecastgenerationplanninggroupresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuforecastgenerationplanninggroupresultDud struct { 
    


    

}

// Buforecastgenerationplanninggroupresult
type Buforecastgenerationplanninggroupresult struct { 
    // PlanningGroupId - The ID of the planning group
    PlanningGroupId string `json:"planningGroupId"`


    // MetricResults - The generation results for the associated planning group
    MetricResults []Buforecasttimeseriesresult `json:"metricResults"`

}

// String returns a JSON representation of the model
func (o *Buforecastgenerationplanninggroupresult) String() string {
    
     o.MetricResults = []Buforecasttimeseriesresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buforecastgenerationplanninggroupresult) MarshalJSON() ([]byte, error) {
    type Alias Buforecastgenerationplanninggroupresult

    if BuforecastgenerationplanninggroupresultMarshalled {
        return []byte("{}"), nil
    }
    BuforecastgenerationplanninggroupresultMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroupId string `json:"planningGroupId"`
        
        MetricResults []Buforecasttimeseriesresult `json:"metricResults"`
        *Alias
    }{

        


        
        MetricResults: []Buforecasttimeseriesresult{{}},
        

        Alias: (*Alias)(u),
    })
}

