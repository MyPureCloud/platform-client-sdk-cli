package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IvrestablishedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IvrestablishedeventDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Ivrestablishedevent
type Ivrestablishedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication
    CommunicationId string `json:"communicationId"`


    // IvrPhoneNumber - The phone number for this IVR, if any is known
    IvrPhoneNumber string `json:"ivrPhoneNumber"`


    // IvrName - A displayable name for this IVR, if any is known.
    IvrName string `json:"ivrName"`


    // Ani - The automatic number identification if it is available for this conversation.
    Ani string `json:"ani"`


    // Dnis - The dialed number identification if it is available for this conversation.
    Dnis string `json:"dnis"`


    // InitialConfiguration - Metadata about this communication.
    InitialConfiguration Initialconfiguration `json:"initialConfiguration"`


    // SourceConfiguration - Metadata about the source of this communication's interaction.
    SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`

}

// String returns a JSON representation of the model
func (o *Ivrestablishedevent) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ivrestablishedevent) MarshalJSON() ([]byte, error) {
    type Alias Ivrestablishedevent

    if IvrestablishedeventMarshalled {
        return []byte("{}"), nil
    }
    IvrestablishedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        IvrPhoneNumber string `json:"ivrPhoneNumber"`
        
        IvrName string `json:"ivrName"`
        
        Ani string `json:"ani"`
        
        Dnis string `json:"dnis"`
        
        InitialConfiguration Initialconfiguration `json:"initialConfiguration"`
        
        SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

