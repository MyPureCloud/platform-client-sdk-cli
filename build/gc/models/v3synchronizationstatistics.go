package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationstatisticsDud struct { 
    


    

}

// V3synchronizationstatistics
type V3synchronizationstatistics struct { 
    // SynchronizedItemCount - The total number of items added, updated or removed in the synchronization.
    SynchronizedItemCount int `json:"synchronizedItemCount"`


    // FailedItemCount - The number of source items that failed to synchronize.
    FailedItemCount int `json:"failedItemCount"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationstatistics) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationstatistics) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationstatistics

    if V3synchronizationstatisticsMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        SynchronizedItemCount int `json:"synchronizedItemCount"`
        
        FailedItemCount int `json:"failedItemCount"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

