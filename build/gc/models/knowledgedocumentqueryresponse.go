package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentqueryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentqueryresponseDud struct { 
    


    


    Total int `json:"total"`


    PageCount int `json:"pageCount"`


    

}

// Knowledgedocumentqueryresponse
type Knowledgedocumentqueryresponse struct { 
    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    


    


    // Results - Documents matching the query.
    Results []Knowledgedocumentqueryresult `json:"results"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentqueryresponse) String() string {
    
    
     o.Results = []Knowledgedocumentqueryresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentqueryresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentqueryresponse

    if KnowledgedocumentqueryresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentqueryresponseMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Results []Knowledgedocumentqueryresult `json:"results"`
        *Alias
    }{

        


        


        


        


        
        Results: []Knowledgedocumentqueryresult{{}},
        

        Alias: (*Alias)(u),
    })
}

