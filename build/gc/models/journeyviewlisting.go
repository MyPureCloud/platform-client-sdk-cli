package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneyviewlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneyviewlistingDud struct { 
    


    


    

}

// Journeyviewlisting
type Journeyviewlisting struct { 
    // Total
    Total int `json:"total"`


    // Entities
    Entities []Journeyview `json:"entities"`


    // SelfUri
    SelfUri string `json:"selfUri"`

}

// String returns a JSON representation of the model
func (o *Journeyviewlisting) String() string {
    
     o.Entities = []Journeyview{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeyviewlisting) MarshalJSON() ([]byte, error) {
    type Alias Journeyviewlisting

    if JourneyviewlistingMarshalled {
        return []byte("{}"), nil
    }
    JourneyviewlistingMarshalled = true

    return json.Marshal(&struct {
        
        Total int `json:"total"`
        
        Entities []Journeyview `json:"entities"`
        
        SelfUri string `json:"selfUri"`
        *Alias
    }{

        


        
        Entities: []Journeyview{{}},
        


        

        Alias: (*Alias)(u),
    })
}

