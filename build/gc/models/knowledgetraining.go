package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgetrainingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgetrainingDud struct { 
    Id string `json:"id"`


    DateTriggered time.Time `json:"dateTriggered"`


    DateCompleted time.Time `json:"dateCompleted"`


    Status string `json:"status"`


    LanguageCode string `json:"languageCode"`


    KnowledgeBase Knowledgebase `json:"knowledgeBase"`


    ErrorMessage string `json:"errorMessage"`


    KnowledgeDocumentsState string `json:"knowledgeDocumentsState"`


    DatePromoted time.Time `json:"datePromoted"`


    SelfUri string `json:"selfUri"`

}

// Knowledgetraining
type Knowledgetraining struct { 
    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgetraining) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgetraining) MarshalJSON() ([]byte, error) {
    type Alias Knowledgetraining

    if KnowledgetrainingMarshalled {
        return []byte("{}"), nil
    }
    KnowledgetrainingMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

