package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationsessionsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationsessionsummaryDud struct { 
    


    


    


    


    


    


    


    


    


    


    Id string `json:"id"`


    Confidence float32 `json:"confidence"`


    


    Communication Entity `json:"communication"`

}

// Conversationsessionsummary
type Conversationsessionsummary struct { 
    // Text - The text of the summary.
    Text string `json:"text"`


    // Status - The status of the conversation summary.
    Status string `json:"status"`


    // MediaType - The media type of the conversation.
    MediaType string `json:"mediaType"`


    // Language - The language of the conversation.
    Language string `json:"language"`


    // PredictedWrapupCodes - The wrapup codes of the conversation summary.
    PredictedWrapupCodes []Conversationsummarywrapupcode `json:"predictedWrapupCodes"`


    // EditedSummary - The edited summary of the conversation.
    EditedSummary Conversationeditedinput `json:"editedSummary"`


    // Reason - The reason of the conversation summary.
    Reason Conversationsummaryreason `json:"reason"`


    // Followup - The followup of the conversation summary.
    Followup Conversationsummaryfollowup `json:"followup"`


    // Resolution - The resolution of the conversation summary.
    Resolution Conversationsummaryresolution `json:"resolution"`


    // DateCreated - The created date of the summary. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    


    


    // Participants - The list of participants.
    Participants []Addressableentityref `json:"participants"`


    

}

// String returns a JSON representation of the model
func (o *Conversationsessionsummary) String() string {
    
    
    
    
     o.PredictedWrapupCodes = []Conversationsummarywrapupcode{{}} 
    
    
    
    
    
     o.Participants = []Addressableentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationsessionsummary) MarshalJSON() ([]byte, error) {
    type Alias Conversationsessionsummary

    if ConversationsessionsummaryMarshalled {
        return []byte("{}"), nil
    }
    ConversationsessionsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Status string `json:"status"`
        
        MediaType string `json:"mediaType"`
        
        Language string `json:"language"`
        
        PredictedWrapupCodes []Conversationsummarywrapupcode `json:"predictedWrapupCodes"`
        
        EditedSummary Conversationeditedinput `json:"editedSummary"`
        
        Reason Conversationsummaryreason `json:"reason"`
        
        Followup Conversationsummaryfollowup `json:"followup"`
        
        Resolution Conversationsummaryresolution `json:"resolution"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        Participants []Addressableentityref `json:"participants"`
        *Alias
    }{

        


        


        


        


        
        PredictedWrapupCodes: []Conversationsummarywrapupcode{{}},
        


        


        


        


        


        


        


        


        
        Participants: []Addressableentityref{{}},
        


        

        Alias: (*Alias)(u),
    })
}

