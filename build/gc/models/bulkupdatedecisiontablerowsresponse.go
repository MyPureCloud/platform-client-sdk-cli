package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdatedecisiontablerowsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdatedecisiontablerowsresponseDud struct { 
    


    

}

// Bulkupdatedecisiontablerowsresponse
type Bulkupdatedecisiontablerowsresponse struct { 
    // TotalUpdated - The total number of rows successfully updated
    TotalUpdated int `json:"totalUpdated"`


    // Rows - The list of updated decision table rows
    Rows []Decisiontablerow `json:"rows"`

}

// String returns a JSON representation of the model
func (o *Bulkupdatedecisiontablerowsresponse) String() string {
    
     o.Rows = []Decisiontablerow{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdatedecisiontablerowsresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdatedecisiontablerowsresponse

    if BulkupdatedecisiontablerowsresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkupdatedecisiontablerowsresponseMarshalled = true

    return json.Marshal(&struct {
        
        TotalUpdated int `json:"totalUpdated"`
        
        Rows []Decisiontablerow `json:"rows"`
        *Alias
    }{

        


        
        Rows: []Decisiontablerow{{}},
        

        Alias: (*Alias)(u),
    })
}

