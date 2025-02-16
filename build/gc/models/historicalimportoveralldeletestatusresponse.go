package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportoveralldeletestatusresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportoveralldeletestatusresponseDud struct { 
    


    Status string `json:"status"`

}

// Historicalimportoveralldeletestatusresponse
type Historicalimportoveralldeletestatusresponse struct { 
    // Entities - List of all the delete jobs
    Entities []Historicaldatajobentitystatus `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Historicalimportoveralldeletestatusresponse) String() string {
     o.Entities = []Historicaldatajobentitystatus{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportoveralldeletestatusresponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportoveralldeletestatusresponse

    if HistoricalimportoveralldeletestatusresponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportoveralldeletestatusresponseMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Historicaldatajobentitystatus `json:"entities"`
        *Alias
    }{

        
        Entities: []Historicaldatajobentitystatus{{}},
        


        

        Alias: (*Alias)(u),
    })
}

