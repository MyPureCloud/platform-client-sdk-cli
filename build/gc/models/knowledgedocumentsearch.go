package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsearchDud struct { 
    


    


    


    SearchId string `json:"searchId"`


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    Results []Knowledgedocumentsearchresult `json:"results"`


    

}

// Knowledgedocumentsearch
type Knowledgedocumentsearch struct { 
    // Query - Query to search content in the knowledge base. Maximum of 30 records per query can be fetched.
    Query string `json:"query"`


    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    


    


    


    


    // Application - The client application details from which search happened.
    Application Knowledgesearchclientapplication `json:"application"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsearch) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsearch) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsearch

    if KnowledgedocumentsearchMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsearchMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Application Knowledgesearchclientapplication `json:"application"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

