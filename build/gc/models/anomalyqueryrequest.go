package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyqueryrequestDud struct { 
    


    


    


    

}

// Anomalyqueryrequest
type Anomalyqueryrequest struct { 
    // Interval - Specifies the date and time range of data being queried. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Filters - Expresses boolean logical predicates as well as dimensional filters
    Filters Anomalyqueryfilters `json:"filters"`


    // PageNumber - The page number to retrieve
    PageNumber int `json:"pageNumber"`


    // PageSize - The number of items to return per page
    PageSize int `json:"pageSize"`

}

// String returns a JSON representation of the model
func (o *Anomalyqueryrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Anomalyqueryrequest

    if AnomalyqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    AnomalyqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Filters Anomalyqueryfilters `json:"filters"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

