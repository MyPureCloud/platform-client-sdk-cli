package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportstatusDud struct { 
    RequestId string `json:"requestId"`


    DateImportEnded time.Time `json:"dateImportEnded"`


    DateImportStarted time.Time `json:"dateImportStarted"`


    Status string `json:"status"`


    VarError string `json:"error"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Active bool `json:"active"`


    VarType string `json:"type"`


    FileName string `json:"fileName"`


    FileSize int `json:"fileSize"`

}

// Historicalimportstatus
type Historicalimportstatus struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Historicalimportstatus) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportstatus) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportstatus

    if HistoricalimportstatusMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportstatusMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

