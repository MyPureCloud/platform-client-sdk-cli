package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeseriesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeseriesDud struct { 
    


    

}

// Timeseries
type Timeseries struct { 
    // PlanningGroup - The planning group ID
    PlanningGroup string `json:"planningGroup"`


    // Weeks - List of data for each week
    Weeks []Weeks `json:"weeks"`

}

// String returns a JSON representation of the model
func (o *Timeseries) String() string {
    
     o.Weeks = []Weeks{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeseries) MarshalJSON() ([]byte, error) {
    type Alias Timeseries

    if TimeseriesMarshalled {
        return []byte("{}"), nil
    }
    TimeseriesMarshalled = true

    return json.Marshal(&struct {
        
        PlanningGroup string `json:"planningGroup"`
        
        Weeks []Weeks `json:"weeks"`
        *Alias
    }{

        


        
        Weeks: []Weeks{{}},
        

        Alias: (*Alias)(u),
    })
}

