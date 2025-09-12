package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataextractionfileurllistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataextractionfileurllistingDud struct { 
    

}

// Dataextractionfileurllisting
type Dataextractionfileurllisting struct { 
    // Entities
    Entities []Dataextractionfileurl `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Dataextractionfileurllisting) String() string {
     o.Entities = []Dataextractionfileurl{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataextractionfileurllisting) MarshalJSON() ([]byte, error) {
    type Alias Dataextractionfileurllisting

    if DataextractionfileurllistingMarshalled {
        return []byte("{}"), nil
    }
    DataextractionfileurllistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dataextractionfileurl `json:"entities"`
        *Alias
    }{

        
        Entities: []Dataextractionfileurl{{}},
        

        Alias: (*Alias)(u),
    })
}

