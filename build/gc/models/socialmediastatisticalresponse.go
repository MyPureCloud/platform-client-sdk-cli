package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialmediastatisticalresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialmediastatisticalresponseDud struct { 
    


    

}

// Socialmediastatisticalresponse
type Socialmediastatisticalresponse struct { 
    // Interval
    Interval string `json:"interval"`


    // Metrics
    Metrics []Socialmediaaggregatemetricdata `json:"metrics"`

}

// String returns a JSON representation of the model
func (o *Socialmediastatisticalresponse) String() string {
    
     o.Metrics = []Socialmediaaggregatemetricdata{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialmediastatisticalresponse) MarshalJSON() ([]byte, error) {
    type Alias Socialmediastatisticalresponse

    if SocialmediastatisticalresponseMarshalled {
        return []byte("{}"), nil
    }
    SocialmediastatisticalresponseMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Metrics []Socialmediaaggregatemetricdata `json:"metrics"`
        *Alias
    }{

        


        
        Metrics: []Socialmediaaggregatemetricdata{{}},
        

        Alias: (*Alias)(u),
    })
}

