package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgemetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgemetadataDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Reportingturnknowledgemetadata
type Reportingturnknowledgemetadata struct { 
    // KnowledgeId - The ID of the knowledge setting or knowledge base
    KnowledgeId string `json:"knowledgeId"`


    // KnowledgeName - The name of the knowledge setting or knowledge base
    KnowledgeName string `json:"knowledgeName"`


    // SearchId - SearchID used in the attempted search
    SearchId string `json:"searchId"`


    // Query - The query used in the knowledge query
    Query string `json:"query"`


    // RetrievalStatus - The result of the knowledge search
    RetrievalStatus string `json:"retrievalStatus"`


    // AnswerGenerationStatus - The result of the knowledge generation
    AnswerGenerationStatus string `json:"answerGenerationStatus"`


    // GeneratedAnswer - The generated answer
    GeneratedAnswer string `json:"generatedAnswer"`


    // FailureReason - Failure reason if knowledge query failed
    FailureReason string `json:"failureReason"`


    // TopConfidence - Highest confidence score of returned knowledgeSources
    TopConfidence float64 `json:"topConfidence"`


    // RetrievedSources - List of the sources retrieved by the knowledge search
    RetrievedSources []Knowledgesource `json:"retrievedSources"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgemetadata) String() string {
    
    
    
    
    
    
    
    
    
     o.RetrievedSources = []Knowledgesource{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgemetadata) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgemetadata

    if ReportingturnknowledgemetadataMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgemetadataMarshalled = true

    return json.Marshal(&struct {
        
        KnowledgeId string `json:"knowledgeId"`
        
        KnowledgeName string `json:"knowledgeName"`
        
        SearchId string `json:"searchId"`
        
        Query string `json:"query"`
        
        RetrievalStatus string `json:"retrievalStatus"`
        
        AnswerGenerationStatus string `json:"answerGenerationStatus"`
        
        GeneratedAnswer string `json:"generatedAnswer"`
        
        FailureReason string `json:"failureReason"`
        
        TopConfidence float64 `json:"topConfidence"`
        
        RetrievedSources []Knowledgesource `json:"retrievedSources"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        RetrievedSources: []Knowledgesource{{}},
        

        Alias: (*Alias)(u),
    })
}

