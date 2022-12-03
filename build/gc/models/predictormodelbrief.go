package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictormodelbriefMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictormodelbriefDud struct { 
    MediaType string `json:"mediaType"`


    DateModified time.Time `json:"dateModified"`


    RetrainingErrors []Predictormodelretrainingerror `json:"retrainingErrors"`


    State string `json:"state"`

}

// Predictormodelbrief
type Predictormodelbrief struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Predictormodelbrief) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictormodelbrief) MarshalJSON() ([]byte, error) {
    type Alias Predictormodelbrief

    if PredictormodelbriefMarshalled {
        return []byte("{}"), nil
    }
    PredictormodelbriefMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

