package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ApiusageorganizationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ApiusageorganizationqueryDud struct { 
    


    


    


    

}

// Apiusageorganizationquery
type Apiusageorganizationquery struct { 
    // Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Granularity - Date granularity of the results
    Granularity string `json:"granularity"`


    // Metrics - Behaves like a SQL SELECT clause. Enables retrieving only named metrics. If omitted, all metrics that are available will be returned (like SELECT *).
    Metrics []string `json:"metrics"`


    // GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
    GroupBy []string `json:"groupBy"`

}

// String returns a JSON representation of the model
func (o *Apiusageorganizationquery) String() string {
    
    
     o.Metrics = []string{""} 
     o.GroupBy = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Apiusageorganizationquery) MarshalJSON() ([]byte, error) {
    type Alias Apiusageorganizationquery

    if ApiusageorganizationqueryMarshalled {
        return []byte("{}"), nil
    }
    ApiusageorganizationqueryMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Granularity string `json:"granularity"`
        
        Metrics []string `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        *Alias
    }{

        


        


        
        Metrics: []string{""},
        


        
        GroupBy: []string{""},
        

        Alias: (*Alias)(u),
    })
}

