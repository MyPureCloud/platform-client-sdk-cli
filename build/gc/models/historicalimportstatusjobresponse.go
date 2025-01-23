package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportstatusjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportstatusjobresponseDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Historicalimportstatusjobresponse
type Historicalimportstatusjobresponse struct { 
    


    // ImportStatusResult - The historical import status result of the import job
    ImportStatusResult Historicalimportstatus `json:"importStatusResult"`


    

}

// String returns a JSON representation of the model
func (o *Historicalimportstatusjobresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportstatusjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportstatusjobresponse

    if HistoricalimportstatusjobresponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportstatusjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        ImportStatusResult Historicalimportstatus `json:"importStatusResult"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

