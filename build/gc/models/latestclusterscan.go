package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LatestclusterscanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LatestclusterscanDud struct { 
    Id string `json:"id"`


    DateCompleted time.Time `json:"dateCompleted"`


    Statistics Clusterscanstatistics `json:"statistics"`


    Settings Clusterscangraphsettings `json:"settings"`


    SelfUri string `json:"selfUri"`

}

// Latestclusterscan
type Latestclusterscan struct { 
    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Latestclusterscan) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Latestclusterscan) MarshalJSON() ([]byte, error) {
    type Alias Latestclusterscan

    if LatestclusterscanMarshalled {
        return []byte("{}"), nil
    }
    LatestclusterscanMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

