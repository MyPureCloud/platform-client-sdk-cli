package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityqueryDud struct { 
    


    


    


    

}

// Useractivityquery
type Useractivityquery struct { 
    // Metrics - List of requested metrics
    Metrics []Useractivityquerymetric `json:"metrics"`


    // GroupBy - Dimension(s) to group by
    GroupBy []string `json:"groupBy"`


    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Useractivityqueryfilter `json:"filter"`


    // Order - Sort the result set in ascending/descending order. Default is ascending
    Order string `json:"order"`

}

// String returns a JSON representation of the model
func (o *Useractivityquery) String() string {
     o.Metrics = []Useractivityquerymetric{{}} 
     o.GroupBy = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityquery) MarshalJSON() ([]byte, error) {
    type Alias Useractivityquery

    if UseractivityqueryMarshalled {
        return []byte("{}"), nil
    }
    UseractivityqueryMarshalled = true

    return json.Marshal(&struct {
        
        Metrics []Useractivityquerymetric `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Useractivityqueryfilter `json:"filter"`
        
        Order string `json:"order"`
        *Alias
    }{

        
        Metrics: []Useractivityquerymetric{{}},
        


        
        GroupBy: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

