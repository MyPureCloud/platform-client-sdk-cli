package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewchartgroupbyattributeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewchartgroupbyattributeDud struct { 
    


    

}

// Journeyviewchartgroupbyattribute - A journey element attribute to group by within the chart
type Journeyviewchartgroupbyattribute struct { 
    // ElementId - The element in the list of elements which is being grouped by
    ElementId string `json:"elementId"`


    // Attribute - The attribute of the element being grouped by
    Attribute string `json:"attribute"`

}

// String returns a JSON representation of the model
func (o *Journeyviewchartgroupbyattribute) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewchartgroupbyattribute) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewchartgroupbyattribute

    if JourneyviewchartgroupbyattributeMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewchartgroupbyattributeMarshalled = true

    return json.Marshal(&struct {
        
        ElementId string `json:"elementId"`
        
        Attribute string `json:"attribute"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

