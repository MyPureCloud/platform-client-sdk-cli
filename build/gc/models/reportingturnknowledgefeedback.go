package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgefeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgefeedbackDud struct { 
    


    


    

}

// Reportingturnknowledgefeedback
type Reportingturnknowledgefeedback struct { 
    // SearchId - The ID of the original knowledge search that this feedback relates to.
    SearchId string `json:"searchId"`


    // Rating - The feedback rating for the search (1.0 - 5.0). 1 = Negative, 5 = Positive.
    Rating int `json:"rating"`


    // Documents - The list of search documents that the feedback applies to.
    Documents []Reportingturnknowledgedocument `json:"documents"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgefeedback) String() string {
    
    
     o.Documents = []Reportingturnknowledgedocument{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgefeedback) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgefeedback

    if ReportingturnknowledgefeedbackMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgefeedbackMarshalled = true

    return json.Marshal(&struct {
        
        SearchId string `json:"searchId"`
        
        Rating int `json:"rating"`
        
        Documents []Reportingturnknowledgedocument `json:"documents"`
        *Alias
    }{

        


        


        
        Documents: []Reportingturnknowledgedocument{{}},
        

        Alias: (*Alias)(u),
    })
}

