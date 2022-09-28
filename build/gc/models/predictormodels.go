package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictormodelsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictormodelsDud struct { 
    


    

}

// Predictormodels
type Predictormodels struct { 
    // Entities
    Entities []Predictormodel `json:"entities"`


    // PredictorModels
    PredictorModels []Predictormodel `json:"predictorModels"`

}

// String returns a JSON representation of the model
func (o *Predictormodels) String() string {
     o.Entities = []Predictormodel{{}} 
     o.PredictorModels = []Predictormodel{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictormodels) MarshalJSON() ([]byte, error) {
    type Alias Predictormodels

    if PredictormodelsMarshalled {
        return []byte("{}"), nil
    }
    PredictormodelsMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Predictormodel `json:"entities"`
        
        PredictorModels []Predictormodel `json:"predictorModels"`
        *Alias
    }{

        
        Entities: []Predictormodel{{}},
        


        
        PredictorModels: []Predictormodel{{}},
        

        Alias: (*Alias)(u),
    })
}

