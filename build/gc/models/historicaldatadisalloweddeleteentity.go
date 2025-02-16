package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaldatadisalloweddeleteentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaldatadisalloweddeleteentityDud struct { 
    


    

}

// Historicaldatadisalloweddeleteentity
type Historicaldatadisalloweddeleteentity struct { 
    // RequestId - The requestId associated with this disallowed entity
    RequestId string `json:"requestId"`


    // Reason - The error code explaining why the delete request for the requestId was disallowed
    Reason string `json:"reason"`

}

// String returns a JSON representation of the model
func (o *Historicaldatadisalloweddeleteentity) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaldatadisalloweddeleteentity) MarshalJSON() ([]byte, error) {
    type Alias Historicaldatadisalloweddeleteentity

    if HistoricaldatadisalloweddeleteentityMarshalled {
        return []byte("{}"), nil
    }
    HistoricaldatadisalloweddeleteentityMarshalled = true

    return json.Marshal(&struct {
        
        RequestId string `json:"requestId"`
        
        Reason string `json:"reason"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

