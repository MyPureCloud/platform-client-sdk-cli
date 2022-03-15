package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ParticipantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ParticipantDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Participant
type Participant struct { 
    // Id - A globally unique identifier for this conversation.
    Id string `json:"id"`


    // StartTime - The timestamp when this participant joined the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The timestamp when this participant disconnected from the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // ConnectedTime - The timestamp when this participant was connected to the conversation in the provider clock. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConnectedTime time.Time `json:"connectedTime"`


    // Name - A human readable name identifying the participant.
    Name string `json:"name"`


    // UserUri - If this participant represents a user, then this will be an URI that can be used to fetch the user.
    UserUri string `json:"userUri"`


    // UserId - If this participant represents a user, then this will be the globally unique identifier for the user.
    UserId string `json:"userId"`


    // ExternalContactId - If this participant represents an external contact, then this will be the globally unique identifier for the external contact.
    ExternalContactId string `json:"externalContactId"`


    // ExternalOrganizationId - If this participant represents an external org, then this will be the globally unique identifier for the external org.
    ExternalOrganizationId string `json:"externalOrganizationId"`


    // QueueId - If present, the queue id that the communication channel came in on.
    QueueId string `json:"queueId"`


    // GroupId - If present, group of users the participant represents.
    GroupId string `json:"groupId"`


    // TeamId - The team id that this participant is a member of when added to the conversation.
    TeamId string `json:"teamId"`


    // QueueName - If present, the queue name that the communication channel came in on.
    QueueName string `json:"queueName"`


    // Purpose - A well known string that specifies the purpose of this participant.
    Purpose string `json:"purpose"`


    // ParticipantType - A well known string that specifies the type of this participant.
    ParticipantType string `json:"participantType"`


    // ConsultParticipantId - If this participant is part of a consult transfer, then this will be the participant id of the participant being transferred.
    ConsultParticipantId string `json:"consultParticipantId"`


    // Address - The address for the this participant. For a phone call this will be the ANI.
    Address string `json:"address"`


    // Ani - The address for the this participant. For a phone call this will be the ANI.
    Ani string `json:"ani"`


    // AniName - The ani-based name for this participant.
    AniName string `json:"aniName"`


    // Dnis - The address for the this participant. For a phone call this will be the ANI.
    Dnis string `json:"dnis"`


    // Locale - An ISO 639 language code specifying the locale for this participant
    Locale string `json:"locale"`


    // WrapupRequired - True iff this participant is required to enter wrapup for this conversation.
    WrapupRequired bool `json:"wrapupRequired"`


    // WrapupPrompt - This field controls how the UI prompts the agent for a wrapup.
    WrapupPrompt string `json:"wrapupPrompt"`


    // WrapupTimeoutMs - Specifies how long a timed ACW session will last.
    WrapupTimeoutMs int `json:"wrapupTimeoutMs"`


    // WrapupSkipped - The UI sets this field when the agent chooses to skip entering a wrapup for this participant.
    WrapupSkipped bool `json:"wrapupSkipped"`


    // Wrapup - Call wrap up or disposition data.
    Wrapup Wrapup `json:"wrapup"`


    // ConversationRoutingData - Information on how a communication should be routed to an agent.
    ConversationRoutingData Conversationroutingdata `json:"conversationRoutingData"`


    // AlertingTimeoutMs - Specifies how long the agent has to answer an interaction before being marked as not responding.
    AlertingTimeoutMs int `json:"alertingTimeoutMs"`


    // MonitoredParticipantId - If this participant is a monitor, then this will be the id of the participant that is being monitored.
    MonitoredParticipantId string `json:"monitoredParticipantId"`


    // CoachedParticipantId - If this participant is a coach, then this will be the id of the participant that is being coached.
    CoachedParticipantId string `json:"coachedParticipantId"`


    // Attributes - Additional participant attributes
    Attributes map[string]string `json:"attributes"`


    // Calls
    Calls []Call `json:"calls"`


    // Callbacks
    Callbacks []Callback `json:"callbacks"`


    // Chats
    Chats []Conversationchat `json:"chats"`


    // Cobrowsesessions
    Cobrowsesessions []Cobrowsesession `json:"cobrowsesessions"`


    // Emails
    Emails []Email `json:"emails"`


    // Messages
    Messages []Message `json:"messages"`


    // Screenshares
    Screenshares []Screenshare `json:"screenshares"`


    // SocialExpressions
    SocialExpressions []Socialexpression `json:"socialExpressions"`


    // Videos
    Videos []Video `json:"videos"`


    // Evaluations
    Evaluations []Evaluation `json:"evaluations"`


    // ScreenRecordingState - The current screen recording state for this participant.
    ScreenRecordingState string `json:"screenRecordingState"`


    // FlaggedReason - The reason specifying why participant flagged the conversation.
    FlaggedReason string `json:"flaggedReason"`


    // StartAcwTime - The timestamp when this participant started after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartAcwTime time.Time `json:"startAcwTime"`


    // EndAcwTime - The timestamp when this participant ended after-call work. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndAcwTime time.Time `json:"endAcwTime"`


    // BargedParticipantId - If this participant barged in a participant's call, then this will be the id of the targeted participant.
    BargedParticipantId string `json:"bargedParticipantId"`

}

