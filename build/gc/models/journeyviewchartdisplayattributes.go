package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartdisplayattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartdisplayattributesDud struct { 
    


    


    


    

}

// Journeyviewchartdisplayattributes - Display attributes for the chart, such as type, labels and legends
type Journeyviewchartdisplayattributes struct { 
    // VarType - The type of chart to display
    VarType string `json:"type"`


    // GroupByTitle - A title for the grouped by attributes (aka the x axis)
    GroupByTitle string `json:"groupByTitle"`


    // MetricsTitle - A title for the metrics (aka the y axis)
    MetricsTitle string `json:"metricsTitle"`


    // ShowLegend - Whether to show a legend
    ShowLegend bool `json:"showLegend"`

}

// String returns a JSON representation of the model
func (o *Journeyviewchartdisplayattributes) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartdisplayattributes) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartdisplayattributes

    if JourneyviewchartdisplayattributesMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartdisplayattributesMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        GroupByTitle string `json:"groupByTitle"`
        
        MetricsTitle string `json:"metricsTitle"`
        
        ShowLegend bool `json:"showLegend"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

