package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Message
type Message struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // InitialState - The initial connection state of this communication.
    InitialState string `json:"initialState"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Held - True if this call is held and the person on this side hears silence.
    Held bool `json:"held"`


    // Segments - The time line of the participant's message, divided into activity segments.
    Segments []Segment `json:"segments"`


    // Direction - The direction of the message.
    Direction string `json:"direction"`


    // RecordingId - A globally unique identifier for the recording associated with this message.
    RecordingId string `json:"recordingId"`


    // ErrorInfo
    ErrorInfo Errorbody `json:"errorInfo"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartHoldTime - The timestamp the message was placed on hold in the cloud clock if the message is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // Provider - The source provider for the message.
    Provider string `json:"provider"`


    // Authenticated - If true, the participant member is authenticated.
    Authenticated bool `json:"authenticated"`


    // VarType - Indicates the type of message platform from which the message originated.
    VarType string `json:"type"`


    // RecipientCountry - Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
    RecipientCountry string `json:"recipientCountry"`


    // RecipientType - The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
    RecipientType string `json:"recipientType"`


    // ScriptId - The UUID of the script to use.
    ScriptId string `json:"scriptId"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // ToAddress - Address and name data for a call endpoint.
    ToAddress Address `json:"toAddress"`


    // FromAddress - Address and name data for a call endpoint.
    FromAddress Address `json:"fromAddress"`


    // Messages - The messages sent on this communication channel.
    Messages []Messagedetails `json:"messages"`


    // JourneyContext - A subset of the Journey System's data relevant to a part of a conversation (for external linkage and internal usage/context).
    JourneyContext Journeycontext `json:"journeyContext"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`


    // AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
    AgentAssistantId string `json:"agentAssistantId"`


    // ByoSmsIntegrationId - The internal id representing the customer supplied sms integration message.
    ByoSmsIntegrationId string `json:"byoSmsIntegrationId"`


    // QueueMediaSettings - Represents the queue settings for this media type.
    QueueMediaSettings Conversationqueuemediasettings `json:"queueMediaSettings"`


    // EngagementSource
    EngagementSource string `json:"engagementSource"`

}

// String returns a JSON representation of the model
func (o *Message) String() string {
    
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Messages = []Messagedetails{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Message) MarshalJSON() ([]byte, error) {
    type Alias Message

    if MessageMarshalled {
        return []byte("{}"), nil
    }
    MessageMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        InitialState string `json:"initialState"`
        
        Id string `json:"id"`
        
        Held bool `json:"held"`
        
        Segments []Segment `json:"segments"`
        
        Direction string `json:"direction"`
        
        RecordingId string `json:"recordingId"`
        
        ErrorInfo Errorbody `json:"errorInfo"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        Provider string `json:"provider"`
        
        Authenticated bool `json:"authenticated"`
        
        VarType string `json:"type"`
        
        RecipientCountry string `json:"recipientCountry"`
        
        RecipientType string `json:"recipientType"`
        
        ScriptId string `json:"scriptId"`
        
        PeerId string `json:"peerId"`
        
        ToAddress Address `json:"toAddress"`
        
        FromAddress Address `json:"fromAddress"`
        
        Messages []Messagedetails `json:"messages"`
        
        JourneyContext Journeycontext `json:"journeyContext"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        
        AgentAssistantId string `json:"agentAssistantId"`
        
        ByoSmsIntegrationId string `json:"byoSmsIntegrationId"`
        
        QueueMediaSettings Conversationqueuemediasettings `json:"queueMediaSettings"`
        
        EngagementSource string `json:"engagementSource"`
        *Alias
    }{

        


        


        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Messages: []Messagedetails{{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

