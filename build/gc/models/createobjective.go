package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateobjectiveMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateobjectiveDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    

}

// Createobjective
type Createobjective struct { 
    


    // TemplateId - The id of this objective's base template
    TemplateId string `json:"templateId"`


    // Zones - Objective zone specifies min,max points and values for the associated metric
    Zones []Objectivezone `json:"zones"`


    // Enabled - A flag for whether this objective is enabled for the related metric
    Enabled bool `json:"enabled"`


    // TopicIds - A list of topic ids for detected topic metrics
    TopicIds []string `json:"topicIds"`


    // MediaTypes - A list of media types for the metric
    MediaTypes []string `json:"mediaTypes"`


    // QueueIds - A list of queue ids for the metric
    QueueIds []string `json:"queueIds"`


    // TopicIdsFilterType - A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is \"or\".
    TopicIdsFilterType string `json:"topicIdsFilterType"`


    // EvaluationFormContextIds - The ids of associated evaluation form context, for Quality Evaluation Score metrics
    EvaluationFormContextIds []string `json:"evaluationFormContextIds"`


    // DateStart - start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`

}

// String returns a JSON representation of the model
func (o *Createobjective) String() string {
    
     o.Zones = []Objectivezone{{}} 
    
     o.TopicIds = []string{""} 
     o.MediaTypes = []string{""} 
     o.QueueIds = []string{""} 
    
     o.EvaluationFormContextIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createobjective) MarshalJSON() ([]byte, error) {
    type Alias Createobjective

    if CreateobjectiveMarshalled {
        return []byte("{}"), nil
    }
    CreateobjectiveMarshalled = true

    return json.Marshal(&struct {
        
        TemplateId string `json:"templateId"`
        
        Zones []Objectivezone `json:"zones"`
        
        Enabled bool `json:"enabled"`
        
        TopicIds []string `json:"topicIds"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        QueueIds []string `json:"queueIds"`
        
        TopicIdsFilterType string `json:"topicIdsFilterType"`
        
        EvaluationFormContextIds []string `json:"evaluationFormContextIds"`
        
        DateStart time.Time `json:"dateStart"`
        *Alias
    }{

        


        


        
        Zones: []Objectivezone{{}},
        


        


        
        TopicIds: []string{""},
        


        
        MediaTypes: []string{""},
        


        
        QueueIds: []string{""},
        


        


        
        EvaluationFormContextIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

