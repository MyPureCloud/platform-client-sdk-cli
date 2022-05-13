package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VoicemailcopyrecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VoicemailcopyrecordDud struct { 
    User User `json:"user"`


    Group Group `json:"group"`


    Date time.Time `json:"date"`

}

// Voicemailcopyrecord
type Voicemailcopyrecord struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Voicemailcopyrecord) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Voicemailcopyrecord) MarshalJSON() ([]byte, error) {
    type Alias Voicemailcopyrecord

    if VoicemailcopyrecordMarshalled {
        return []byte("{}"), nil
    }
    VoicemailcopyrecordMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

