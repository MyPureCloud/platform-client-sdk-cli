package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterlistDud struct { 
    


    


    


    


    

}

// Clusterlist
type Clusterlist struct { 
    // Entities
    Entities []Cluster `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // Cursors - The cursor that points to the next set of entities being returned.
    Cursors Cursors `json:"cursors"`

}

// String returns a JSON representation of the model
func (o *Clusterlist) String() string {
     o.Entities = []Cluster{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clusterlist) MarshalJSON() ([]byte, error) {
    type Alias Clusterlist

    if ClusterlistMarshalled {
        return []byte("{}"), nil
    }
    ClusterlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Cluster `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Cluster{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

