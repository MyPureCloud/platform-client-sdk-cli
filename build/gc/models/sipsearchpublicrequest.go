package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SipsearchpublicrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SipsearchpublicrequestDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Sipsearchpublicrequest
type Sipsearchpublicrequest struct { 
    


    // Name
    Name string `json:"name"`


    // CallId - unique identification of the placed call
    CallId string `json:"callId"`


    // ToUser - SIP user to who the call was placed
    ToUser string `json:"toUser"`


    // FromUser - SIP user who placed the call
    FromUser string `json:"fromUser"`


    // ConversationId - Unique identification of the conversation
    ConversationId string `json:"conversationId"`


    // ParticipantId - Unique identification of the participant
    ParticipantId string `json:"participantId"`


    // DateStart - Start date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStart time.Time `json:"dateStart"`


    // DateEnd - End date of the search. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateEnd time.Time `json:"dateEnd"`


    

}

// String returns a JSON representation of the model
func (o *Sipsearchpublicrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sipsearchpublicrequest) MarshalJSON() ([]byte, error) {
    type Alias Sipsearchpublicrequest

    if SipsearchpublicrequestMarshalled {
        return []byte("{}"), nil
    }
    SipsearchpublicrequestMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        CallId string `json:"callId"`
        
        ToUser string `json:"toUser"`
        
        FromUser string `json:"fromUser"`
        
        ConversationId string `json:"conversationId"`
        
        ParticipantId string `json:"participantId"`
        
        DateStart time.Time `json:"dateStart"`
        
        DateEnd time.Time `json:"dateEnd"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

