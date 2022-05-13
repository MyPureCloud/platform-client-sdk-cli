package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ObjectiveMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ObjectiveDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    

}

// Objective
type Objective struct { 
    


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


    // DateStart - start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    DateStart time.Time `json:"dateStart"`

}

// String returns a JSON representation of the model
func (o *Objective) String() string {
    
     o.Zones = []Objectivezone{{}} 
    
     o.MediaTypes = []string{""} 
     o.Queues = []Addressableentityref{{}} 
     o.Topics = []Addressableentityref{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Objective) MarshalJSON() ([]byte, error) {
    type Alias Objective

    if ObjectiveMarshalled {
        return []byte("{}"), nil
    }
    ObjectiveMarshalled = true

    return json.Marshal(&struct {
        
        TemplateId string `json:"templateId"`
        
        Zones []Objectivezone `json:"zones"`
        
        Enabled bool `json:"enabled"`
        
        MediaTypes []string `json:"mediaTypes"`
        
        Queues []Addressableentityref `json:"queues"`
        
        Topics []Addressableentityref `json:"topics"`
        
        TopicIdsFilterType string `json:"topicIdsFilterType"`
        
        DateStart time.Time `json:"dateStart"`
        *Alias
    }{

        


        


        
        Zones: []Objectivezone{{}},
        


        


        
        MediaTypes: []string{""},
        


        
        Queues: []Addressableentityref{{}},
        


        
        Topics: []Addressableentityref{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

