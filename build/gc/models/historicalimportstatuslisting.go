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


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

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
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Entities: []Historicalimportstatus{{}},
        


        

        Alias: (*Alias)(u),
    })
}

