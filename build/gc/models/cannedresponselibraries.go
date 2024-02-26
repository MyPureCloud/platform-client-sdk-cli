package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CannedresponselibrariesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CannedresponselibrariesDud struct { 
    


    

}

// Cannedresponselibraries
type Cannedresponselibraries struct { 
    // LibraryIds - Set of canned response library IDs associated with the queue only when mode is SelectedOnly.
    LibraryIds []string `json:"libraryIds"`


    // Mode - The association mode of canned response libraries to queue
    Mode string `json:"mode"`

}

// String returns a JSON representation of the model
func (o *Cannedresponselibraries) String() string {
     o.LibraryIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cannedresponselibraries) MarshalJSON() ([]byte, error) {
    type Alias Cannedresponselibraries

    if CannedresponselibrariesMarshalled {
        return []byte("{}"), nil
    }
    CannedresponselibrariesMarshalled = true

    return json.Marshal(&struct {
        
        LibraryIds []string `json:"libraryIds"`
        
        Mode string `json:"mode"`
        *Alias
    }{

        
        LibraryIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

