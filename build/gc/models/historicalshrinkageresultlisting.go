package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalshrinkageresultlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalshrinkageresultlistingDud struct { 
    

}

// Historicalshrinkageresultlisting
type Historicalshrinkageresultlisting struct { 
    // Entities
    Entities []Historicalshrinkageresult `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Historicalshrinkageresultlisting) String() string {
     o.Entities = []Historicalshrinkageresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalshrinkageresultlisting) MarshalJSON() ([]byte, error) {
    type Alias Historicalshrinkageresultlisting

    if HistoricalshrinkageresultlistingMarshalled {
        return []byte("{}"), nil
    }
    HistoricalshrinkageresultlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Historicalshrinkageresult `json:"entities"`
        *Alias
    }{

        
        Entities: []Historicalshrinkageresult{{}},
        

        Alias: (*Alias)(u),
    })
}

