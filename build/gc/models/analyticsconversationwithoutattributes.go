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
    // ConversationEnd - The end time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationEnd time.Time `json:"conversationEnd"`


    // ConversationId - Unique identifier for the conversation
    ConversationId string `json:"conversationId"`


    // ConversationStart - The start time of a conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ConversationStart time.Time `json:"conversationStart"`


    // DivisionIds - Identifier(s) of division(s) associated with a conversation
    DivisionIds []string `json:"divisionIds"`


    // ExternalTag - External tag for the conversation
    ExternalTag string `json:"externalTag"`


    // MediaStatsMinConversationMos - The lowest estimated average MOS among all the audio streams belonging to this conversation
    MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos"`


    // MediaStatsMinConversationRFactor - The lowest R-factor value among all of the audio streams belonging to this conversation
    MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor"`


    // OriginatingDirection - The original direction of the conversation
    OriginatingDirection string `json:"originatingDirection"`


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
        ConversationEnd time.Time `json:"conversationEnd"`
        
        ConversationId string `json:"conversationId"`
        
        ConversationStart time.Time `json:"conversationStart"`
        
        DivisionIds []string `json:"divisionIds"`
        
        ExternalTag string `json:"externalTag"`
        
        MediaStatsMinConversationMos float64 `json:"mediaStatsMinConversationMos"`
        
        MediaStatsMinConversationRFactor float64 `json:"mediaStatsMinConversationRFactor"`
        
        OriginatingDirection string `json:"originatingDirection"`
        
        Evaluations []Analyticsevaluation `json:"evaluations"`
        
        Surveys []Analyticssurvey `json:"surveys"`
        
        Resolutions []Analyticsresolution `json:"resolutions"`
        
        Participants []Analyticsparticipantwithoutattributes `json:"participants"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        DivisionIds: []string{""},
        

        

        

        

        

        

        

        

        

        

        
        Evaluations: []Analyticsevaluation{{}},
        

        

        
        Surveys: []Analyticssurvey{{}},
        

        

        
        Resolutions: []Analyticsresolution{{}},
        

        

        
        Participants: []Analyticsparticipantwithoutattributes{{}},
        

        
        Alias: (*Alias)(u),
    })
}

