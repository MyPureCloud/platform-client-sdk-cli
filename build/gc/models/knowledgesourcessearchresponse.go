package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesourcessearchresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesourcessearchresponseDud struct { 
    


    SearchId string `json:"searchId"`


    SessionId string `json:"sessionId"`


    Result Knowledgesearchresult `json:"result"`


    KnowledgeSettingId string `json:"knowledgeSettingId"`


    ConversationContext Knowledgev3conversationcontextresponse `json:"conversationContext"`


    Application V3knowledgesearchclientapplication `json:"application"`


    QueryType string `json:"queryType"`


    GenerationLanguage string `json:"generationLanguage"`


    AnswerGeneration bool `json:"answerGeneration"`

}

// Knowledgesourcessearchresponse
type Knowledgesourcessearchresponse struct { 
    // Query - Query to search content in the knowledge base.
    Query string `json:"query"`


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgesourcessearchresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesourcessearchresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesourcessearchresponse

    if KnowledgesourcessearchresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesourcessearchresponseMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

