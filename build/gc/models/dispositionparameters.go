package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DispositionparametersMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DispositionparametersDud struct { 
    

}

// Dispositionparameters
type Dispositionparameters struct { 
    // AdjustableLiveSpeakerDetection - ALSD evaluation inputs and output (isPersonalLikely) of the ALSD detector the last time it ran on the call (could be multiple times)
    AdjustableLiveSpeakerDetection Adjustablelivespeakerdetection `json:"adjustableLiveSpeakerDetection"`

}

// String returns a JSON representation of the model
func (o *Dispositionparameters) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dispositionparameters) MarshalJSON() ([]byte, error) {
    type Alias Dispositionparameters

    if DispositionparametersMarshalled {
        return []byte("{}"), nil
    }
    DispositionparametersMarshalled = true

    return json.Marshal(&struct {
        
        AdjustableLiveSpeakerDetection Adjustablelivespeakerdetection `json:"adjustableLiveSpeakerDetection"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

