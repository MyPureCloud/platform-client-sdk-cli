package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScreenrecordingsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScreenrecordingsessionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Screenrecordingsession
type Screenrecordingsession struct { 
    


    // Name
    Name string `json:"name"`


    // User
    User User `json:"user"`


    // CommunicationId - The id of the communication that is being recorded on the conversation
    CommunicationId string `json:"communicationId"`


    // Conversation
    Conversation Conversation `json:"conversation"`


    // StartTime - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    

}

// String returns a JSON representation of the model
func (o *Screenrecordingsession) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Screenrecordingsession) MarshalJSON() ([]byte, error) {
    type Alias Screenrecordingsession

    if ScreenrecordingsessionMarshalled {
        return []byte("{}"), nil
    }
    ScreenrecordingsessionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        User User `json:"user"`
        
        CommunicationId string `json:"communicationId"`
        
        Conversation Conversation `json:"conversation"`
        
        StartTime time.Time `json:"startTime"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

