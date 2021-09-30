package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueueobservationqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueueobservationqueryresponseDud struct { 
    


    

}

// Queueobservationqueryresponse
type Queueobservationqueryresponse struct { 
    // SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
    SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`


    // Results
    Results []Queueobservationdatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Queueobservationqueryresponse) String() string {
    
    
     o.SystemToOrganizationMappings = map[string][]string{"": {}} 
    
    
    
     o.Results = []Queueobservationdatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queueobservationqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Queueobservationqueryresponse

    if QueueobservationqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    QueueobservationqueryresponseMarshalled = true

    return json.Marshal(&struct { 
        SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`
        
        Results []Queueobservationdatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        SystemToOrganizationMappings: map[string][]string{"": {}},
        

        

        
        Results: []Queueobservationdatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

