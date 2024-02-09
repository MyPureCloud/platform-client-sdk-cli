package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RatelimitaggregatedatacontainerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RatelimitaggregatedatacontainerDud struct { 
    


    

}

// Ratelimitaggregatedatacontainer
type Ratelimitaggregatedatacontainer struct { 
    // Group - A mapping from dimension to value
    Group map[string]string `json:"group"`


    // Data
    Data []Statisticalresponse `json:"data"`

}

// String returns a JSON representation of the model
func (o *Ratelimitaggregatedatacontainer) String() string {
     o.Group = map[string]string{"": ""} 
     o.Data = []Statisticalresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ratelimitaggregatedatacontainer) MarshalJSON() ([]byte, error) {
    type Alias Ratelimitaggregatedatacontainer

    if RatelimitaggregatedatacontainerMarshalled {
        return []byte("{}"), nil
    }
    RatelimitaggregatedatacontainerMarshalled = true

    return json.Marshal(&struct {
        
        Group map[string]string `json:"group"`
        
        Data []Statisticalresponse `json:"data"`
        *Alias
    }{

        
        Group: map[string]string{"": ""},
        


        
        Data: []Statisticalresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

