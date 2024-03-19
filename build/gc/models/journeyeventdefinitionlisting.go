package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyeventdefinitionlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyeventdefinitionlistingDud struct { 
    


    


    

}

// Journeyeventdefinitionlisting
type Journeyeventdefinitionlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Journeyeventdefinition `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Journeyeventdefinitionlisting) String() string {
    
     o.Entities = []Journeyeventdefinition{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyeventdefinitionlisting) MarshalJSON() ([]byte, error) {
    type Alias Journeyeventdefinitionlisting

    if JourneyeventdefinitionlistingMarshalled {
        return []byte("{}"), nil
    }
    JourneyeventdefinitionlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Journeyeventdefinition `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Journeyeventdefinition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

