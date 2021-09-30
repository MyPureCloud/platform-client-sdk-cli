package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseraggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseraggregatequeryresponseDud struct { 
    


    

}

// Useraggregatequeryresponse
type Useraggregatequeryresponse struct { 
    // SystemToOrganizationMappings - A mapping from system presence to a list of organization presence ids
    SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`


    // Results
    Results []Useraggregatedatacontainer `json:"results"`

}

// String returns a JSON representation of the model
func (o *Useraggregatequeryresponse) String() string {
    
    
     o.SystemToOrganizationMappings = map[string][]string{"": {}} 
    
    
    
     o.Results = []Useraggregatedatacontainer{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useraggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Useraggregatequeryresponse

    if UseraggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    UseraggregatequeryresponseMarshalled = true

    return json.Marshal(&struct { 
        SystemToOrganizationMappings map[string][]string `json:"systemToOrganizationMappings"`
        
        Results []Useraggregatedatacontainer `json:"results"`
        
        *Alias
    }{
        

        
        SystemToOrganizationMappings: map[string][]string{"": {}},
        

        

        
        Results: []Useraggregatedatacontainer{{}},
        

        
        Alias: (*Alias)(u),
    })
}

