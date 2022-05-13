package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingjobfailedrecordingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingjobfailedrecordingDud struct { 
    Conversation Addressableentityref `json:"conversation"`


    Recording Addressableentityref `json:"recording"`

}

// Recordingjobfailedrecording
type Recordingjobfailedrecording struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Recordingjobfailedrecording) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingjobfailedrecording) MarshalJSON() ([]byte, error) {
    type Alias Recordingjobfailedrecording

    if RecordingjobfailedrecordingMarshalled {
        return []byte("{}"), nil
    }
    RecordingjobfailedrecordingMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

