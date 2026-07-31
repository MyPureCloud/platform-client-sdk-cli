package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesourceDud struct { 
    


    


    


    

}

// Knowledgesource
type Knowledgesource struct { 
    // SourceName - The name of the knowledge source.
    SourceName string `json:"sourceName"`


    // Text - The retrieved source text content.
    Text string `json:"text"`


    // Url - URL of the source document.
    Url string `json:"url"`


    // Confidence - Confidence score for this knowledge source.
    Confidence float64 `json:"confidence"`

}

// String returns a JSON representation of the model
func (o *Knowledgesource) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesource) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesource

    if KnowledgesourceMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesourceMarshalled = true

    return json.Marshal(&struct {
        
        SourceName string `json:"sourceName"`
        
        Text string `json:"text"`
        
        Url string `json:"url"`
        
        Confidence float64 `json:"confidence"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

