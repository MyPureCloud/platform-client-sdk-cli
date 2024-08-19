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


    Results []Knowledgedocumentresponse `json:"results"`

}

// Knowledgedocumentqueryresponse
type Knowledgedocumentqueryresponse struct { 
    // PageSize - Page size of the returned results.
    PageSize int `json:"pageSize"`


    // PageNumber - Page number of the returned results.
    PageNumber int `json:"pageNumber"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentqueryresponse) String() string {
    
    

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
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

