package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopilotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopilotDud struct { 
    Enabled bool `json:"enabled"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Copilot
type Copilot struct { 
    


    // LiveOnQueue - Copilot is live on selected queue.
    LiveOnQueue bool `json:"liveOnQueue"`


    // DefaultLanguage - Copilot default language, e.g. [en-US, es-US, es-ES]. Once set, it can not be modified.
    DefaultLanguage string `json:"defaultLanguage"`


    // KnowledgeAnswerConfig - Knowledge answer configuration.
    KnowledgeAnswerConfig Knowledgeanswerconfig `json:"knowledgeAnswerConfig"`


    // SummaryGenerationConfig - Copilot generated summary configuration.
    SummaryGenerationConfig Summarygenerationconfig `json:"summaryGenerationConfig"`


    // WrapupCodePredictionConfig - Copilot generated wrapup code prediction configuration.
    WrapupCodePredictionConfig Wrapupcodepredictionconfig `json:"wrapupCodePredictionConfig"`


    // AnswerGenerationConfig - Answer generation configuration.
    AnswerGenerationConfig Answergenerationconfig `json:"answerGenerationConfig"`


    // NluEngineType - Language understanding engine type.
    NluEngineType string `json:"nluEngineType"`


    // NluConfig - NLU configuration.
    NluConfig Nluconfig `json:"nluConfig"`


    // RuleEngineConfig - Rule engine configuration.
    RuleEngineConfig Ruleengineconfig `json:"ruleEngineConfig"`


    

}

// String returns a JSON representation of the model
func (o *Copilot) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copilot) MarshalJSON() ([]byte, error) {
    type Alias Copilot

    if CopilotMarshalled {
        return []byte("{}"), nil
    }
    CopilotMarshalled = true

    return json.Marshal(&struct {
        
        LiveOnQueue bool `json:"liveOnQueue"`
        
        DefaultLanguage string `json:"defaultLanguage"`
        
        KnowledgeAnswerConfig Knowledgeanswerconfig `json:"knowledgeAnswerConfig"`
        
        SummaryGenerationConfig Summarygenerationconfig `json:"summaryGenerationConfig"`
        
        WrapupCodePredictionConfig Wrapupcodepredictionconfig `json:"wrapupCodePredictionConfig"`
        
        AnswerGenerationConfig Answergenerationconfig `json:"answerGenerationConfig"`
        
        NluEngineType string `json:"nluEngineType"`
        
        NluConfig Nluconfig `json:"nluConfig"`
        
        RuleEngineConfig Ruleengineconfig `json:"ruleEngineConfig"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

