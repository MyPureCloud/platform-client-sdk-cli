package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewelementattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewelementattributesDud struct { 
    


    


    

}

// Journeyviewelementattributes - Attributes on an element in a journey view
type Journeyviewelementattributes struct { 
    // VarType - The type of the element (e.g. Event)
    VarType string `json:"type"`


    // Id - The identifier for the element based on its type
    Id string `json:"id"`


    // Source - The source for the element (e.g. IVR, Voice, Chat). Used for informational purposes only
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Journeyviewelementattributes) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewelementattributes) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewelementattributes

    if JourneyviewelementattributesMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewelementattributesMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Id string `json:"id"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

