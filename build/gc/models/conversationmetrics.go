package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationmetricsDud struct { 
    


    


    


    


    

}

// Conversationmetrics
type Conversationmetrics struct { 
    // Conversation - The Conversation Reference
    Conversation Addressableentityref `json:"conversation"`


    // SentimentScore - The Sentiment Score
    SentimentScore float64 `json:"sentimentScore"`


    // SentimentTrend - The Sentiment Trend
    SentimentTrend float64 `json:"sentimentTrend"`


    // SentimentTrendClass - The Sentiment Trend Class
    SentimentTrendClass string `json:"sentimentTrendClass"`


    // ParticipantMetrics - The Participant Metrics
    ParticipantMetrics Participantmetrics `json:"participantMetrics"`

}

// String returns a JSON representation of the model
func (o *Conversationmetrics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationmetrics) MarshalJSON() ([]byte, error) {
    type Alias Conversationmetrics

    if ConversationmetricsMarshalled {
        return []byte("{}"), nil
    }
    ConversationmetricsMarshalled = true

    return json.Marshal(&struct { 
        Conversation Addressableentityref `json:"conversation"`
        
        SentimentScore float64 `json:"sentimentScore"`
        
        SentimentTrend float64 `json:"sentimentTrend"`
        
        SentimentTrendClass string `json:"sentimentTrendClass"`
        
        ParticipantMetrics Participantmetrics `json:"participantMetrics"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

