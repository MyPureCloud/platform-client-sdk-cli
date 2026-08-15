package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContinuousforecasttimeseriesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContinuousforecasttimeseriesDud struct { 
    


    

}

// Continuousforecasttimeseries
type Continuousforecasttimeseries struct { 
    // PlanningGroup - The planning group ID
    PlanningGroup string `json:"planningGroup"`


    // Weeks - List of data for each week
    Weeks []Continuousforecastweeks `json:"weeks"`

}

// String returns a JSON representation of the model
func (o *Continuousforecasttimeseries) String() string {
    
     o.Weeks = []Continuousforecastweeks{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Continuousforecasttimeseries) MarshalJSON() ([]byte, error) {
    type Alias Continuousforecasttimeseries

    if ContinuousforecasttimeseriesMarshalled {
        return []byte("{}"), nil
    }
    ContinuousforecasttimeseriesMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup string `json:"planningGroup"`
        
        Weeks []Continuousforecastweeks `json:"weeks"`
        *Alias
    }{

        


        
        Weeks: []Continuousforecastweeks{{}},
        

        Alias: (*Alias)(u),
    })
}

