package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictormodelfeaturelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictormodelfeaturelistingDud struct { 
    

}

// Predictormodelfeaturelisting
type Predictormodelfeaturelisting struct { 
    // Entities
    Entities []Predictormodelfeature `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Predictormodelfeaturelisting) String() string {
     o.Entities = []Predictormodelfeature{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictormodelfeaturelisting) MarshalJSON() ([]byte, error) {
    type Alias Predictormodelfeaturelisting

    if PredictormodelfeaturelistingMarshalled {
        return []byte("{}"), nil
    }
    PredictormodelfeaturelistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Predictormodelfeature `json:"entities"`
        *Alias
    }{

        
        Entities: []Predictormodelfeature{{}},
        

        Alias: (*Alias)(u),
    })
}

