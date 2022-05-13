package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationchatMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationchatDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Conversationchat
type Conversationchat struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // RoomId - The room id for the chat.
    RoomId string `json:"roomId"`


    // RecordingId - A globally unique identifier for the recording associated with this chat.
    RecordingId string `json:"recordingId"`


    // Segments - The time line of the participant's chat, divided into activity segments.
    Segments []Segment `json:"segments"`


    // Held - True if this call is held and the person on this side hears silence.
    Held bool `json:"held"`


    // Direction - The direction of the chat
    Direction string `json:"direction"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartHoldTime - The timestamp the chat was placed on hold in the cloud clock if the chat is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // Provider - The source provider for the email.
    Provider string `json:"provider"`


    // ScriptId - The UUID of the script to use.
    ScriptId string `json:"scriptId"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // AvatarImageUrl - If available, the URI to the avatar image of this communication.
    AvatarImageUrl string `json:"avatarImageUrl"`


    // JourneyContext - A subset of the Journey System's data relevant to a part of a conversation (for external linkage and internal usage/context).
    JourneyContext Journeycontext `json:"journeyContext"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`

}

// String returns a JSON representation of the model
func (o *Conversationchat) String() string {
    
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationchat) MarshalJSON() ([]byte, error) {
    type Alias Conversationchat

    if ConversationchatMarshalled {
        return []byte("{}"), nil
    }
    ConversationchatMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Id string `json:"id"`
        
        RoomId string `json:"roomId"`
        
        RecordingId string `json:"recordingId"`
        
        Segments []Segment `json:"segments"`
        
        Held bool `json:"held"`
        
        Direction string `json:"direction"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        Provider string `json:"provider"`
        
        ScriptId string `json:"scriptId"`
        
        PeerId string `json:"peerId"`
        
        AvatarImageUrl string `json:"avatarImageUrl"`
        
        JourneyContext Journeycontext `json:"journeyContext"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        *Alias
    }{

        


        


        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