// String returns a JSON representation of the model
func (o *Participant) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Attributes = map[string]string{"": ""} 
    
    
    
     o.Calls = []Call{{}} 
    
    
    
     o.Callbacks = []Callback{{}} 
    
    
    
     o.Chats = []Conversationchat{{}} 
    
    
    
     o.Cobrowsesessions = []Cobrowsesession{{}} 
    
    
    
     o.Emails = []Email{{}} 
    
    
    
     o.Messages = []Message{{}} 
    
    
    
     o.Screenshares = []Screenshare{{}} 
    
    
    
     o.SocialExpressions = []Socialexpression{{}} 
    
    
    
     o.Videos = []Video{{}} 
    
    
    
     o.Evaluations = []Evaluation{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Participant) MarshalJSON() ([]byte, error) {
    type Alias Participant

    if ParticipantMarshalled {
        return []byte("{}"), nil
    }
    ParticipantMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        ConnectedTime time.Time `json:"connectedTime"`
        
        Name string `json:"name"`
        
        UserUri string `json:"userUri"`
        
        UserId string `json:"userId"`
        
        ExternalContactId string `json:"externalContactId"`
        
        ExternalOrganizationId string `json:"externalOrganizationId"`
        
        QueueId string `json:"queueId"`
        
        GroupId string `json:"groupId"`
        
        TeamId string `json:"teamId"`
        
        QueueName string `json:"queueName"`
        
        Purpose string `json:"purpose"`
        
        ParticipantType string `json:"participantType"`
        
        ConsultParticipantId string `json:"consultParticipantId"`
        
        Address string `json:"address"`
        
        Ani string `json:"ani"`
        
        AniName string `json:"aniName"`
        
        Dnis string `json:"dnis"`
        
        Locale string `json:"locale"`
        
        WrapupRequired bool `json:"wrapupRequired"`
        
        WrapupPrompt string `json:"wrapupPrompt"`
        
        WrapupTimeoutMs int `json:"wrapupTimeoutMs"`
        
        WrapupSkipped bool `json:"wrapupSkipped"`
        
        Wrapup Wrapup `json:"wrapup"`
        
        ConversationRoutingData Conversationroutingdata `json:"conversationRoutingData"`
        
        AlertingTimeoutMs int `json:"alertingTimeoutMs"`
        
        MonitoredParticipantId string `json:"monitoredParticipantId"`
        
        CoachedParticipantId string `json:"coachedParticipantId"`
        
        Attributes map[string]string `json:"attributes"`
        
        Calls []Call `json:"calls"`
        
        Callbacks []Callback `json:"callbacks"`
        
        Chats []Conversationchat `json:"chats"`
        
        Cobrowsesessions []Cobrowsesession `json:"cobrowsesessions"`
        
        Emails []Email `json:"emails"`
        
        Messages []Message `json:"messages"`
        
        Screenshares []Screenshare `json:"screenshares"`
        
        SocialExpressions []Socialexpression `json:"socialExpressions"`
        
        Videos []Video `json:"videos"`
        
        Evaluations []Evaluation `json:"evaluations"`
        
        ScreenRecordingState string `json:"screenRecordingState"`
        
        FlaggedReason string `json:"flaggedReason"`
        
        StartAcwTime time.Time `json:"startAcwTime"`
        
        EndAcwTime time.Time `json:"endAcwTime"`
        
        BargedParticipantId string `json:"bargedParticipantId"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Attributes: map[string]string{"": ""},
        

        

        
        Calls: []Call{{}},
        

        

        
        Callbacks: []Callback{{}},
        

        

        
        Chats: []Conversationchat{{}},
        

        

        
        Cobrowsesessions: []Cobrowsesession{{}},
        

        

        
        Emails: []Email{{}},
        

        

        
        Messages: []Message{{}},
        

        

        
        Screenshares: []Screenshare{{}},
        

        

        
        SocialExpressions: []Socialexpression{{}},
        

        

        
        Videos: []Video{{}},
        

        

        
        Evaluations: []Evaluation{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

