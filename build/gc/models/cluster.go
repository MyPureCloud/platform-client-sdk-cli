package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterDud struct { 
    Id string `json:"id"`


    Division Starrabledivision `json:"division"`


    ClusterScan Clusterscan `json:"clusterScan"`


    MergeInfo Mergeinfo `json:"mergeInfo"`


    Graph Graph `json:"graph"`


    DateCreated time.Time `json:"dateCreated"`


    SelfUri string `json:"selfUri"`

}

// Cluster
type Cluster struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Cluster) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cluster) MarshalJSON() ([]byte, error) {
    type Alias Cluster

    if ClusterMarshalled {
        return []byte("{}"), nil
    }
    ClusterMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

