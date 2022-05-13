package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediaparticipantrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediaparticipantrequestDud struct { 
    


    


    


    


    


    


    

}

// Mediaparticipantrequest
type Mediaparticipantrequest struct { 
    // Wrapup - Wrap-up to assign to this participant.
    Wrapup Wrapup `json:"wrapup"`


    // State - The state to update to set for this participant's communications.  Possible values are: 'connected' and 'disconnected'.
    State string `json:"state"`


    // Recording - True to enable recording of this participant, otherwise false to disable recording.
    Recording bool `json:"recording"`


    // Muted - True to mute this conversation participant.
    Muted bool `json:"muted"`


    // Confined - True to confine this conversation participant.  Should only be used for ad-hoc conferences
    Confined bool `json:"confined"`


    // Held - True to hold this conversation participant.
    Held bool `json:"held"`


    // WrapupSkipped - True to skip wrap-up for this participant.
    WrapupSkipped bool `json:"wrapupSkipped"`

}

// String returns a JSON representation of the model
func (o *Mediaparticipantrequest) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediaparticipantrequest) MarshalJSON() ([]byte, error) {
    type Alias Mediaparticipantrequest

    if MediaparticipantrequestMarshalled {
        return []byte("{}"), nil
    }
    MediaparticipantrequestMarshalled = true

    return json.Marshal(&struct {
        
        Wrapup Wrapup `json:"wrapup"`
        
        State string `json:"state"`
        
        Recording bool `json:"recording"`
        
        Muted bool `json:"muted"`
        
        Confined bool `json:"confined"`
        
        Held bool `json:"held"`
        
        WrapupSkipped bool `json:"wrapupSkipped"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

