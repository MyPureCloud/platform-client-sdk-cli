package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnsweroptionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnsweroptionDud struct { 
    


    


    


    


    

}

// Answeroption
type Answeroption struct { 
    // Id
    Id string `json:"id"`


    // BuiltInType - The built-in type of this answer option. Only used for built-in answer options such as selection states for Multiple Select answer options. Possible values include: Selected, Unselected
    BuiltInType string `json:"builtInType"`


    // Text
    Text string `json:"text"`


    // Value
    Value int `json:"value"`


    // AssistanceConditions - List of assistance conditions which are combined together with a logical AND operator. Eg ( assistanceCondtion1 && assistanceCondition2 ) wherein assistanceCondition could be ( EXISTS topic1 || topic2 || ... ) or (NOTEXISTS topic3 || topic4 || ...).
    AssistanceConditions []Assistancecondition `json:"assistanceConditions"`

}

// String returns a JSON representation of the model
func (o *Answeroption) String() string {
    
    
    
    
     o.AssistanceConditions = []Assistancecondition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Answeroption) MarshalJSON() ([]byte, error) {
    type Alias Answeroption

    if AnsweroptionMarshalled {
        return []byte("{}"), nil
    }
    AnsweroptionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        BuiltInType string `json:"builtInType"`
        
        Text string `json:"text"`
        
        Value int `json:"value"`
        
        AssistanceConditions []Assistancecondition `json:"assistanceConditions"`
        *Alias
    }{

        


        


        


        


        
        AssistanceConditions: []Assistancecondition{{}},
        

        Alias: (*Alias)(u),
    })
}

