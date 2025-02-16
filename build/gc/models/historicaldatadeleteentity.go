package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaldatadeleteentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaldatadeleteentityDud struct { 
    


    

}

// Historicaldatadeleteentity
type Historicaldatadeleteentity struct { 
    // RequestId
    RequestId string `json:"requestId"`


    // Status
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Historicaldatadeleteentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaldatadeleteentity) MarshalJSON() ([]byte, error) {
    type Alias Historicaldatadeleteentity

    if HistoricaldatadeleteentityMarshalled {
        return []byte("{}"), nil
    }
    HistoricaldatadeleteentityMarshalled = true

    return json.Marshal(&struct {
        
        RequestId string `json:"requestId"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

