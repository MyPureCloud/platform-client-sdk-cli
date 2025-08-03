package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentchunkresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentchunkresponseDud struct { 
    


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    


    


    


    SearchId string `json:"searchId"`


    


    


    


    


    

}

// Knowledgedocumentchunkresponse
type Knowledgedocumentchunkresponse struct { 
    // Query - Query to search chunks in the knowledge base.
    Query string `json:"query"`


    


    


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    // QueryType - The type of the query that initiates the chunks search.
    QueryType string `json:"queryType"`


    


    // PreprocessQuery - Indicates whether the chunks search query should be preprocessed.
    PreprocessQuery bool `json:"preprocessQuery"`


    // ConfidenceThreshold - The confidence threshold for the chunk results. If applied, the returned results will have an equal or higher chunk confidence than the threshold.
    ConfidenceThreshold float32 `json:"confidenceThreshold"`


    // Results - Chunks matching the search query.
    Results []Documentchunkblock `json:"results"`


    // Application - The client application details from which chunks search happened.
    Application Knowledgesearchclientapplication `json:"application"`


    // ConversationContext - Conversation context information if the chunks search is initiated in the context of a conversation.
    ConversationContext Knowledgeconversationcontextresponse `json:"conversationContext"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentchunkresponse) String() string {
    
    
    
    
    
    
     o.Results = []Documentchunkblock{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentchunkresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentchunkresponse

    if KnowledgedocumentchunkresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentchunkresponseMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        QueryType string `json:"queryType"`
        
        PreprocessQuery bool `json:"preprocessQuery"`
        
        ConfidenceThreshold float32 `json:"confidenceThreshold"`
        
        Results []Documentchunkblock `json:"results"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        
        ConversationContext Knowledgeconversationcontextresponse `json:"conversationContext"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Results: []Documentchunkblock{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

