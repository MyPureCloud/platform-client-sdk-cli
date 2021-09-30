package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryresponsedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryresponsedataDud struct { 
    


    

}

// Learningassignmentaggregatequeryresponsedata
type Learningassignmentaggregatequeryresponsedata struct { 
    // Interval - Specifies the range of due dates to be used for filtering. A maximum of 1 year can be specified in the range. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Metrics - The list of aggregated metrics
    Metrics []Learningassignmentaggregatequeryresponsemetric `json:"metrics"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsedata) String() string {
    
    
    
    
    
    
     o.Metrics = []Learningassignmentaggregatequeryresponsemetric{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryresponsedata) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryresponsedata

    if LearningassignmentaggregatequeryresponsedataMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryresponsedataMarshalled = true

    return json.Marshal(&struct { 
        Interval string `json:"interval"`
        
        Metrics []Learningassignmentaggregatequeryresponsemetric `json:"metrics"`
        
        *Alias
    }{
        

        

        

        
        Metrics: []Learningassignmentaggregatequeryresponsemetric{{}},
        

        
        Alias: (*Alias)(u),
    })
}

