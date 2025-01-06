package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstateagentsessionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstateagentsessionresultDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Analyticsagentstateagentsessionresult
type Analyticsagentstateagentsessionresult struct { 
    // ConversationId - Conversation Id
    ConversationId string `json:"conversationId"`


    // SessionId - Session Id
    SessionId string `json:"sessionId"`


    // SessionStart - Session start datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SessionStart time.Time `json:"sessionStart"`


    // SegmentStart - Segment start datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SegmentStart time.Time `json:"segmentStart"`


    // SegmentType - Segment type
    SegmentType string `json:"segmentType"`


    // RoutedQueueId - Routed queue Id
    RoutedQueueId string `json:"routedQueueId"`


    // RequestedRoutingSkillIds - List of requested routing skill Id
    RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`


    // RequestedLanguageId - Requested language Id
    RequestedLanguageId string `json:"requestedLanguageId"`


    // OriginatingDirection - Originating direction
    OriginatingDirection string `json:"originatingDirection"`


    // MediaType - Media type
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentsessionresult) String() string {
    
    
    
    
    
    
     o.RequestedRoutingSkillIds = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstateagentsessionresult) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstateagentsessionresult

    if AnalyticsagentstateagentsessionresultMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstateagentsessionresultMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        SessionId string `json:"sessionId"`
        
        SessionStart time.Time `json:"sessionStart"`
        
        SegmentStart time.Time `json:"segmentStart"`
        
        SegmentType string `json:"segmentType"`
        
        RoutedQueueId string `json:"routedQueueId"`
        
        RequestedRoutingSkillIds []string `json:"requestedRoutingSkillIds"`
        
        RequestedLanguageId string `json:"requestedLanguageId"`
        
        OriginatingDirection string `json:"originatingDirection"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        


        


        


        


        
        RequestedRoutingSkillIds: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

