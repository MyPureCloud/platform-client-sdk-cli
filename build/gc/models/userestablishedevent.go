package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserestablishedeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserestablishedeventDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Userestablishedevent
type Userestablishedevent struct { 
    // EventId - A unique (V4 UUID) eventId for this event
    EventId string `json:"eventId"`


    // EventDateTime - A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDateTime time.Time `json:"eventDateTime"`


    // ConversationId - A unique Id (V4 UUID) identifying this conversation
    ConversationId string `json:"conversationId"`


    // CommunicationId - A unique Id (V4 UUID) identifying this communication
    CommunicationId string `json:"communicationId"`


    // PhoneNumber - Identifies the phone number used to reach this user if it is different from the information that would be accessed by userId.
    PhoneNumber string `json:"phoneNumber"`


    // UserId - The userId (V4 UUID) for the user this communication belongs to.
    UserId string `json:"userId"`


    // StationId - A Station ID (V4 UUID) that identifies the station being used if the user is using a station and the stationId is known.
    StationId string `json:"stationId"`


    // AfterCallWorkRequired - Indicates whether or not this user will be required to complete after call work.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`


    // QueueId - The id (V4 UUID) of the queue that the user is calling on behalf of. Applies to outbound calls only.
    QueueId string `json:"queueId"`


    // InitialConfiguration - Metadata about this communication.
    InitialConfiguration Initialconfiguration `json:"initialConfiguration"`


    // SourceConfiguration - Metadata about the source of this communication's interaction.
    SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`

}

// String returns a JSON representation of the model
func (o *Userestablishedevent) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userestablishedevent) MarshalJSON() ([]byte, error) {
    type Alias Userestablishedevent

    if UserestablishedeventMarshalled {
        return []byte("{}"), nil
    }
    UserestablishedeventMarshalled = true

    return json.Marshal(&struct {
        
        EventId string `json:"eventId"`
        
        EventDateTime time.Time `json:"eventDateTime"`
        
        ConversationId string `json:"conversationId"`
        
        CommunicationId string `json:"communicationId"`
        
        PhoneNumber string `json:"phoneNumber"`
        
        UserId string `json:"userId"`
        
        StationId string `json:"stationId"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        
        QueueId string `json:"queueId"`
        
        InitialConfiguration Initialconfiguration `json:"initialConfiguration"`
        
        SourceConfiguration Sourceconfiguration `json:"sourceConfiguration"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

