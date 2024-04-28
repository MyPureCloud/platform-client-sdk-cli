package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentfeedbackresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentfeedbackresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SessionId string `json:"sessionId"`


    DateCreated time.Time `json:"dateCreated"`


    


    


    


    


    


    


    User Addressableentityref `json:"user"`


    SelfUri string `json:"selfUri"`

}

// Knowledgedocumentfeedbackresponse
type Knowledgedocumentfeedbackresponse struct { 
    


    // DocumentVariation - The variation of the document on which feedback was given.
    DocumentVariation Entityreference `json:"documentVariation"`


    // Rating - Feedback rating.
    Rating string `json:"rating"`


    // Reason - Feedback reason.
    Reason string `json:"reason"`


    // Comment - Free-text comment of the feedback. Maximum length: 2000 characters.
    Comment string `json:"comment"`


    // Search - The search that surfaced the document on which feedback was given.
    Search Entityreference `json:"search"`


    


    


    // QueryType - The type of the query that surfaced the document on which the feedback was given.
    QueryType string `json:"queryType"`


    // SurfacingMethod - The method how knowledge was surfaced. Article: Full article was shown. Snippet: A snippet from the article was shown. Highlight: A highlighted answer in a snippet was shown.
    SurfacingMethod string `json:"surfacingMethod"`


    // State - The state of the feedback.
    State string `json:"state"`


    // Document - The document on which feedback was given.
    Document Knowledgedocumentversionreference `json:"document"`


    // Application - The client application from which feedback was given.
    Application Knowledgesearchclientapplication `json:"application"`


    // ConversationContext - Conversation context information if the feedback is given in the context of a conversation.
    ConversationContext Knowledgeconversationcontextresponse `json:"conversationContext"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentfeedbackresponse) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentfeedbackresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentfeedbackresponse

    if KnowledgedocumentfeedbackresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentfeedbackresponseMarshalled = true

    return json.Marshal(&struct {
        
        DocumentVariation Entityreference `json:"documentVariation"`
        
        Rating string `json:"rating"`
        
        Reason string `json:"reason"`
        
        Comment string `json:"comment"`
        
        Search Entityreference `json:"search"`
        
        QueryType string `json:"queryType"`
        
        SurfacingMethod string `json:"surfacingMethod"`
        
        State string `json:"state"`
        
        Document Knowledgedocumentversionreference `json:"document"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        
        ConversationContext Knowledgeconversationcontextresponse `json:"conversationContext"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

