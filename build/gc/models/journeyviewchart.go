package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartDud struct { 
    Id string `json:"id"`


    


    Version int `json:"version"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Journeyviewchart - A chart within the context of the elements of the the journey view
type Journeyviewchart struct { 
    


    // Name
    Name string `json:"name"`


    


    // GroupByTime - A time unit to group the metrics by. There is a limit on the number of groupBy properties which can be specified.
    GroupByTime string `json:"groupByTime"`


    // GroupByAttributes - A list of attributes to group the metrics by. There is a limit on the number of groupBy properties which can be specified.
    GroupByAttributes []Journeyviewchartgroupbyattribute `json:"groupByAttributes"`


    // Metrics - A list of metrics to calculate within the chart by (aka the y axis)
    Metrics []Journeyviewchartmetric `json:"metrics"`


    

}

// String returns a JSON representation of the model
func (o *Journeyviewchart) String() string {
    
    
     o.GroupByAttributes = []Journeyviewchartgroupbyattribute{{}} 
     o.Metrics = []Journeyviewchartmetric{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchart) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchart

    if JourneyviewchartMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        GroupByTime string `json:"groupByTime"`
        
        GroupByAttributes []Journeyviewchartgroupbyattribute `json:"groupByAttributes"`
        
        Metrics []Journeyviewchartmetric `json:"metrics"`
        *Alias
    }{

        


        


        


        


        
        GroupByAttributes: []Journeyviewchartgroupbyattribute{{}},
        


        
        Metrics: []Journeyviewchartmetric{{}},
        


        

        Alias: (*Alias)(u),
    })
}

