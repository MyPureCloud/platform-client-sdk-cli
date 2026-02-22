package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributesbulkupdateresponsemapMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributesbulkupdateresponsemapDud struct { 
    

}

// Customattributesbulkupdateresponsemap
type Customattributesbulkupdateresponsemap struct { 
    // Results - The map of Custom Attributes record ids to their results after updating.
    Results map[string]Customattributesbulkupdateresponse `json:"results"`

}

// String returns a JSON representation of the model
func (o *Customattributesbulkupdateresponsemap) String() string {
     o.Results = map[string]Customattributesbulkupdateresponse{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributesbulkupdateresponsemap) MarshalJSON() ([]byte, error) {
    type Alias Customattributesbulkupdateresponsemap

    if CustomattributesbulkupdateresponsemapMarshalled {
        return []byte("{}"), nil
    }
    CustomattributesbulkupdateresponsemapMarshalled = true

    return json.Marshal(&struct {
        
        Results map[string]Customattributesbulkupdateresponse `json:"results"`
        *Alias
    }{

        
        Results: map[string]Customattributesbulkupdateresponse{"": {}},
        

        Alias: (*Alias)(u),
    })
}

