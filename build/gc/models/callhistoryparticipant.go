package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallhistoryparticipantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallhistoryparticipantDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Callhistoryparticipant
type Callhistoryparticipant struct { 
    // Id - The unique participant ID.
    Id string `json:"id"`


    // Name - The display friendly name of the participant.
    Name string `json:"name"`


    // Address - The participant address.
    Address string `json:"address"`


    // StartTime - The time when this participant first joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartTime time.Time `json:"startTime"`


    // EndTime - The time when this participant went disconnected for this media (eg: video disconnected time). Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EndTime time.Time `json:"endTime"`


    // Purpose - The participant's purpose.  Values can be: 'agent', 'user', 'customer', 'external', 'acd', 'ivr
    Purpose string `json:"purpose"`


    // Direction - The participant's direction.  Values can be: 'inbound' or 'outbound'
    Direction string `json:"direction"`


    // Ani - The call ANI.
    Ani string `json:"ani"`


    // Dnis - The call DNIS.
    Dnis string `json:"dnis"`


    // User - The PureCloud user for this participant.
    User User `json:"user"`


    // Queue - The PureCloud queue for this participant.
    Queue Queue `json:"queue"`


    // Group - The group involved in the group ring call.
    Group Group `json:"group"`


    // DisconnectType - The reason the participant was disconnected from the conversation.
    DisconnectType string `json:"disconnectType"`


    // ExternalContact - The PureCloud external contact
    ExternalContact Externalcontact `json:"externalContact"`


    // ExternalOrganization - The PureCloud external organization
    ExternalOrganization Externalorganization `json:"externalOrganization"`


    // DidInteract - Indicates whether the contact ever connected
    DidInteract bool `json:"didInteract"`


    // SipResponseCodes - Indicates SIP Response codes associated with the participant
    SipResponseCodes []int `json:"sipResponseCodes"`


    // FlaggedReason - The reason specifying why participant flagged the conversation.
    FlaggedReason string `json:"flaggedReason"`


    // OutboundCampaign - The outbound campaign associated with the participant
    OutboundCampaign Campaign `json:"outboundCampaign"`

}

// String returns a JSON representation of the model
func (o *Callhistoryparticipant) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.SipResponseCodes = []int{0} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callhistoryparticipant) MarshalJSON() ([]byte, error) {
    type Alias Callhistoryparticipant

    if CallhistoryparticipantMarshalled {
        return []byte("{}"), nil
    }
    CallhistoryparticipantMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Address string `json:"address"`
        
        StartTime time.Time `json:"startTime"`
        
        EndTime time.Time `json:"endTime"`
        
        Purpose string `json:"purpose"`
        
        Direction string `json:"direction"`
        
        Ani string `json:"ani"`
        
        Dnis string `json:"dnis"`
        
        User User `json:"user"`
        
        Queue Queue `json:"queue"`
        
        Group Group `json:"group"`
        
        DisconnectType string `json:"disconnectType"`
        
        ExternalContact Externalcontact `json:"externalContact"`
        
        ExternalOrganization Externalorganization `json:"externalOrganization"`
        
        DidInteract bool `json:"didInteract"`
        
        SipResponseCodes []int `json:"sipResponseCodes"`
        
        FlaggedReason string `json:"flaggedReason"`
        
        OutboundCampaign Campaign `json:"outboundCampaign"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        SipResponseCodes: []int{0},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

