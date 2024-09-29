package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeintegrationfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeintegrationfilterDud struct { 
    


    


    

}

// Knowledgeintegrationfilter
type Knowledgeintegrationfilter struct { 
    // Name - Filter name, which is the name of a setting in a knowledge source.
    Name string `json:"name"`


    // VarType - Filter type.
    VarType string `json:"type"`


    // Values - Available options of the filter setting.
    Values []Knowledgeintegrationfiltervalue `json:"values"`

}

// String returns a JSON representation of the model
func (o *Knowledgeintegrationfilter) String() string {
    
    
     o.Values = []Knowledgeintegrationfiltervalue{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeintegrationfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeintegrationfilter

    if KnowledgeintegrationfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeintegrationfilterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Values []Knowledgeintegrationfiltervalue `json:"values"`
        *Alias
    }{

        


        


        
        Values: []Knowledgeintegrationfiltervalue{{}},
        

        Alias: (*Alias)(u),
    })
}

