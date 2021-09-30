package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnknowledgesearchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnknowledgesearchDud struct { 
    


    


    

}

// Reportingturnknowledgesearch
type Reportingturnknowledgesearch struct { 
    // SearchId - The ID of this knowledge search.
    SearchId string `json:"searchId"`


    // Documents - The list of search documents captured during this reporting turn.
    Documents []Reportingturnknowledgedocument `json:"documents"`


    // Query - The search query that was used to search the Knowledge Base documents for a matching question.
    Query string `json:"query"`

}

// String returns a JSON representation of the model
func (o *Reportingturnknowledgesearch) String() string {
    
    
    
    
    
    
     o.Documents = []Reportingturnknowledgedocument{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturnknowledgesearch) MarshalJSON() ([]byte, error) {
    type Alias Reportingturnknowledgesearch

    if ReportingturnknowledgesearchMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnknowledgesearchMarshalled = true

    return json.Marshal(&struct { 
        SearchId string `json:"searchId"`
        
        Documents []Reportingturnknowledgedocument `json:"documents"`
        
        Query string `json:"query"`
        
        *Alias
    }{
        

        

        

        
        Documents: []Reportingturnknowledgedocument{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

