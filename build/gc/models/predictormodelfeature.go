package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictormodelfeatureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictormodelfeatureDud struct { 
    Id string `json:"id"`


    VarType string `json:"type"`


    PercentageImportance float64 `json:"percentageImportance"`

}

// Predictormodelfeature
type Predictormodelfeature struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Predictormodelfeature) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictormodelfeature) MarshalJSON() ([]byte, error) {
    type Alias Predictormodelfeature

    if PredictormodelfeatureMarshalled {
        return []byte("{}"), nil
    }
    PredictormodelfeatureMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

