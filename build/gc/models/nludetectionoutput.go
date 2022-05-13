package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludetectionoutputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludetectionoutputDud struct { 
    Intents []Detectedintent `json:"intents"`


    DialogActs []Detecteddialogact `json:"dialogActs"`

}

// Nludetectionoutput
type Nludetectionoutput struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Nludetectionoutput) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludetectionoutput) MarshalJSON() ([]byte, error) {
    type Alias Nludetectionoutput

    if NludetectionoutputMarshalled {
        return []byte("{}"), nil
    }
    NludetectionoutputMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

