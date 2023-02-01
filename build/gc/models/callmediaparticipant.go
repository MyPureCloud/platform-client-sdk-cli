package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallmediaparticipantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallmediaparticipantDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Callmediaparticipant
type Callmediaparticipant struct { 
    // Id - The unique participant ID.
    Id string `json:"id"`


    // Name - The display friendly name of the participant.
    Name string `json:"name"`


    // Address - The participant address.
    Address string `json:"address"`


    // StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // ConnectedTime - The time when this participant went connected for this media (eg: video connected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // StartHoldTime - The time when this participant's hold started. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartHoldTime time.Time `json:"startHoldTime"`


    // Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
    Purpose string `json:"purpose"`


    // State - The participant's state.  Values can be: 'alerting', 'connected', 'disconnected', 'dialing', 'contacting
    State string `json:"state"`


    // Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
    Direction string `json:"direction"`


    // DisconnectType - The reason the participant was disconnected from the conversation.
    DisconnectType string `json:"disconnectType"`


    // Held - Value is true when the participant is on hold.
    Held bool `json:"held"`


    // WrapupRequired - Value is true when the participant requires wrap-up.
    WrapupRequired bool `json:"wrapupRequired"`


    // WrapupPrompt - The wrap-up prompt indicating the type of wrap-up to be performed.
    WrapupPrompt string `json:"wrapupPrompt"`


    // MediaRoles - List of roles this participant's media has had on the conversation, ie monitor, coach, etc
    MediaRoles []string `json:"mediaRoles"`


    // User - The PureCloud user for this participant.
    User Domainentityref `json:"user"`


    // Queue - The PureCloud queue for this participant.
    Queue Domainentityref `json:"queue"`


    // Team - The PureCloud team for this participant.
    Team Domainentityref `json:"team"`


    // Attributes - A list of ad-hoc attributes for the participant.
    Attributes map[string]string `json:"attributes"`


    // ErrorInfo - If the conversation ends in error, contains additional error details.
    ErrorInfo Errorinfo `json:"errorInfo"`


    // Script - The Engage script that should be used by this participant.
    Script Domainentityref `json:"script"`


    // WrapupTimeoutMs - The amount of time the participant has to complete wrap-up.
    WrapupTimeoutMs int `json:"wrapupTimeoutMs"`


    // WrapupSkipped - Value is true when the participant has skipped wrap-up.
    WrapupSkipped bool `json:"wrapupSkipped"`


    // AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
    AlertingTimeoutMs int `json:"alertingTimeoutMs"`


    // Provider - The source provider for the communication.
    Provider string `json:"provider"`


    // ExternalContact - If this participant represents an external contact, then this will be the reference for the external contact.
    ExternalContact Domainentityref `json:"externalContact"`


    // ExternalOrganization - If this participant represents an external org, then this will be the reference for the external org.
    ExternalOrganization Domainentityref `json:"externalOrganization"`


    // Wrapup - Wrapup for this participant, if it has been applied.
    Wrapup Wrapup `json:"wrapup"`


    // Peer - The peer communication corresponding to a matching leg for this communication.
    Peer string `json:"peer"`


    // FlaggedReason - The reason specifying why participant flagged the conversation.
    FlaggedReason string `json:"flaggedReason"`


    // JourneyContext - Journey System data/context that is applicable to this communication.  When used for historical purposes, the context should be immutable.  When null, there is no applicable Journey System context.
    JourneyContext Journeycontext `json:"journeyContext"`


    // ConversationRoutingData - Information on how a communication should be routed to an agent.
    ConversationRoutingData Conversationroutingdata `json:"conversationRoutingData"`


    // StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAcwTime time.Time `json:"startAcwTime"`


    // EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndAcwTime time.Time `json:"endAcwTime"`


    // Muted - Value is true when the call is muted.
    Muted bool `json:"muted"`


    // Confined - Value is true when the call is confined.
    Confined bool `json:"confined"`


    // Recording - Value is true when the call is being recorded.
    Recording bool `json:"recording"`


    // RecordingState - The state of the call recording.
    RecordingState string `json:"recordingState"`


    // Group - The group involved in the group ring call.
    Group Domainentityref `json:"group"`


    // Ani - The call ANI.
    Ani string `json:"ani"`


    // Dnis - The call DNIS.
    Dnis string `json:"dnis"`


    // DocumentId - The ID of the Content Management document if the call is a fax.
    DocumentId string `json:"documentId"`


    // FaxStatus - Extra fax information if the call is a fax.
    FaxStatus Faxstatus `json:"faxStatus"`


    // MonitoredParticipantId - The ID of the participant being monitored when performing a call monitor.
    MonitoredParticipantId string `json:"monitoredParticipantId"`


    // CoachedParticipantId - The ID of the participant being coached when performing a call coach.
    CoachedParticipantId string `json:"coachedParticipantId"`


    // BargedParticipantId - If this participant barged in a participant's call, then this will be the id of the targeted participant.
    BargedParticipantId string `json:"bargedParticipantId"`


    // ConsultParticipantId - The ID of the consult transfer target participant when performing a consult transfer.
    ConsultParticipantId string `json:"consultParticipantId"`


    // UuiData - User-to-User information which maps to a SIP header field defined in RFC7433. UUI data is used in the Public Switched Telephone Network (PSTN) for use cases described in RFC6567.
    UuiData string `json:"uuiData"`


    // BargedTime - The timestamp when this participant was connected to the barge conference in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    BargedTime time.Time `json:"bargedTime"`

}

