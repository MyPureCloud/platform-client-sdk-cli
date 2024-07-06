package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionDud struct { 
    Id string `json:"id"`


    VarType string `json:"type"`


    Faq Faq `json:"faq"`


    Article Article `json:"article"`


    DateCreated time.Time `json:"dateCreated"`


    AnswerRecordId string `json:"answerRecordId"`


    TriggerType string `json:"triggerType"`


    Context Suggestioncontext `json:"context"`


    State string `json:"state"`


    KnowledgeSearch Suggestionknowledgesearch `json:"knowledgeSearch"`


    KnowledgeArticle Suggestionknowledgearticle `json:"knowledgeArticle"`


    CannedResponse Suggestioncannedresponse `json:"cannedResponse"`


    Script Suggestionscript `json:"script"`


    SelfUri string `json:"selfUri"`


    Conversation Addressableentityref `json:"conversation"`


    Assistant Addressableentityref `json:"assistant"`

}

// Suggestion
type Suggestion struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestion) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestion) MarshalJSON() ([]byte, error) {
    type Alias Suggestion

    if SuggestionMarshalled {
        return []byte("{}"), nil
    }
    SuggestionMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

