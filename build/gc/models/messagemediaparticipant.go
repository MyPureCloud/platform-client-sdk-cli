package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessagemediaparticipantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessagemediaparticipantDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Messagemediaparticipant
type Messagemediaparticipant struct { 
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


    // ToAddress - Address for the participant on receiving side of the message conversation. If the address is a phone number, E.164 format is recommended.
    ToAddress Address `json:"toAddress"`


    // FromAddress - Address for the participant on the sending side of the message conversation. If the address is a phone number, E.164 format is recommended.
    FromAddress Address `json:"fromAddress"`


    // Messages - Message instance details on the communication.
    Messages []Messagedetails `json:"messages"`


    // VarType - Indicates the type of message platform from which the message originated.
    VarType string `json:"type"`


    // RecipientCountry - Indicates the country where the recipient is associated in ISO 3166-1 alpha-2 format.
    RecipientCountry string `json:"recipientCountry"`


    // RecipientType - The type of the recipient. Eg: Provisioned phoneNumber is the recipient for sms message type.
    RecipientType string `json:"recipientType"`


    // Authenticated - If true, the participant member is authenticated.
    Authenticated bool `json:"authenticated"`

}

// String returns a JSON representation of the model
func (o *Messagemediaparticipant) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.MediaRoles = []string{""} 
    
    
    
     o.Attributes = map[string]string{"": ""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Messages = []Messagedetails{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messagemediaparticipant) MarshalJSON() ([]byte, error) {
    type Alias Messagemediaparticipant

    if MessagemediaparticipantMarshalled {
        return []byte("{}"), nil
    }
    MessagemediaparticipantMarshalled = true

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
        
        ToAddress Address `json:"toAddress"`
        
        FromAddress Address `json:"fromAddress"`
        
        Messages []Messagedetails `json:"messages"`
        
        VarType string `json:"type"`
        
        RecipientCountry string `json:"recipientCountry"`
        
        RecipientType string `json:"recipientType"`
        
        Authenticated bool `json:"authenticated"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        MediaRoles: []string{""},
        


        


        


        


        
        Attributes: map[string]string{"": ""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Messages: []Messagedetails{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

