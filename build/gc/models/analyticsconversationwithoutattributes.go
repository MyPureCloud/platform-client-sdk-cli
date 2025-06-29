package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsconversationwithoutattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsconversationwithoutattributesDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsconversationwithoutattributes
type Analyticsconversationwithoutattributes struct { 
    // ConferenceStart - The start time of a conference call. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConferenceStart time.Time `json:"conferenceStart"`


    // ConversationEnd - The end time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationEnd time.Time `json:"conversationEnd"`


    // ConversationId - Unique identifier for the conversation
    ConversationId string `json:"conversationId"`


    // ConversationInitiator - Indicates the participant purpose of the participant initiating a message conversation
    ConversationInitiator string `json:"conversationInitiator"`


    // ConversationStart - The start time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationStart time.Time `json:"conversationStart"`


    // CustomerParticipation - Indicates a messaging conversation in which the customer participated by sending at least one message
    CustomerParticipation bool `json:"customerParticipation"`


    // DivisionIds - Identifier(s) of division(s) associated with a conversation
    DivisionIds []string `json:"divisionIds"`


    // ExternalTag - External tag for the conversation
    ExternalTag string `json:"externalTag"`


    // InactivityTimeout - The time in the future, after which this conversation would be considered inactive. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    InactivityTimeout time.Time `json:"inactivityTimeout"`


    // KnowledgeBaseIds - The unique identifier(s) of the knowledge base(s) used
    KnowledgeBaseIds []string `json:"knowledgeBaseIds"`


    // MediaStatsMinConversationMos - The lowest estimated average MOS among all the audio streams belonging to this conversation
    MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos"`


    // MediaStatsMinConversationRFactor - The lowest R-factor value among all of the audio streams belonging to this conversation
    MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor"`


    // OriginatingDirection - The original direction of the conversation
    OriginatingDirection string `json:"originatingDirection"`


    // OriginatingSocialMediaPublic - Indicates that the conversation originated from a public message on social media
    OriginatingSocialMediaPublic bool `json:"originatingSocialMediaPublic"`


    // SelfServed - Indicates whether all flow sessions were self serviced
    SelfServed bool `json:"selfServed"`


    // Evaluations - Evaluations associated with this conversation
    Evaluations []Analyticsevaluation `json:"evaluations"`


    // Surveys - Surveys associated with this conversation
    Surveys []Analyticssurvey `json:"surveys"`


    // Resolutions - Resolutions associated with this conversation
    Resolutions []Analyticsresolution `json:"resolutions"`


    // Participants - Participants in the conversation
    Participants []Analyticsparticipantwithoutattributes `json:"participants"`

}

// String returns a JSON representation of the model
func (o *Analyticsconversationwithoutattributes) String() string {
    
    
    
    
    
    
     o.DivisionIds = []string{""} 
    
    
     o.KnowledgeBaseIds = []string{""} 
    
    
    
    
    
     o.Evaluations = []Analyticsevaluation{{}} 
     o.Surveys = []Analyticssurvey{{}} 
     o.Resolutions = []Analyticsresolution{{}} 
     o.Participants = []Analyticsparticipantwithoutattributes{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsconversationwithoutattributes) MarshalJSON() ([]byte, error) {
    type Alias Analyticsconversationwithoutattributes

    if AnalyticsconversationwithoutattributesMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsconversationwithoutattributesMarshalled = true

    return json.Marshal(&struct {
        
        ConferenceStart time.Time `json:"conferenceStart"`
        
        ConversationEnd time.Time `json:"conversationEnd"`
        
        ConversationId string `json:"conversationId"`
        
        ConversationInitiator string `json:"conversationInitiator"`
        
        ConversationStart time.Time `json:"conversationStart"`
        
        CustomerParticipation bool `json:"customerParticipation"`
        
        DivisionIds []string `json:"divisionIds"`
        
        ExternalTag string `json:"externalTag"`
        
        InactivityTimeout time.Time `json:"inactivityTimeout"`
        
        KnowledgeBaseIds []string `json:"knowledgeBaseIds"`
        
        MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos"`
        
        MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor"`
        
        OriginatingDirection string `json:"originatingDirection"`
        
        OriginatingSocialMediaPublic bool `json:"originatingSocialMediaPublic"`
        
        SelfServed bool `json:"selfServed"`
        
        Evaluations []Analyticsevaluation `json:"evaluations"`
        
        Surveys []Analyticssurvey `json:"surveys"`
        
        Resolutions []Analyticsresolution `json:"resolutions"`
        
        Participants []Analyticsparticipantwithoutattributes `json:"participants"`
        *Alias
    }{

        


        


        


        


        


        


        
        DivisionIds: []string{""},
        


        


        


        
        KnowledgeBaseIds: []string{""},
        


        


        


        


        


        


        
        Evaluations: []Analyticsevaluation{{}},
        


        
        Surveys: []Analyticssurvey{{}},
        


        
        Resolutions: []Analyticsresolution{{}},
        


        
        Participants: []Analyticsparticipantwithoutattributes{{}},
        

        Alias: (*Alias)(u),
    })
}

