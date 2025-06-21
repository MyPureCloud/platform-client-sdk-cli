package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InternalmessagedataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InternalmessagedataDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    NormalizedMessage Conversationnormalizedmessage `json:"normalizedMessage"`


    SelfUri string `json:"selfUri"`

}

// Internalmessagedata
type Internalmessagedata struct { 
    


    // Name
    Name string `json:"name"`


    // Conversation - The conversation of this message.
    Conversation Addressableentityref `json:"conversation"`


    // CommunicationId - The id of the communication of this message.
    CommunicationId string `json:"communicationId"`


    // Timestamp - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Sender - The sender of the text message.
    Sender Userreference `json:"sender"`


    // Recipient - The recipient of the text message.
    Recipient Userreference `json:"recipient"`


    


    

}

// String returns a JSON representation of the model
func (o *Internalmessagedata) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Internalmessagedata) MarshalJSON() ([]byte, error) {
    type Alias Internalmessagedata

    if InternalmessagedataMarshalled {
        return []byte("{}"), nil
    }
    InternalmessagedataMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Conversation Addressableentityref `json:"conversation"`
        
        CommunicationId string `json:"communicationId"`
        
        Timestamp time.Time `json:"timestamp"`
        
        Sender Userreference `json:"sender"`
        
        Recipient Userreference `json:"recipient"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

