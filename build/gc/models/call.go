package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Call
type Call struct { 
    // State - The connection state of this communication.
    State string `json:"state"`


    // InitialState - The initial connection state of this communication.
    InitialState string `json:"initialState"`


    // Id - A globally unique identifier for this communication.
    Id string `json:"id"`


    // Direction - The direction of the call
    Direction string `json:"direction"`


    // Recording - True if this call is being recorded.
    Recording bool `json:"recording"`


    // RecordingState - State of recording on this call.
    RecordingState string `json:"recordingState"`


    // RecordersState - Contains the states of different recorders.
    RecordersState Recordersstate `json:"recordersState"`


    // Muted - True if this call is muted so that remote participants can't hear any audio from this end.
    Muted bool `json:"muted"`


    // Confined - True if this call is held and the person on this side hears hold music.
    Confined bool `json:"confined"`


    // Held - True if this call is held and the person on this side hears silence.
    Held bool `json:"held"`


    // SecurePause - True when the recording of this call is in secure pause status.
    SecurePause bool `json:"securePause"`


    // RecordingId - A globally unique identifier for the recording associated with this call.
    RecordingId string `json:"recordingId"`


    // Segments - The time line of the participant's call, divided into activity segments.
    Segments []Segment `json:"segments"`


    // ErrorInfo
    ErrorInfo Errorinfo `json:"errorInfo"`


    // DisconnectType - System defined string indicating what caused the communication to disconnect. Will be null until the communication disconnects.
    DisconnectType string `json:"disconnectType"`


    // StartHoldTime - The timestamp the call was placed on hold in the cloud clock if the call is currently on hold. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // DocumentId - If call is an outbound fax of a document from content management, then this is the id in content management.
    DocumentId string `json:"documentId"`


    // StartAlertingTime - The timestamp the communication has when it is first put into an alerting state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAlertingTime time.Time `json:"startAlertingTime"`


    // ConnectedTime - The timestamp when this communication was connected in the cloud clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // DisconnectedTime - The timestamp when this communication disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DisconnectedTime time.Time `json:"disconnectedTime"`


    // DisconnectReasons - List of reasons that this call was disconnected. This will be set once the call disconnects.
    DisconnectReasons []Disconnectreason `json:"disconnectReasons"`


    // FaxStatus - Extra information on fax transmission.
    FaxStatus Faxstatus `json:"faxStatus"`


    // Provider - The source provider for the call.
    Provider string `json:"provider"`


    // ScriptId - The UUID of the script to use.
    ScriptId string `json:"scriptId"`


    // PeerId - The id of the peer communication corresponding to a matching leg for this communication.
    PeerId string `json:"peerId"`


    // UuiData - User to User Information (UUI) data managed by SIP session application.
    UuiData string `json:"uuiData"`


    // Self - Address and name data for a call endpoint.
    Self Address `json:"self"`


    // Other - Address and name data for a call endpoint.
    Other Address `json:"other"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // AfterCallWork - After-call work for the communication.
    AfterCallWork Aftercallwork `json:"afterCallWork"`


    // AfterCallWorkRequired - Indicates if after-call work is required for a communication. Only used when the ACW Setting is Agent Requested.
    AfterCallWorkRequired bool `json:"afterCallWorkRequired"`


    // AgentAssistantId - UUID of virtual agent assistant that provide suggestions to the agent participant during the conversation.
    AgentAssistantId string `json:"agentAssistantId"`


    // TransferSource - Indicates how call reaches the agent.
    TransferSource string `json:"transferSource"`


    // QueueMediaSettings - Represents the queue settings for this media type.
    QueueMediaSettings Conversationqueuemediasettings `json:"queueMediaSettings"`


    // Disposition - Call resolution data for Dialer bulk make calls commands.
    Disposition Disposition `json:"disposition"`

}

// String returns a JSON representation of the model
func (o *Call) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Segments = []Segment{{}} 
    
    
    
    
    
    
    
     o.DisconnectReasons = []Disconnectreason{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Call) MarshalJSON() ([]byte, error) {
    type Alias Call

    if CallMarshalled {
        return []byte("{}"), nil
    }
    CallMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        InitialState string `json:"initialState"`
        
        Id string `json:"id"`
        
        Direction string `json:"direction"`
        
        Recording bool `json:"recording"`
        
        RecordingState string `json:"recordingState"`
        
        RecordersState Recordersstate `json:"recordersState"`
        
        Muted bool `json:"muted"`
        
        Confined bool `json:"confined"`
        
        Held bool `json:"held"`
        
        SecurePause bool `json:"securePause"`
        
        RecordingId string `json:"recordingId"`
        
        Segments []Segment `json:"segments"`
        
        ErrorInfo Errorinfo `json:"errorInfo"`
        
        DisconnectType string `json:"disconnectType"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        DocumentId string `json:"documentId"`
        
        StartAlertingTime time.Time `json:"startAlertingTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        DisconnectedTime time.Time `json:"disconnectedTime"`
        
        DisconnectReasons []Disconnectreason `json:"disconnectReasons"`
        
        FaxStatus Faxstatus `json:"faxStatus"`
        
        Provider string `json:"provider"`
        
        ScriptId string `json:"scriptId"`
        
        PeerId string `json:"peerId"`
        
        UuiData string `json:"uuiData"`
        
        Self Address `json:"self"`
        
        Other Address `json:"other"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        AfterCallWork Aftercallwork `json:"afterCallWork"`
        
        AfterCallWorkRequired bool `json:"afterCallWorkRequired"`
        
        AgentAssistantId string `json:"agentAssistantId"`
        
        TransferSource string `json:"transferSource"`
        
        QueueMediaSettings Conversationqueuemediasettings `json:"queueMediaSettings"`
        
        Disposition Disposition `json:"disposition"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Segments: []Segment{{}},
        


        


        


        


        


        


        


        


        
        DisconnectReasons: []Disconnectreason{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

