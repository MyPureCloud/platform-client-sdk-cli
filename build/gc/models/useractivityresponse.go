package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityresponseDud struct { 
    


    

}

// Useractivityresponse
type Useractivityresponse struct { 
    // Results - Query results
    Results []Useractivitydata `json:"results"`


    // EntityIdDimension - Dimension that is used as an entityId
    EntityIdDimension string `json:"entityIdDimension"`

}

// String returns a JSON representation of the model
func (o *Useractivityresponse) String() string {
     o.Results = []Useractivitydata{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivityresponse) MarshalJSON() ([]byte, error) {
    type Alias Useractivityresponse

    if UseractivityresponseMarshalled {
        return []byte("{}"), nil
    }
    UseractivityresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Useractivitydata `json:"results"`
        
        EntityIdDimension string `json:"entityIdDimension"`
        *Alias
    }{

        
        Results: []Useractivitydata{{}},
        


        

        Alias: (*Alias)(u),
    })
}

