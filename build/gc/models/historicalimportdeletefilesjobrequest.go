package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportdeletefilesjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportdeletefilesjobrequestDud struct { 
    

}

// Historicalimportdeletefilesjobrequest
type Historicalimportdeletefilesjobrequest struct { 
    // RequestIds - List of requestIds to be deleted. Max number of RequestIds should be 100
    RequestIds []string `json:"requestIds"`

}

// String returns a JSON representation of the model
func (o *Historicalimportdeletefilesjobrequest) String() string {
     o.RequestIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportdeletefilesjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportdeletefilesjobrequest

    if HistoricalimportdeletefilesjobrequestMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportdeletefilesjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        RequestIds []string `json:"requestIds"`
        *Alias
    }{

        
        RequestIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

