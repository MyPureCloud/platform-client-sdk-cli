package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RoutingactivityqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RoutingactivityqueryDud struct { 
    


    


    


    

}

// Routingactivityquery
type Routingactivityquery struct { 
    // Metrics - List of requested metrics
    Metrics []Routingactivityquerymetric `json:"metrics"`


    // GroupBy - Dimension(s) to group by
    GroupBy []string `json:"groupBy"`


    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Routingactivityqueryfilter `json:"filter"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Routingactivityquery) String() string {
     o.Metrics = []Routingactivityquerymetric{{}} 
     o.GroupBy = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Routingactivityquery) MarshalJSON() ([]byte, error) {
    type Alias Routingactivityquery

    if RoutingactivityqueryMarshalled {
        return []byte("{}"), nil
    }
    RoutingactivityqueryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Routingactivityquerymetric `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Routingactivityqueryfilter `json:"filter"`
        
        Order string `json:"order"`
        *Alias
    }{

        
        Metrics: []Routingactivityquerymetric{{}},
        


        
        GroupBy: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

