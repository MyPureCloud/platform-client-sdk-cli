package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterstatisticsDud struct { 
    TotalClusters int `json:"totalClusters"`


    TotalAutomaticMergeSuccesses int `json:"totalAutomaticMergeSuccesses"`


    TotalAutomaticMergeFailures int `json:"totalAutomaticMergeFailures"`

}

// Clusterstatistics
type Clusterstatistics struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Clusterstatistics) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clusterstatistics) MarshalJSON() ([]byte, error) {
    type Alias Clusterstatistics

    if ClusterstatisticsMarshalled {
        return []byte("{}"), nil
    }
    ClusterstatisticsMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

