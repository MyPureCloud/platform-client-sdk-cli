package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticssessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticssessionDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticssession
type Analyticssession struct { 
    // ActiveSkillIds - ID(s) of Skill(s) that are active on the conversation
    ActiveSkillIds []string `json:"activeSkillIds"`


    // AcwSkipped - Marker for an agent that skipped after call work
    AcwSkipped bool `json:"acwSkipped"`


    // AddressFrom - The address that initiated an action
    AddressFrom string `json:"addressFrom"`


    // AddressOther - The email address for the participant on the other side of the email conversation
    AddressOther string `json:"addressOther"`


    // AddressSelf - The email address for the participant on this side of the email conversation
    AddressSelf string `json:"addressSelf"`


    // AddressTo - The address receiving an action
    AddressTo string `json:"addressTo"`


    // AgentAssistantId - Unique identifier of the active virtual agent assistant
    AgentAssistantId string `json:"agentAssistantId"`


    // AgentBullseyeRing - Bullseye ring of the targeted agent
    AgentBullseyeRing int `json:"agentBullseyeRing"`


    // AgentOwned - Flag indicating an agent-owned callback
    AgentOwned bool `json:"agentOwned"`


    // Ani - Automatic Number Identification (caller's number)
    Ani string `json:"ani"`


    // AssignerId - ID of the user that manually assigned a conversation
    AssignerId string `json:"assignerId"`


    // Authenticated - Flag that indicates that the identity of the customer has been asserted as verified by the provider.
    Authenticated bool `json:"authenticated"`


    // CallbackNumbers - Callback phone number(s)
    CallbackNumbers []string `json:"callbackNumbers"`


    // CallbackScheduledTime - Scheduled callback date/time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CallbackScheduledTime time.Time `json:"callbackScheduledTime"`


    // CallbackUserName - The name of the user requesting a call back
    CallbackUserName string `json:"callbackUserName"`


    // CoachedParticipantId - The participantId being coached (if someone (e.g. an agent) is being coached, this would correspond to one of the other participantIds present in the conversation)
    CoachedParticipantId string `json:"coachedParticipantId"`


    // CobrowseRole - Describes side of the cobrowse (sharer or viewer)
    CobrowseRole string `json:"cobrowseRole"`


    // CobrowseRoomId - A unique identifier for a PureCloud cobrowse room
    CobrowseRoomId string `json:"cobrowseRoomId"`


    // DeliveryStatus - The email or SMS delivery status
    DeliveryStatus string `json:"deliveryStatus"`


    // DeliveryStatusChangeDate - Date and time of the most recent delivery status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DeliveryStatusChangeDate time.Time `json:"deliveryStatusChangeDate"`


    // DestinationAddresses - Destination address(es) of transfers or consults
    DestinationAddresses []string `json:"destinationAddresses"`


    // Direction - The direction of the communication
    Direction string `json:"direction"`


    // DispositionAnalyzer - (Dialer) Analyzer (for example speech.person)
    DispositionAnalyzer string `json:"dispositionAnalyzer"`


    // DispositionName - (Dialer) Result of the analysis (for example disposition.classification.callable.machine)
    DispositionName string `json:"dispositionName"`


    // Dnis - Dialed number identification service (number dialed by the calling party)
    Dnis string `json:"dnis"`


    // EdgeId - Unique identifier of the edge device
    EdgeId string `json:"edgeId"`


    // EligibleAgentCounts - Number of eligible agents for each predictive routing attempt
    EligibleAgentCounts []int `json:"eligibleAgentCounts"`


    // ExtendedDeliveryStatus - Extended email delivery status
    ExtendedDeliveryStatus string `json:"extendedDeliveryStatus"`


    // FlowInType - Type of flow in that occurred when entering ACD.
    FlowInType string `json:"flowInType"`


    // FlowOutType - Type of flow out that occurred when emitting tFlowOut.
    FlowOutType string `json:"flowOutType"`


    // JourneyActionId - Identifier of the journey action.
    JourneyActionId string `json:"journeyActionId"`


    // JourneyActionMapId - Identifier of the journey action map that triggered the action.
    JourneyActionMapId string `json:"journeyActionMapId"`


    // JourneyActionMapVersion - Version of the journey action map that triggered the action.
    JourneyActionMapVersion int `json:"journeyActionMapVersion"`


    // JourneyCustomerId - Primary identifier of the journey customer in the source where the activities originate from.
    JourneyCustomerId string `json:"journeyCustomerId"`


    // JourneyCustomerIdType - Type of primary identifier of the journey customer (e.g. cookie).
    JourneyCustomerIdType string `json:"journeyCustomerIdType"`


    // JourneyCustomerSessionId - Unique identifier of the journey session.
    JourneyCustomerSessionId string `json:"journeyCustomerSessionId"`


    // JourneyCustomerSessionIdType - Type or category of journey sessions (e.g. web, ticket, delivery, atm).
    JourneyCustomerSessionIdType string `json:"journeyCustomerSessionIdType"`


    // MediaBridgeId - Media bridge ID for the conference session consistent across all participants
    MediaBridgeId string `json:"mediaBridgeId"`


    // MediaCount - Count of any media (images, files, etc) included in this session
    MediaCount int `json:"mediaCount"`


    // MediaType - The session media type
    MediaType string `json:"mediaType"`


    // MessageType - Message type for messaging services. E.g.: sms, facebook, twitter, line
    MessageType string `json:"messageType"`


    // MonitoredParticipantId - The participantId being monitored (if someone (e.g. an agent) is being monitored, this would correspond to one of the other participantIds present in the conversation)
    MonitoredParticipantId string `json:"monitoredParticipantId"`


    // OutboundCampaignId - (Dialer) Unique identifier of the outbound campaign
    OutboundCampaignId string `json:"outboundCampaignId"`


    // OutboundContactId - (Dialer) Unique identifier of the contact
    OutboundContactId string `json:"outboundContactId"`


    // OutboundContactListId - (Dialer) Unique identifier of the contact list that this contact belongs to
    OutboundContactListId string `json:"outboundContactListId"`


    // PeerId - This identifies pairs of related sessions on a conversation. E.g. an external session’s peerId will be the session that the call originally connected to, e.g. if an IVR was dialed, the IVR session, which will also have the external session’s ID as its peer. After that point, any transfers of that session to other internal components (acd, agent, etc.) will all spawn new sessions whose peerIds point back to that original external session.
    PeerId string `json:"peerId"`


    // ProtocolCallId - The original voice protocol call ID, e.g. a SIP call ID
    ProtocolCallId string `json:"protocolCallId"`


    // Provider - The source provider for the communication.
    Provider string `json:"provider"`


    // Recording - Flag determining if an audio recording was started or not
    Recording bool `json:"recording"`


    // Remote - Name, phone number, or email address of the remote party.
    Remote string `json:"remote"`


    // RemoteNameDisplayable - Unique identifier for the remote party
    RemoteNameDisplayable string `json:"remoteNameDisplayable"`


    // RemovedSkillIds - ID(s) of Skill(s) that have been removed by bullseye routing
    RemovedSkillIds []string `json:"removedSkillIds"`


    // RequestedRoutings - Routing type(s) for requested/attempted routing methods.
    RequestedRoutings []string `json:"requestedRoutings"`


    // RoomId - Unique identifier for the room
    RoomId string `json:"roomId"`


    // RoutingRing - Routing ring for bullseye or preferred agent routing
    RoutingRing int `json:"routingRing"`


    // ScreenShareAddressSelf - Direct ScreenShare address
    ScreenShareAddressSelf string `json:"screenShareAddressSelf"`


    // ScreenShareRoomId - A unique identifier for a PureCloud ScreenShare room
    ScreenShareRoomId string `json:"screenShareRoomId"`


    // ScriptId - A unique identifier for a script
    ScriptId string `json:"scriptId"`


    // SelectedAgentId - Selected agent ID
    SelectedAgentId string `json:"selectedAgentId"`


    // SelectedAgentRank - Selected agent GPR rank
    SelectedAgentRank int `json:"selectedAgentRank"`


    // SessionDnis - Dialed number for the current session; this can be different from dnis, e.g. if the call was transferred
    SessionDnis string `json:"sessionDnis"`


    // SessionId - The unique identifier of this session
    SessionId string `json:"sessionId"`


    // SharingScreen - Flag determining if screenShare is started or not (true/false)
    SharingScreen bool `json:"sharingScreen"`


    // SkipEnabled - (Dialer) Whether the agent can skip the dialer contact
    SkipEnabled bool `json:"skipEnabled"`


    // TimeoutSeconds - The number of seconds before PureCloud begins the call for a call back (0 disables automatic calling)
    TimeoutSeconds int `json:"timeoutSeconds"`


    // UsedRouting - Complete routing method
    UsedRouting string `json:"usedRouting"`


    // VideoAddressSelf - Direct Video address
    VideoAddressSelf string `json:"videoAddressSelf"`


    // VideoRoomId - A unique identifier for a PureCloud video room
    VideoRoomId string `json:"videoRoomId"`


    // WaitingInteractionCounts - Number of waiting interactions for each predictive routing attempt
    WaitingInteractionCounts []int `json:"waitingInteractionCounts"`


    // ProposedAgents - Proposed agents
    ProposedAgents []Analyticsproposedagent `json:"proposedAgents"`


    // MediaEndpointStats - MediaEndpointStats associated with this session
    MediaEndpointStats []Analyticsmediaendpointstat `json:"mediaEndpointStats"`


    // Flow - IVR flow execution associated with this session
    Flow Analyticsflow `json:"flow"`


    // Metrics - List of metrics for this session
    Metrics []Analyticssessionmetric `json:"metrics"`


    // Segments - List of segments for this session
    Segments []Analyticsconversationsegment `json:"segments"`

}

