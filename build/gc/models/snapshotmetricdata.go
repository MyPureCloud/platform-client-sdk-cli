package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnapshotmetricdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnapshotmetricdataDud struct { 
    


    

}

// Snapshotmetricdata
type Snapshotmetricdata struct { 
    // Weekly - Weekly time series for snapshot data
    Weekly Weekly `json:"weekly"`


    // QuarterHour - Quarter hour time series for snapshot data
    QuarterHour Quarterhourly `json:"quarterHour"`

}

// String returns a JSON representation of the model
func (o *Snapshotmetricdata) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snapshotmetricdata) MarshalJSON() ([]byte, error) {
    type Alias Snapshotmetricdata

    if SnapshotmetricdataMarshalled {
        return []byte("{}"), nil
    }
    SnapshotmetricdataMarshalled = true

    return json.Marshal(&struct {
        
        Weekly Weekly `json:"weekly"`
        
        QuarterHour Quarterhourly `json:"quarterHour"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

