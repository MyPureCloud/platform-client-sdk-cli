package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterscanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterscanDud struct { 
    Id string `json:"id"`


    DateCompleted time.Time `json:"dateCompleted"`


    Statistics Clusterscanstatistics `json:"statistics"`


    SelfUri string `json:"selfUri"`

}

// Clusterscan
type Clusterscan struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Clusterscan) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clusterscan) MarshalJSON() ([]byte, error) {
    type Alias Clusterscan

    if ClusterscanMarshalled {
        return []byte("{}"), nil
    }
    ClusterscanMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

