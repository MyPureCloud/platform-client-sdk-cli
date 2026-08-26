package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClusterscanlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClusterscanlistDud struct { 
    


    


    


    


    

}

// Clusterscanlist
type Clusterscanlist struct { 
    // Entities
    Entities []Clusterscan `json:"entities"`


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
func (o *Clusterscanlist) String() string {
     o.Entities = []Clusterscan{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clusterscanlist) MarshalJSON() ([]byte, error) {
    type Alias Clusterscanlist

    if ClusterscanlistMarshalled {
        return []byte("{}"), nil
    }
    ClusterscanlistMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Clusterscan `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        Cursors Cursors `json:"cursors"`
        *Alias
    }{

        
        Entities: []Clusterscan{{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

