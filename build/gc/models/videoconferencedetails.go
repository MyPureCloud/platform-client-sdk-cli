package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    VideoconferencedetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type VideoconferencedetailsDud struct { 
    


    


    

}

// Videoconferencedetails
type Videoconferencedetails struct { 
    // ConferenceId - The conferenceId.
    ConferenceId string `json:"conferenceId"`


    // ConversationId - The conversationId of the video conference.
    ConversationId string `json:"conversationId"`


    // ParticipantInfo - Information about the participants of the video conference.
    ParticipantInfo Participantinfo `json:"participantInfo"`

}

// String returns a JSON representation of the model
func (o *Videoconferencedetails) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Videoconferencedetails) MarshalJSON() ([]byte, error) {
    type Alias Videoconferencedetails

    if VideoconferencedetailsMarshalled {
        return []byte("{}"), nil
    }
    VideoconferencedetailsMarshalled = true

    return json.Marshal(&struct {
        
        ConferenceId string `json:"conferenceId"`
        
        ConversationId string `json:"conversationId"`
        
        ParticipantInfo Participantinfo `json:"participantInfo"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

