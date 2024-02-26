package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsearchrequestDud struct { 
    


    


    


    SearchId string `json:"searchId"`


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    


    


    


    


    


    


    


    


    

}

// Knowledgedocumentsearchrequest
type Knowledgedocumentsearchrequest struct { 
    // Query - Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    


    


    


    // QueryType - The type of the query that initiates the search.
    QueryType string `json:"queryType"`


    // IncludeDraftDocuments - Indicates whether the search results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`


    // Interval - Retrieves the documents created/modified/published in specified date and time range.
    Interval Documentqueryinterval `json:"interval"`


    // Filter - Filter for the document search.
    Filter Documentquery `json:"filter"`


    // SortOrder - The sort order for search results.
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the documents that you want to sort the search results by.
    SortBy string `json:"sortBy"`


    // Application - The client application details from which search request was sent.
    Application Knowledgesearchclientapplication `json:"application"`


    // ConversationContext - Conversation context information if the search is initiated in the context of a conversation.
    ConversationContext Knowledgeconversationcontext `json:"conversationContext"`


    // ConfidenceThreshold - The confidence threshold for the search results. If applied, the returned results will have an equal or higher confidence than the threshold. The value should be between 0 to 1.
    ConfidenceThreshold float32 `json:"confidenceThreshold"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearchrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsearchrequest

    if KnowledgedocumentsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        QueryType string `json:"queryType"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        
        Interval Documentqueryinterval `json:"interval"`
        
        Filter Documentquery `json:"filter"`
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        
        ConversationContext Knowledgeconversationcontext `json:"conversationContext"`
        
        ConfidenceThreshold float32 `json:"confidenceThreshold"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

