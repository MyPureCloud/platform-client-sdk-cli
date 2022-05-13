package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgedocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgedocumentDud struct { 
    


    


    


    

}

// Reportingturnknowledgedocument
type Reportingturnknowledgedocument struct { 
    // Id - The ID of the knowledge document.
    Id string `json:"id"`


    // Question - The the question that was used to match against the search query.
    Question string `json:"question"`


    // Answer - The corresponding answer to the question.
    Answer string `json:"answer"`


    // Confidence - The confidence score of how well the question matched the search query.
    Confidence float64 `json:"confidence"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgedocument) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgedocument) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgedocument

    if ReportingturnknowledgedocumentMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgedocumentMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Question string `json:"question"`
        
        Answer string `json:"answer"`
        
        Confidence float64 `json:"confidence"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

