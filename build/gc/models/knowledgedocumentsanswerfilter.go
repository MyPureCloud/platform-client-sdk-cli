package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgedocumentsanswerfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgedocumentsanswerfilterDud struct { 
    


    


    


    


    


    


    


    

}

// Knowledgedocumentsanswerfilter
type Knowledgedocumentsanswerfilter struct { 
    // Query - The search query.
    Query string `json:"query"`


    // Language - The language of the documents.
    Language string `json:"language"`


    // AppType - The appType
    AppType string `json:"appType"`


    // QueryType - The query type
    QueryType string `json:"queryType"`


    // SearchId - The search id.
    SearchId string `json:"searchId"`


    // InsertHighlightIntoVariationContent - If specified - insert highlight data into the variation content.
    InsertHighlightIntoVariationContent bool `json:"insertHighlightIntoVariationContent"`


    // AnswerMode - Allows extracted answers from an article (AnswerHighlight) and/or AI-generated answers (AnswerGeneration). Default mode: AnswerHighlight
    AnswerMode []string `json:"answerMode"`


    // VariationIds - The variation Ids to answer.
    VariationIds []string `json:"variationIds"`

}

// String returns a JSON representation of the model
func (o *Knowledgedocumentsanswerfilter) String() string {
    
    
    
    
    
    
     o.AnswerMode = []string{""} 
     o.VariationIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgedocumentsanswerfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgedocumentsanswerfilter

    if KnowledgedocumentsanswerfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgedocumentsanswerfilterMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        Language string `json:"language"`
        
        AppType string `json:"appType"`
        
        QueryType string `json:"queryType"`
        
        SearchId string `json:"searchId"`
        
        InsertHighlightIntoVariationContent bool `json:"insertHighlightIntoVariationContent"`
        
        AnswerMode []string `json:"answerMode"`
        
        VariationIds []string `json:"variationIds"`
        *Alias
    }{

        


        


        


        


        


        


        
        AnswerMode: []string{""},
        


        
        VariationIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

