package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamactivityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamactivityresponseDud struct { 
    


    

}

// Teamactivityresponse
type Teamactivityresponse struct { 
    // Results - Query results
    Results []Teamactivitydata `json:"results"`


    // EntityIdDimension - Dimension that is used as an entityId
    EntityIdDimension string `json:"entityIdDimension"`

}

// String returns a JSON representation of the model
func (o *Teamactivityresponse) String() string {
     o.Results = []Teamactivitydata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamactivityresponse) MarshalJSON() ([]byte, error) {
    type Alias Teamactivityresponse

    if TeamactivityresponseMarshalled {
        return []byte("{}"), nil
    }
    TeamactivityresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Teamactivitydata `json:"results"`
        
        EntityIdDimension string `json:"entityIdDimension"`
        *Alias
    }{

        
        Results: []Teamactivitydata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

