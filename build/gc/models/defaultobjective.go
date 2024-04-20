package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DefaultobjectiveMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DefaultobjectiveDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    

}

// Defaultobjective
type Defaultobjective struct { 
    


    // TemplateId - The id of this objective's base template
    TemplateId string `json:"templateId"`


    // Zones - Objective zone specifies min,max points and values for the associated metric
    Zones []Objectivezone `json:"zones"`


    // Enabled - A flag for whether this objective is enabled for the related metric
    Enabled bool `json:"enabled"`


    // MediaTypes - A list of media types for the metric
    MediaTypes []string `json:"mediaTypes"`


    // Queues - A list of queues for the metric
    Queues []Addressableentityref `json:"queues"`


    // Topics - A list of topic ids for detected topic metrics
    Topics []Addressableentityref `json:"topics"`


    // TopicIdsFilterType - A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is \"or\".
    TopicIdsFilterType string `json:"topicIdsFilterType"`


    // EvaluationFormContextIds - The ids of associated evaluation form context, for Quality Evaluation Score metrics
    EvaluationFormContextIds []string `json:"evaluationFormContextIds"`


    // InitialDirection - The initial direction to filter on
    InitialDirection string `json:"initialDirection"`

}

// String returns a JSON representation of the model
func (o *Defaultobjective) String() string {
    
     o.Zones = []Objectivezone{{}} 
    
     o.MediaTypes = []string{""} 
     o.Queues = []Addressableentityref{{}} 
     o.Topics = []Addressableentityref{{}} 
    
     o.EvaluationFormContextIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Defaultobjective) MarshalJSON() ([]byte, error) {
    type Alias Defaultobjective

    if DefaultobjectiveMarshalled {
        return []byte("{}"), nil
    }
    DefaultobjectiveMarshalled = true

    return json.Marshal(&struct {
        
        TemplateId string `json:"templateId"`
        
        Zones []Objectivezone `json:"zones"`
        
        Enabled bool `json:"enabled"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        Queues []Addressableentityref `json:"queues"`
        
        Topics []Addressableentityref `json:"topics"`
        
        TopicIdsFilterType string `json:"topicIdsFilterType"`
        
        EvaluationFormContextIds []string `json:"evaluationFormContextIds"`
        
        InitialDirection string `json:"initialDirection"`
        *Alias
    }{

        


        


        
        Zones: []Objectivezone{{}},
        


        


        
        MediaTypes: []string{""},
        


        
        Queues: []Addressableentityref{{}},
        


        
        Topics: []Addressableentityref{{}},
        


        


        
        EvaluationFormContextIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

