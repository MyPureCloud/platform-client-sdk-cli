package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentguestsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentguestsearchrequestDud struct { 
    


    


    


    SearchId string `json:"searchId"`


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    


    SessionId string `json:"sessionId"`


    

}

// Knowledgedocumentguestsearchrequest
type Knowledgedocumentguestsearchrequest struct { 
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

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearchrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentguestsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentguestsearchrequest

    if KnowledgedocumentguestsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentguestsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        QueryType string `json:"queryType"`
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

