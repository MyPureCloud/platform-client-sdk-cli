package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkupdateagentunavailabletimesresultitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkupdateagentunavailabletimesresultitemDud struct { 
    


    

}

// Bulkupdateagentunavailabletimesresultitem
type Bulkupdateagentunavailabletimesresultitem struct { 
    // UnavailableTime - The unavailable time that was created, updated, or deleted. Populated when the operation completed successfully
    UnavailableTime Targetunavailabletime `json:"unavailableTime"`


    // Status - The status of the operation
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Bulkupdateagentunavailabletimesresultitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkupdateagentunavailabletimesresultitem) MarshalJSON() ([]byte, error) {
    type Alias Bulkupdateagentunavailabletimesresultitem

    if BulkupdateagentunavailabletimesresultitemMarshalled {
        return []byte("{}"), nil
    }
    BulkupdateagentunavailabletimesresultitemMarshalled = true

    return json.Marshal(&struct {
        
        UnavailableTime Targetunavailabletime `json:"unavailableTime"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

