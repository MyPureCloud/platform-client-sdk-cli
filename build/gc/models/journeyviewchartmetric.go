package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartmetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartmetricDud struct { 
    


    


    


    

}

// Journeyviewchartmetric - A metric to measure within the chart
type Journeyviewchartmetric struct { 
    // Id - The unique identifier of the metric within the chart
    Id string `json:"id"`


    // ElementId - The element in the list of elements which the metric is measuring
    ElementId string `json:"elementId"`


    // Aggregate - How to aggregate the given element, defaults to EventCount
    Aggregate string `json:"aggregate"`


    // DisplayLabel - A display label for the metric
    DisplayLabel string `json:"displayLabel"`

}

// String returns a JSON representation of the model
func (o *Journeyviewchartmetric) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartmetric) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartmetric

    if JourneyviewchartmetricMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartmetricMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ElementId string `json:"elementId"`
        
        Aggregate string `json:"aggregate"`
        
        DisplayLabel string `json:"displayLabel"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

