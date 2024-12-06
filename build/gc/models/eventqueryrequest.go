package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventqueryrequestDud struct { 
    


    


    


    

}

// Eventqueryrequest
type Eventqueryrequest struct { 
    // Interval - Date and time range to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // EventDefinitionIds - Filter events by a list of event definition ids
    EventDefinitionIds []string `json:"eventDefinitionIds"`


    // SearchTerm - Only return events that contain the search term
    SearchTerm string `json:"searchTerm"`


    // SortOrder - Order of results. Default order is DESC.
    SortOrder string `json:"sortOrder"`

}

// String returns a JSON representation of the model
func (o *Eventqueryrequest) String() string {
    
     o.EventDefinitionIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Eventqueryrequest

    if EventqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    EventqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        EventDefinitionIds []string `json:"eventDefinitionIds"`
        
        SearchTerm string `json:"searchTerm"`
        
        SortOrder string `json:"sortOrder"`
        *Alias
    }{

        


        
        EventDefinitionIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

