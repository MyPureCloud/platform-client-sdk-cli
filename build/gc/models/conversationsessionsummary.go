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
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Conversationsessionsummary
type Conversationsessionsummary struct { 
    // Id - The id of the summary.
    Id string `json:"id"`


    // Text - The text of the summary.
    Text string `json:"text"`


    // Confidence - The AI confidence value.
    Confidence float32 `json:"confidence"`


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


    // LatestCommunication - The most recent communication whose corresponding content was used for generating the summary. This field will be null if not available.
    LatestCommunication Entity `json:"latestCommunication"`


    // Labels - The summary labels.
    Labels []Conversationsummarylabel `json:"labels"`


    // Queue - The queue that handled the conversation, if any.
    Queue Addressableentityref `json:"queue"`


    // TriggerSource - The configuration source that triggered summarization.
    TriggerSource Conversationsummarytriggersource `json:"triggerSource"`


    // SummarySourceType - The service used to generate the summary.
    SummarySourceType string `json:"summarySourceType"`


    // TriggerType - How summarization was triggered, either on-demand by an agent or automatically after the conversation disconnect.
    TriggerType string `json:"triggerType"`


    // LastEditedBy - The user that last edited the summary, if the summary was edited.
    LastEditedBy Conversationsummaryeditedby `json:"lastEditedBy"`


    // Communication - The communication object of the summary.
    Communication Entity `json:"communication"`


    // Participants - The list of participants.
    Participants []Addressableentityref `json:"participants"`

}

// String returns a JSON representation of the model
func (o *Conversationsessionsummary) String() string {
    
    
    
    
    
    
     o.PredictedWrapupCodes = []Conversationsummarywrapupcode{{}} 
    
    
    
    
    
    
     o.Labels = []Conversationsummarylabel{{}} 
    
    
    
    
    
    
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
        
        Id string `json:"id"`
        
        Text string `json:"text"`
        
        Confidence float32 `json:"confidence"`
        
        Status string `json:"status"`
        
        MediaType string `json:"mediaType"`
        
        Language string `json:"language"`
        
        PredictedWrapupCodes []Conversationsummarywrapupcode `json:"predictedWrapupCodes"`
        
        EditedSummary Conversationeditedinput `json:"editedSummary"`
        
        Reason Conversationsummaryreason `json:"reason"`
        
        Followup Conversationsummaryfollowup `json:"followup"`
        
        Resolution Conversationsummaryresolution `json:"resolution"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        LatestCommunication Entity `json:"latestCommunication"`
        
        Labels []Conversationsummarylabel `json:"labels"`
        
        Queue Addressableentityref `json:"queue"`
        
        TriggerSource Conversationsummarytriggersource `json:"triggerSource"`
        
        SummarySourceType string `json:"summarySourceType"`
        
        TriggerType string `json:"triggerType"`
        
        LastEditedBy Conversationsummaryeditedby `json:"lastEditedBy"`
        
        Communication Entity `json:"communication"`
        
        Participants []Addressableentityref `json:"participants"`
        *Alias
    }{

        


        


        


        


        


        


        
        PredictedWrapupCodes: []Conversationsummarywrapupcode{{}},
        


        


        


        


        


        


        


        
        Labels: []Conversationsummarylabel{{}},
        


        


        


        


        


        


        


        
        Participants: []Addressableentityref{{}},
        

        Alias: (*Alias)(u),
    })
}

