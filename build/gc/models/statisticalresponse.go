package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StatisticalresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StatisticalresponseDud struct { 
    


    


    

}

// Statisticalresponse
type Statisticalresponse struct { 
    // Interval
    Interval string `json:"interval"`


    // Metrics
    Metrics []Aggregatemetricdata `json:"metrics"`


    // Views
    Views []Aggregateviewdata `json:"views"`

}

// String returns a JSON representation of the model
func (o *Statisticalresponse) String() string {
    
     o.Metrics = []Aggregatemetricdata{{}} 
     o.Views = []Aggregateviewdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Statisticalresponse) MarshalJSON() ([]byte, error) {
    type Alias Statisticalresponse

    if StatisticalresponseMarshalled {
        return []byte("{}"), nil
    }
    StatisticalresponseMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Metrics []Aggregatemetricdata `json:"metrics"`
        
        Views []Aggregateviewdata `json:"views"`
        *Alias
    }{

        


        
        Metrics: []Aggregatemetricdata{{}},
        


        
        Views: []Aggregateviewdata{{}},
        

        Alias: (*Alias)(u),
    })
}

