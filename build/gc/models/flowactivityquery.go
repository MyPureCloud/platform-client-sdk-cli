package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivityqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivityqueryDud struct { 
    


    


    


    

}

// Flowactivityquery
type Flowactivityquery struct { 
    // Metrics - List of requested metrics
    Metrics []Flowactivityquerymetric `json:"metrics"`


    // GroupBy - Dimension(s) to group by
    GroupBy []string `json:"groupBy"`


    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Flowactivityqueryfilter `json:"filter"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Flowactivityquery) String() string {
     o.Metrics = []Flowactivityquerymetric{{}} 
     o.GroupBy = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivityquery) MarshalJSON() ([]byte, error) {
    type Alias Flowactivityquery

    if FlowactivityqueryMarshalled {
        return []byte("{}"), nil
    }
    FlowactivityqueryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Flowactivityquerymetric `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Flowactivityqueryfilter `json:"filter"`
        
        Order string `json:"order"`
        *Alias
    }{

        
        Metrics: []Flowactivityquerymetric{{}},
        


        
        GroupBy: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

