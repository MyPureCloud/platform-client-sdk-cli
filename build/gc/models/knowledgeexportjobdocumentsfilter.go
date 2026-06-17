package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeexportjobdocumentsfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeexportjobdocumentsfilterDud struct { 
    


    


    

}

// Knowledgeexportjobdocumentsfilter
type Knowledgeexportjobdocumentsfilter struct { 
    // Interval - Retrieves the documents modified in specified date and time range. Cannot be used together with entities filter. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Interval string `json:"interval"`


    // Entities - Retrieves the documents with the given ids. Cannot be used together with internal filter.
    Entities []Entity `json:"entities"`


    // SourceId
    SourceId string `json:"sourceId"`

}

// String returns a JSON representation of the model
func (o *Knowledgeexportjobdocumentsfilter) String() string {
    
     o.Entities = []Entity{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeexportjobdocumentsfilter) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeexportjobdocumentsfilter

    if KnowledgeexportjobdocumentsfilterMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeexportjobdocumentsfilterMarshalled = true

    return json.Marshal(&struct {
        
        Interval string `json:"interval"`
        
        Entities []Entity `json:"entities"`
        
        SourceId string `json:"sourceId"`
        *Alias
    }{

        


        
        Entities: []Entity{{}},
        


        

        Alias: (*Alias)(u),
    })
}