// String returns a JSON representation of the model
func (o *Callmediaparticipant) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MediaRoles = []string{""} 
    
    
    
     o.Attributes = map[string]string{"": ""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callmediaparticipant) MarshalJSON() ([]byte, error) {
    type Alias Callmediaparticipant

    if CallmediaparticipantMarshalled {
        return []byte("{}"), nil
    }
    CallmediaparticipantMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Address string `json:"address"`
        
        StartTime time.Time `json:"startTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        EndTime time.Time `json:"endTime"`
        
        StartHoldTime time.Time `json:"startHoldTime"`
        
        Purpose string `json:"purpose"`
        
        State string `json:"state"`
        
        Direction string `json:"direction"`
        
        DisconnectType string `json:"disconnectType"`
        
        Held bool `json:"held"`
        
        WrapupRequired bool `json:"wrapupRequired"`
        
        WrapupPrompt string `json:"wrapupPrompt"`
        
        MediaRoles []string `json:"mediaRoles"`
        
        User Domainentityref `json:"user"`
        
        Queue Domainentityref `json:"queue"`
        
        Team Domainentityref `json:"team"`
        
        Attributes map[string]string `json:"attributes"`
        
        ErrorInfo Errorinfo `json:"errorInfo"`
        
        Script Domainentityref `json:"script"`
        
        WrapupTimeoutMs int `json:"wrapupTimeoutMs"`
        
        WrapupSkipped bool `json:"wrapupSkipped"`
        
        AlertingTimeoutMs int `json:"alertingTimeoutMs"`
        
        Provider string `json:"provider"`
        
        ExternalContact Domainentityref `json:"externalContact"`
        
        ExternalOrganization Domainentityref `json:"externalOrganization"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        Peer string `json:"peer"`
        
        FlaggedReason string `json:"flaggedReason"`
        
        JourneyContext Journeycontext `json:"journeyContext"`
        
        ConversationRoutingData Conversationroutingdata `json:"conversationRoutingData"`
        
        StartAcwTime time.Time `json:"startAcwTime"`
        
        EndAcwTime time.Time `json:"endAcwTime"`
        
        Muted bool `json:"muted"`
        
        Confined bool `json:"confined"`
        
        Recording bool `json:"recording"`
        
        RecordingState string `json:"recordingState"`
        
        Group Domainentityref `json:"group"`
        
        Ani string `json:"ani"`
        
        Dnis string `json:"dnis"`
        
        DocumentId string `json:"documentId"`
        
        FaxStatus Faxstatus `json:"faxStatus"`
        
        MonitoredParticipantId string `json:"monitoredParticipantId"`
        
        CoachedParticipantId string `json:"coachedParticipantId"`
        
        BargedParticipantId string `json:"bargedParticipantId"`
        
        ConsultParticipantId string `json:"consultParticipantId"`
        
        UuiData string `json:"uuiData"`
        
        BargedTime time.Time `json:"bargedTime"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        MediaRoles: []string{""},
        


        


        


        


        
        Attributes: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

