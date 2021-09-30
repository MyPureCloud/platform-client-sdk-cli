package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DevelopmentactivityaggregatequeryresponsedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DevelopmentactivityaggregatequeryresponsedataDud struct { 
    


    

}

// Developmentactivityaggregatequeryresponsedata
type Developmentactivityaggregatequeryresponsedata struct { 
    // Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Metrics - The list of aggregated metrics
    Metrics []Developmentactivityaggregatequeryresponsemetric `json:"metrics"`

}

// String returns a JSON representation of the model
func (o *Developmentactivityaggregatequeryresponsedata) String() string {
    
    
    
    
    
    
     o.Metrics = []Developmentactivityaggregatequeryresponsemetric{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Developmentactivityaggregatequeryresponsedata) MarshalJSON() ([]byte, error) {
    type Alias Developmentactivityaggregatequeryresponsedata

    if DevelopmentactivityaggregatequeryresponsedataMarshalled {
        return []byte("{}"), nil
    }
    DevelopmentactivityaggregatequeryresponsedataMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        Metrics []Developmentactivityaggregatequeryresponsemetric `json:"metrics"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []Developmentactivityaggregatequeryresponsemetric{{}},
        

        
        Alias: (*Alias)(u),
    })
}

