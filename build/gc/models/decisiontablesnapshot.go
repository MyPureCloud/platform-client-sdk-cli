package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablesnapshotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablesnapshotDud struct { 
    


    


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Addressableentityref `json:"createdBy"`

}

// Decisiontablesnapshot
type Decisiontablesnapshot struct { 
    // Name - Snapshot name
    Name string `json:"name"`


    // Notes - Optional snapshot notes
    Notes string `json:"notes"`


    


    

}

// String returns a JSON representation of the model
func (o *Decisiontablesnapshot) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablesnapshot) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablesnapshot

    if DecisiontablesnapshotMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablesnapshotMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