// String returns a JSON representation of the model
func (o *Analyticssession) String() string {
     o.ActiveSkillIds = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
     o.CallbackNumbers = []string{""} 
    
    
    
    
    
    
    
     o.DestinationAddresses = []string{""} 
    
    
    
    
    
     o.EligibleAgentCounts = []int{0} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RemovedSkillIds = []string{""} 
     o.RequestedRoutings = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.WaitingInteractionCounts = []int{0} 
     o.ProposedAgents = []Analyticsproposedagent{{}} 
     o.MediaEndpointStats = []Analyticsmediaendpointstat{{}} 
    
     o.Metrics = []Analyticssessionmetric{{}} 
     o.Segments = []Analyticsconversationsegment{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticssession) MarshalJSON() ([]byte, error) {
    type Alias Analyticssession

    if AnalyticssessionMarshalled {
        return []byte("{}"), nil
    }
    AnalyticssessionMarshalled = true

    return json.Marshal(&struct {
        
        ActiveSkillIds []string `json:"activeSkillIds"`
        
        AcwSkipped bool `json:"acwSkipped"`
        
        AddressFrom string `json:"addressFrom"`
        
        AddressOther string `json:"addressOther"`
        
        AddressSelf string `json:"addressSelf"`
        
        AddressTo string `json:"addressTo"`
        
        AgentAssistantId string `json:"agentAssistantId"`
        
        AgentBullseyeRing int `json:"agentBullseyeRing"`
        
        AgentOwned bool `json:"agentOwned"`
        
        Ani string `json:"ani"`
        
        AssignerId string `json:"assignerId"`
        
        Authenticated bool `json:"authenticated"`
        
        CallbackNumbers []string `json:"callbackNumbers"`
        
        CallbackScheduledTime time.Time `json:"callbackScheduledTime"`
        
        CallbackUserName string `json:"callbackUserName"`
        
        CoachedParticipantId string `json:"coachedParticipantId"`
        
        CobrowseRole string `json:"cobrowseRole"`
        
        CobrowseRoomId string `json:"cobrowseRoomId"`
        
        DeliveryStatus string `json:"deliveryStatus"`
        
        DeliveryStatusChangeDate time.Time `json:"deliveryStatusChangeDate"`
        
        DestinationAddresses []string `json:"destinationAddresses"`
        
        Direction string `json:"direction"`
        
        DispositionAnalyzer string `json:"dispositionAnalyzer"`
        
        DispositionName string `json:"dispositionName"`
        
        Dnis string `json:"dnis"`
        
        EdgeId string `json:"edgeId"`
        
        EligibleAgentCounts []int `json:"eligibleAgentCounts"`
        
        ExtendedDeliveryStatus string `json:"extendedDeliveryStatus"`
        
        FlowInType string `json:"flowInType"`
        
        FlowOutType string `json:"flowOutType"`
        
        JourneyActionId string `json:"journeyActionId"`
        
        JourneyActionMapId string `json:"journeyActionMapId"`
        
        JourneyActionMapVersion int `json:"journeyActionMapVersion"`
        
        JourneyCustomerId string `json:"journeyCustomerId"`
        
        JourneyCustomerIdType string `json:"journeyCustomerIdType"`
        
        JourneyCustomerSessionId string `json:"journeyCustomerSessionId"`
        
        JourneyCustomerSessionIdType string `json:"journeyCustomerSessionIdType"`
        
        MediaBridgeId string `json:"mediaBridgeId"`
        
        MediaCount int `json:"mediaCount"`
        
        MediaType string `json:"mediaType"`
        
        MessageType string `json:"messageType"`
        
        MonitoredParticipantId string `json:"monitoredParticipantId"`
        
        OutboundCampaignId string `json:"outboundCampaignId"`
        
        OutboundContactId string `json:"outboundContactId"`
        
        OutboundContactListId string `json:"outboundContactListId"`
        
        PeerId string `json:"peerId"`
        
        ProtocolCallId string `json:"protocolCallId"`
        
        Provider string `json:"provider"`
        
        Recording bool `json:"recording"`
        
        Remote string `json:"remote"`
        
        RemoteNameDisplayable string `json:"remoteNameDisplayable"`
        
        RemovedSkillIds []string `json:"removedSkillIds"`
        
        RequestedRoutings []string `json:"requestedRoutings"`
        
        RoomId string `json:"roomId"`
        
        RoutingRing int `json:"routingRing"`
        
        ScreenShareAddressSelf string `json:"screenShareAddressSelf"`
        
        ScreenShareRoomId string `json:"screenShareRoomId"`
        
        ScriptId string `json:"scriptId"`
        
        SelectedAgentId string `json:"selectedAgentId"`
        
        SelectedAgentRank int `json:"selectedAgentRank"`
        
        SessionDnis string `json:"sessionDnis"`
        
        SessionId string `json:"sessionId"`
        
        SharingScreen bool `json:"sharingScreen"`
        
        SkipEnabled bool `json:"skipEnabled"`
        
        TimeoutSeconds int `json:"timeoutSeconds"`
        
        UsedRouting string `json:"usedRouting"`
        
        VideoAddressSelf string `json:"videoAddressSelf"`
        
        VideoRoomId string `json:"videoRoomId"`
        
        WaitingInteractionCounts []int `json:"waitingInteractionCounts"`
        
        ProposedAgents []Analyticsproposedagent `json:"proposedAgents"`
        
        MediaEndpointStats []Analyticsmediaendpointstat `json:"mediaEndpointStats"`
        
        Flow Analyticsflow `json:"flow"`
        
        Metrics []Analyticssessionmetric `json:"metrics"`
        
        Segments []Analyticsconversationsegment `json:"segments"`
        *Alias
    }{

        
        ActiveSkillIds: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        
        CallbackNumbers: []string{""},
        


        


        


        


        


        


        


        


        
        DestinationAddresses: []string{""},
        


        


        


        


        


        


        
        EligibleAgentCounts: []int{0},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        RemovedSkillIds: []string{""},
        


        
        RequestedRoutings: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        WaitingInteractionCounts: []int{0},
        


        
        ProposedAgents: []Analyticsproposedagent{{}},
        


        
        MediaEndpointStats: []Analyticsmediaendpointstat{{}},
        


        


        
        Metrics: []Analyticssessionmetric{{}},
        


        
        Segments: []Analyticsconversationsegment{{}},
        

        Alias: (*Alias)(u),
    })
}

