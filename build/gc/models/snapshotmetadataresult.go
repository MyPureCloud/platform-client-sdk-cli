package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotmetadataresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotmetadataresultDud struct { 
    


    

}

// Snapshotmetadataresult
type Snapshotmetadataresult struct { 
    // SnapshotInfo - Information about the snapshot
    SnapshotInfo Snapshotinfo `json:"snapshotInfo"`


    // DateForecastStart - Start date of the forecast. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateForecastStart time.Time `json:"dateForecastStart"`

}

// String returns a JSON representation of the model
func (o *Snapshotmetadataresult) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshotmetadataresult) MarshalJSON() ([]byte, error) {
    type Alias Snapshotmetadataresult

    if SnapshotmetadataresultMarshalled {
        return []byte("{}"), nil
    }
    SnapshotmetadataresultMarshalled = true

    return json.Marshal(&struct {
        
        SnapshotInfo Snapshotinfo `json:"snapshotInfo"`
        
        DateForecastStart time.Time `json:"dateForecastStart"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

