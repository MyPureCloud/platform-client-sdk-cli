package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivityqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivityqueryDud struct { 
    


    


    


    

}

// Teamactivityquery
type Teamactivityquery struct { 
    // Metrics - List of requested metrics
    Metrics []Teamactivityquerymetric `json:"metrics"`


    // GroupBy - Dimension(s) to group by
    GroupBy []string `json:"groupBy"`


    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Teamactivityqueryfilter `json:"filter"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Teamactivityquery) String() string {
     o.Metrics = []Teamactivityquerymetric{{}} 
     o.GroupBy = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivityquery) MarshalJSON() ([]byte, error) {
    type Alias Teamactivityquery

    if TeamactivityqueryMarshalled {
        return []byte("{}"), nil
    }
    TeamactivityqueryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Teamactivityquerymetric `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Teamactivityqueryfilter `json:"filter"`
        
        Order string `json:"order"`
        *Alias
    }{

        
        Metrics: []Teamactivityquerymetric{{}},
        


        
        GroupBy: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

