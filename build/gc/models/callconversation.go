package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallconversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallconversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Callconversation
type Callconversation struct { 
    


    // Name
    Name string `json:"name"`


    // Participants - The list of participants involved in the conversation.
    Participants []Callmediaparticipant `json:"participants"`


    // OtherMediaUris - The list of other media channels involved in the conversation.
    OtherMediaUris []string `json:"otherMediaUris"`


    // RecordingState
    RecordingState string `json:"recordingState"`


    // MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
    MaxParticipants int `json:"maxParticipants"`


    

}

// String returns a JSON representation of the model
func (o *Callconversation) String() string {
    
     o.Participants = []Callmediaparticipant{{}} 
     o.OtherMediaUris = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callconversation) MarshalJSON() ([]byte, error) {
    type Alias Callconversation

    if CallconversationMarshalled {
        return []byte("{}"), nil
    }
    CallconversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Participants []Callmediaparticipant `json:"participants"`
        
        OtherMediaUris []string `json:"otherMediaUris"`
        
        RecordingState string `json:"recordingState"`
        
        MaxParticipants int `json:"maxParticipants"`
        *Alias
    }{

        


        


        
        Participants: []Callmediaparticipant{{}},
        


        
        OtherMediaUris: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

