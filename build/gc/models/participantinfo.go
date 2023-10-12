package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantinfoDud struct { 
    


    

}

// Participantinfo
type Participantinfo struct { 
    // ActiveParticipantCount - The number of active participants in the video conference.
    ActiveParticipantCount int `json:"activeParticipantCount"`


    // Version - The version of the participant count.
    Version int `json:"version"`

}

// String returns a JSON representation of the model
func (o *Participantinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participantinfo) MarshalJSON() ([]byte, error) {
    type Alias Participantinfo

    if ParticipantinfoMarshalled {
        return []byte("{}"), nil
    }
    ParticipantinfoMarshalled = true

    return json.Marshal(&struct {
        
        ActiveParticipantCount int `json:"activeParticipantCount"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

