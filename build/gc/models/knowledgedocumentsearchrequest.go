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


    


    


    


    // IncludeDraftDocuments - Indicates whether the search results would also include draft documents.
    IncludeDraftDocuments bool `json:"includeDraftDocuments"`

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
        
        IncludeDraftDocuments bool `json:"includeDraftDocuments"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

