package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentchunkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentchunkrequestDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Knowledgedocumentchunkrequest
type Knowledgedocumentchunkrequest struct { 
    // Query - Query to search chunks in the knowledge base.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    // Filter - Filter for the document chunks.
    Filter Documentquery `json:"filter"`


    // QueryType - The type of the query that initiates the chunks search.
    QueryType string `json:"queryType"`


    // PreprocessQuery - Indicates whether the chunks search query should be preprocessed.
    PreprocessQuery bool `json:"preprocessQuery"`


    // IncludeDraftDocuments - Indicates whether the chunk results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`


    // Application - The client application details from which chunks request was sent.
    Application Knowledgesearchclientapplication `json:"application"`


    // ConversationContext - Conversation context information if the chunks search is initiated in the context of a conversation.
    ConversationContext Knowledgeconversationcontext `json:"conversationContext"`


    // ConfidenceThreshold - The confidence threshold for the chunk results. If applied, the returned results will have an equal or higher confidence than the threshold. The value should be between 0 to 1.
    ConfidenceThreshold float32 `json:"confidenceThreshold"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentchunkrequest) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentchunkrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentchunkrequest

    if KnowledgedocumentchunkrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentchunkrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Filter Documentquery `json:"filter"`
        
        QueryType string `json:"queryType"`
        
        PreprocessQuery bool `json:"preprocessQuery"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        
        ConversationContext Knowledgeconversationcontext `json:"conversationContext"`
        
        ConfidenceThreshold float32 `json:"confidenceThreshold"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

