package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementdisplayattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementdisplayattributesDud struct { 
    


    


    

}

// Journeyviewelementdisplayattributes - Display attributes for an element in a journey view
type Journeyviewelementdisplayattributes struct { 
    // X - The horizontal position (x-coordinate) of the element on the journey view canvas
    X int `json:"x"`


    // Y - The vertical position (y-coordinate) of the element on the journey view canvas
    Y int `json:"y"`


    // Col - The column position for the element in the journey view canvas
    Col int `json:"col"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementdisplayattributes) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementdisplayattributes) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementdisplayattributes

    if JourneyviewelementdisplayattributesMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementdisplayattributesMarshalled = true

    return json.Marshal(&struct {
        
        X int `json:"x"`
        
        Y int `json:"y"`
        
        Col int `json:"col"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

