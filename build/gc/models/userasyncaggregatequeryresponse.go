package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserasyncaggregatequeryresponseDud struct { 
    


    


    

}

// Userasyncaggregatequeryresponse
type Userasyncaggregatequeryresponse struct { 
    // SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
    SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`


    // Results
    Results []Useraggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Userasyncaggregatequeryresponse) String() string {
     o.SystemToOrganizationMappings = map[string][]string{"": {}} 
     o.Results = []Useraggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Userasyncaggregatequeryresponse

    if UserasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    UserasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`
        
        Results []Useraggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        SystemToOrganizationMappings: map[string][]string{"": {}},
        


        
        Results: []Useraggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

