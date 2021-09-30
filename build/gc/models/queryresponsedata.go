package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryresponsedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryresponsedataDud struct { 
    


    

}

// Queryresponsedata
type Queryresponsedata struct { 
    // Interval - Interval with start and end represented as ISO-8601 string. i.e: yyyy-MM-dd'T'HH:mm:ss.SSS'Z'/yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
    Interval string `json:"interval"`


    // Metrics - A list of aggregated metrics
    Metrics []Queryresponsemetric `json:"metrics"`

}

// String returns a JSON representation of the model
func (o *Queryresponsedata) String() string {
    
    
    
    
    
    
     o.Metrics = []Queryresponsemetric{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryresponsedata) MarshalJSON() ([]byte, error) {
    type Alias Queryresponsedata

    if QueryresponsedataMarshalled {
        return []byte("{}"), nil
    }
    QueryresponsedataMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        Metrics []Queryresponsemetric `json:"metrics"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []Queryresponsemetric{{}},
        

        
        Alias: (*Alias)(u),
    })
}

