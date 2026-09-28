package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateagentunavailabletimesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateagentunavailabletimesresponseDud struct { 
    


    

}

// Bulkupdateagentunavailabletimesresponse
type Bulkupdateagentunavailabletimesresponse struct { 
    // Results - The result of each unavailable time operation, in the order the operations were requested
    Results []Bulkupdateagentunavailabletimesresultitem `json:"results"`


    // VarError - The error that stopped processing, populated when one or more operations failed
    VarError Errorbody `json:"error"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateagentunavailabletimesresponse) String() string {
     o.Results = []Bulkupdateagentunavailabletimesresultitem{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateagentunavailabletimesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateagentunavailabletimesresponse

    if BulkupdateagentunavailabletimesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateagentunavailabletimesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Bulkupdateagentunavailabletimesresultitem `json:"results"`
        
        VarError Errorbody `json:"error"`
        *Alias
    }{

        
        Results: []Bulkupdateagentunavailabletimesresultitem{{}},
        


        

        Alias: (*Alias)(u),
    })
}

