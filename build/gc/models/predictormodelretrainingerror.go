package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictormodelretrainingerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictormodelretrainingerrorDud struct { 
    Id string `json:"id"`


    ErrorCode string `json:"errorCode"`


    DateOfFirstOccurrence time.Time `json:"dateOfFirstOccurrence"`

}

// Predictormodelretrainingerror
type Predictormodelretrainingerror struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Predictormodelretrainingerror) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictormodelretrainingerror) MarshalJSON() ([]byte, error) {
    type Alias Predictormodelretrainingerror

    if PredictormodelretrainingerrorMarshalled {
        return []byte("{}"), nil
    }
    PredictormodelretrainingerrorMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

