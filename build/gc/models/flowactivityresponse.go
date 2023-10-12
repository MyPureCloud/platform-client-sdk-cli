package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowactivityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowactivityresponseDud struct { 
    


    

}

// Flowactivityresponse
type Flowactivityresponse struct { 
    // Results - Query results
    Results []Flowactivitydata `json:"results"`


    // EntityIdDimension - Dimension that is used as an entityId
    EntityIdDimension string `json:"entityIdDimension"`

}

// String returns a JSON representation of the model
func (o *Flowactivityresponse) String() string {
     o.Results = []Flowactivitydata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowactivityresponse) MarshalJSON() ([]byte, error) {
    type Alias Flowactivityresponse

    if FlowactivityresponseMarshalled {
        return []byte("{}"), nil
    }
    FlowactivityresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Flowactivitydata `json:"results"`
        
        EntityIdDimension string `json:"entityIdDimension"`
        *Alias
    }{

        
        Results: []Flowactivitydata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

