package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentguestsearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentguestsearchDud struct { 
    


    


    


    SearchId string `json:"searchId"`


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    


    SessionId string `json:"sessionId"`


    Results []Knowledgedocumentguestsearchresult `json:"results"`

}

// Knowledgedocumentguestsearch
type Knowledgedocumentguestsearch struct { 
    // Query - Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    


    


    


    // QueryType - The type of the query that initiates the search.
    QueryType string `json:"queryType"`


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentguestsearch) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentguestsearch) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentguestsearch

    if KnowledgedocumentguestsearchMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentguestsearchMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        QueryType string `json:"queryType"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

