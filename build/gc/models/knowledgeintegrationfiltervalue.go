package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeintegrationfiltervalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeintegrationfiltervalueDud struct { 
    


    

}

// Knowledgeintegrationfiltervalue
type Knowledgeintegrationfiltervalue struct { 
    // Key - The key that can be used as a value of a filter setting in a knowledge source.
    Key string `json:"key"`


    // Value - The display label of the filter value.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Knowledgeintegrationfiltervalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeintegrationfiltervalue) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeintegrationfiltervalue

    if KnowledgeintegrationfiltervalueMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeintegrationfiltervalueMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

