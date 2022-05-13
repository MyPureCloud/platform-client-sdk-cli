package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchrequestDud struct { 
    


    


    


    


    


    

}

// Knowledgesearchrequest
type Knowledgesearchrequest struct { 
    // Query - Input query to search content in the knowledge base
    Query string `json:"query"`


    // PageSize - Page size of the returned results
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results
    PageNumber int `json:"pageNumber"`


    // DocumentType - Document type to be used while searching
    DocumentType string `json:"documentType"`


    // LanguageCode - query search for specific languageCode
    LanguageCode string `json:"languageCode"`


    // SearchOnDraftDocuments - If true the search query will be executed on draft documents, else it will be on active documents
    SearchOnDraftDocuments bool `json:"searchOnDraftDocuments"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchrequest

    if KnowledgesearchrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        DocumentType string `json:"documentType"`
        
        LanguageCode string `json:"languageCode"`
        
        SearchOnDraftDocuments bool `json:"searchOnDraftDocuments"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

