package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationaggregationquerymeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationaggregationquerymeDud struct { 
    


    


    


    


    


    

}

// Evaluationaggregationqueryme
type Evaluationaggregationqueryme struct { 
    // Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    // GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
    GroupBy []string `json:"groupBy"`


    // Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
    Metrics []string `json:"metrics"`


    // AlternateTimeDimension - Dimension to use as the alternative timestamp for data in the aggregate.  Choosing \"eventTime\" uses the actual time of the data event.
    AlternateTimeDimension string `json:"alternateTimeDimension"`


    // ContextId - Evaluation context Id
    ContextId string `json:"contextId"`

}

// String returns a JSON representation of the model
func (o *Evaluationaggregationqueryme) String() string {
    
    
     o.GroupBy = []string{""} 
     o.Metrics = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationaggregationqueryme) MarshalJSON() ([]byte, error) {
    type Alias Evaluationaggregationqueryme

    if EvaluationaggregationquerymeMarshalled {
        return []byte("{}"), nil
    }
    EvaluationaggregationquerymeMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        TimeZone string `json:"timeZone"`
        
        GroupBy []string `json:"groupBy"`
        
        Metrics []string `json:"metrics"`
        
        AlternateTimeDimension string `json:"alternateTimeDimension"`
        
        ContextId string `json:"contextId"`
        *Alias
    }{

        


        


        
        GroupBy: []string{""},
        


        
        Metrics: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

