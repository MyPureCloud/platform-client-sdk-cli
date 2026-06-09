package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkadddecisiontablerowsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkadddecisiontablerowsresponseDud struct { 
    


    

}

// Bulkadddecisiontablerowsresponse
type Bulkadddecisiontablerowsresponse struct { 
    // TotalCreated - The total number of rows successfully created
    TotalCreated int `json:"totalCreated"`


    // Rows - The list of created decision table rows
    Rows []Decisiontablerow `json:"rows"`

}

// String returns a JSON representation of the model
func (o *Bulkadddecisiontablerowsresponse) String() string {
    
     o.Rows = []Decisiontablerow{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkadddecisiontablerowsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkadddecisiontablerowsresponse

    if BulkadddecisiontablerowsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkadddecisiontablerowsresponseMarshalled = true

    return json.Marshal(&struct {
        
        TotalCreated int `json:"totalCreated"`
        
        Rows []Decisiontablerow `json:"rows"`
        *Alias
    }{

        


        
        Rows: []Decisiontablerow{{}},
        

        Alias: (*Alias)(u),
    })
}

