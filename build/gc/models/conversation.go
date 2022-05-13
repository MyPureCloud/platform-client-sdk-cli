package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Conversation
type Conversation struct { 
    


    // Name
    Name string `json:"name"`


    // ExternalTag - The external tag associated with the conversation.
    ExternalTag string `json:"externalTag"`


    // StartTime - The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // Address - The address of the conversation as seen from an external participant. For phone calls this will be the DNIS for inbound calls and the ANI for outbound calls. For other media types this will be the address of the destination participant for inbound and the address of the initiating participant for outbound.
    Address string `json:"address"`


    // Participants - The list of all participants in the conversation.
    Participants []Participant `json:"participants"`


    // ConversationIds - A list of conversations to merge into this conversation to create a conference. This field is null except when being used to create a conference.
    ConversationIds []string `json:"conversationIds"`


    // MaxParticipants - If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
    MaxParticipants int `json:"maxParticipants"`


    // RecordingState - On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings; otherwise indicates state of conversation recording.
    RecordingState string `json:"recordingState"`


    // State - The conversation's state
    State string `json:"state"`


    // Divisions - Identifiers of divisions associated with this conversation
    Divisions []Conversationdivisionmembership `json:"divisions"`


    

}

// String returns a JSON representation of the model
func (o *Conversation) String() string {
    
    
    
    
    
     o.Participants = []Participant{{}} 
     o.ConversationIds = []string{""} 
    
    
    
     o.Divisions = []Conversationdivisionmembership{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversation) MarshalJSON() ([]byte, error) {
    type Alias Conversation

    if ConversationMarshalled {
        return []byte("{}"), nil
    }
    ConversationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ExternalTag string `json:"externalTag"`
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        Address string `json:"address"`
        
        Participants []Participant `json:"participants"`
        
        ConversationIds []string `json:"conversationIds"`
        
        MaxParticipants int `json:"maxParticipants"`
        
        RecordingState string `json:"recordingState"`
        
        State string `json:"state"`
        
        Divisions []Conversationdivisionmembership `json:"divisions"`
        *Alias
    }{

        


        


        


        


        


        


        
        Participants: []Participant{{}},
        


        
        ConversationIds: []string{""},
        


        


        


        


        
        Divisions: []Conversationdivisionmembership{{}},
        


        

        Alias: (*Alias)(u),
    })
}

