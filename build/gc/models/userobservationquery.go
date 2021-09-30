package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserobservationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserobservationqueryDud struct { 
    


    


    

}

// Userobservationquery
type Userobservationquery struct { 
    // Filter - Filter to return a subset of observations. Expresses boolean logical predicates as well as dimensional filters
    Filter Userobservationqueryfilter `json:"filter"`


    // Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
    Metrics []string `json:"metrics"`


    // DetailMetrics - Metrics for which to include additional detailed observations
    DetailMetrics []string `json:"detailMetrics"`

}

// String returns a JSON representation of the model
func (o *Userobservationquery) String() string {
    
    
    
    
    
    
     o.Metrics = []string{""} 
    
    
    
     o.DetailMetrics = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userobservationquery) MarshalJSON() ([]byte, error) {
    type Alias Userobservationquery

    if UserobservationqueryMarshalled {
        return []byte("{}"), nil
    }
    UserobservationqueryMarshalled = true

    return json.Marshal(&struct { 
        Filter Userobservationqueryfilter `json:"filter"`
        
        Metrics []string `json:"metrics"`
        
        DetailMetrics []string `json:"detailMetrics"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []string{""},
        

        

        
        DetailMetrics: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

