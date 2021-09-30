package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsparticipantwithoutattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsparticipantwithoutattributesDud struct { 
    


    


    


    


    


    


    


    


    

}

// Analyticsparticipantwithoutattributes
type Analyticsparticipantwithoutattributes struct { 
    // ExternalContactId - External contact identifier
    ExternalContactId string `json:"externalContactId"`


    // ExternalOrganizationId - External organization identifier
    ExternalOrganizationId string `json:"externalOrganizationId"`


    // FlaggedReason - Reason for which participant flagged conversation
    FlaggedReason string `json:"flaggedReason"`


    // ParticipantId - Unique identifier for the participant
    ParticipantId string `json:"participantId"`


    // ParticipantName - A human readable name identifying the participant
    ParticipantName string `json:"participantName"`


    // Purpose - The participant's purpose
    Purpose string `json:"purpose"`


    // TeamId - The team ID the user is a member of
    TeamId string `json:"teamId"`


    // UserId - Unique identifier for the user
    UserId string `json:"userId"`


    // Sessions - List of sessions associated to this participant
    Sessions []Analyticssession `json:"sessions"`

}

// String returns a JSON representation of the model
func (o *Analyticsparticipantwithoutattributes) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Sessions = []Analyticssession{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsparticipantwithoutattributes) MarshalJSON() ([]byte, error) {
    type Alias Analyticsparticipantwithoutattributes

    if AnalyticsparticipantwithoutattributesMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsparticipantwithoutattributesMarshalled = true

    return json.Marshal(&struct { 
        ExternalContactId string `json:"externalContactId"`
        
        ExternalOrganizationId string `json:"externalOrganizationId"`
        
        FlaggedReason string `json:"flaggedReason"`
        
        ParticipantId string `json:"participantId"`
        
        ParticipantName string `json:"participantName"`
        
        Purpose string `json:"purpose"`
        
        TeamId string `json:"teamId"`
        
        UserId string `json:"userId"`
        
        Sessions []Analyticssession `json:"sessions"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Sessions: []Analyticssession{{}},
        

        
        Alias: (*Alias)(u),
    })
}

