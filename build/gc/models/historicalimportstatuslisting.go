package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportstatuslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportstatuslistingDud struct { 
    

}

// Historicalimportstatuslisting
type Historicalimportstatuslisting struct { 
    // Entities
    Entities []Historicalimportstatus `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Historicalimportstatuslisting) String() string {
    
    
     o.Entities = []Historicalimportstatus{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportstatuslisting) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportstatuslisting

    if HistoricalimportstatuslistingMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportstatuslistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Historicalimportstatus `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Historicalimportstatus{{}},
        

        
        Alias: (*Alias)(u),
    })
}

