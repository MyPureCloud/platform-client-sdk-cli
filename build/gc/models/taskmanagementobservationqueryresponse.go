package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationqueryresponseDud struct { 
    


    


    


    

}

// Taskmanagementobservationqueryresponse
type Taskmanagementobservationqueryresponse struct { 
    // Results - Query results grouped by the specified dimensions supplied in the groupBy parameter. Each result contains metrics for a specific group combination.
    Results []Taskmanagementobservationresult `json:"results"`


    // Details - Details about entities contained in results. Provides expanded information when requested through the expands parameter.
    Details Taskmanagementobservationdetailcontainer `json:"details"`


    // Cursors - Cursor tokens to be used for navigating paginated results
    Cursors Cursors `json:"cursors"`


    // NextUri - A URI to the next page in the listing.
    NextUri string `json:"nextUri"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationqueryresponse) String() string {
     o.Results = []Taskmanagementobservationresult{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationqueryresponse

    if TaskmanagementobservationqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Taskmanagementobservationresult `json:"results"`
        
        Details Taskmanagementobservationdetailcontainer `json:"details"`
        
        Cursors Cursors `json:"cursors"`
        
        NextUri string `json:"nextUri"`
        *Alias
    }{

        
        Results: []Taskmanagementobservationresult{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

