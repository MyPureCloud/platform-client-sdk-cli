package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummaryaggregationqueryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummaryaggregationqueryDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Summaryaggregationquery
type Summaryaggregationquery struct { 
    // Interval - Behaves like one clause in a SQL WHERE. Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Granularity - Granularity aggregates metrics into subpartitions within the time interval specified. The default granularity is the same duration as the interval. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
    Granularity string `json:"granularity"`


    // TimeZone - Time zone context used to calculate response intervals (this allows resolving DST changes). The interval offset is used even when timeZone is specified. Default is UTC. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZone string `json:"timeZone"`


    // GroupBy - Behaves like a SQL GROUPBY. Allows for multiple levels of grouping as a list of dimensions. Partitions resulting aggregate computations into distinct named subgroups rather than across the entire result set as if it were one group.
    GroupBy []string `json:"groupBy"`


    // Filter - Behaves like a SQL WHERE clause. This is ANDed with the interval parameter. Expresses boolean logical predicates as well as dimensional filters
    Filter Summaryaggregatequeryfilter `json:"filter"`


    // Metrics - Behaves like a SQL SELECT clause. Only named metrics will be retrieved.
    Metrics []string `json:"metrics"`


    // FlattenMultivaluedDimensions - Flattens any multivalued dimensions used in response groups (e.g. ['a','b','c']->'a,b,c')
    FlattenMultivaluedDimensions bool `json:"flattenMultivaluedDimensions"`


    // Views - Custom derived metric views
    Views []Summaryaggregationview `json:"views"`


    // AlternateTimeDimension - Dimension to use as the alternative timestamp for data in the aggregate.  Choosing \"eventTime\" uses the actual time of the data event.
    AlternateTimeDimension string `json:"alternateTimeDimension"`


    // QueryType - Query type to use. Use groupBy for all matching results, and topN for just top N results for the requested metric (group by exactly 1 dimension)
    QueryType string `json:"queryType"`


    // Limit - How many results you want in the topN list. Only applicable for topN query type.
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Summaryaggregationquery) String() string {
    
    
    
     o.GroupBy = []string{""} 
    
     o.Metrics = []string{""} 
    
     o.Views = []Summaryaggregationview{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summaryaggregationquery) MarshalJSON() ([]byte, error) {
    type Alias Summaryaggregationquery

    if SummaryaggregationqueryMarshalled {
        return []byte("{}"), nil
    }
    SummaryaggregationqueryMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Granularity string `json:"granularity"`
        
        TimeZone string `json:"timeZone"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Summaryaggregatequeryfilter `json:"filter"`
        
        Metrics []string `json:"metrics"`
        
        FlattenMultivaluedDimensions bool `json:"flattenMultivaluedDimensions"`
        
        Views []Summaryaggregationview `json:"views"`
        
        AlternateTimeDimension string `json:"alternateTimeDimension"`
        
        QueryType string `json:"queryType"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        


        


        
        GroupBy: []string{""},
        


        


        
        Metrics: []string{""},
        


        


        
        Views: []Summaryaggregationview{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

