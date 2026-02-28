package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchpreviewresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchpreviewresponseDud struct { 
    


    SearchId string `json:"searchId"`


    SessionId string `json:"sessionId"`


    Result Knowledgesearchresult `json:"result"`

}

// Knowledgesearchpreviewresponse
type Knowledgesearchpreviewresponse struct { 
    // Query - Query to search content in the knowledge base.
    Query string `json:"query"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgesearchpreviewresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchpreviewresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchpreviewresponse

    if KnowledgesearchpreviewresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchpreviewresponseMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

