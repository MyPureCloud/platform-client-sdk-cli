package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TopiccriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TopiccriteriaDud struct { 
    


    

}

// Topiccriteria
type Topiccriteria struct { 
    // TopicId - The ID of the topic.
    TopicId string `json:"topicId"`


    // DataIngestionRules - The data ingestion rules for this topic.
    DataIngestionRules []Dataingestionrulecriteria `json:"dataIngestionRules"`

}

// String returns a JSON representation of the model
func (o *Topiccriteria) String() string {
    
     o.DataIngestionRules = []Dataingestionrulecriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Topiccriteria) MarshalJSON() ([]byte, error) {
    type Alias Topiccriteria

    if TopiccriteriaMarshalled {
        return []byte("{}"), nil
    }
    TopiccriteriaMarshalled = true

    return json.Marshal(&struct {
        
        TopicId string `json:"topicId"`
        
        DataIngestionRules []Dataingestionrulecriteria `json:"dataIngestionRules"`
        *Alias
    }{

        


        
        DataIngestionRules: []Dataingestionrulecriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

