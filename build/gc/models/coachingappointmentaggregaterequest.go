package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingappointmentaggregaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingappointmentaggregaterequestDud struct { 
    


    


    


    

}

// Coachingappointmentaggregaterequest
type Coachingappointmentaggregaterequest struct { 
    // Interval - Interval to aggregate across. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Metrics - A list of metrics to aggregate.  If omitted, all metrics are returned.
    Metrics []string `json:"metrics"`


    // GroupBy - An optional list of items by which to group the result data.
    GroupBy []string `json:"groupBy"`


    // Filter - The filter applied to the data
    Filter Queryrequestfilter `json:"filter"`

}

// String returns a JSON representation of the model
func (o *Coachingappointmentaggregaterequest) String() string {
    
    
    
    
    
    
     o.Metrics = []string{""} 
    
    
    
     o.GroupBy = []string{""} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingappointmentaggregaterequest) MarshalJSON() ([]byte, error) {
    type Alias Coachingappointmentaggregaterequest

    if CoachingappointmentaggregaterequestMarshalled {
        return []byte("{}"), nil
    }
    CoachingappointmentaggregaterequestMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        Metrics []string `json:"metrics"`
        
        GroupBy []string `json:"groupBy"`
        
        Filter Queryrequestfilter `json:"filter"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []string{""},
        

        

        
        GroupBy: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

