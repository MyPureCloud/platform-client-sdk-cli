package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationbasicMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationbasicDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    Participants []Participantbasic `json:"participants"`

}

// Conversationbasic
type Conversationbasic struct { 
    


    // Name
    Name string `json:"name"`


    // ExternalTag - The external tag associated with the conversation.
    ExternalTag string `json:"externalTag"`


    // StartTime - The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // Divisions - Identifiers of divisions associated with this conversation
    Divisions []Conversationdivisionmembership `json:"divisions"`


    


    

}

// String returns a JSON representation of the model
func (o *Conversationbasic) String() string {
    
    
    
    
     o.Divisions = []Conversationdivisionmembership{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationbasic) MarshalJSON() ([]byte, error) {
    type Alias Conversationbasic

    if ConversationbasicMarshalled {
        return []byte("{}"), nil
    }
    ConversationbasicMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ExternalTag string `json:"externalTag"`
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        Divisions []Conversationdivisionmembership `json:"divisions"`
        *Alias
    }{

        


        


        


        


        


        
        Divisions: []Conversationdivisionmembership{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

