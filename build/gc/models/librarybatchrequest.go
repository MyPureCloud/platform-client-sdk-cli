package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LibrarybatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LibrarybatchrequestDud struct { 
    

}

// Librarybatchrequest
type Librarybatchrequest struct { 
    // LibraryIds - List of Library IDs
    LibraryIds []string `json:"libraryIds"`

}

// String returns a JSON representation of the model
func (o *Librarybatchrequest) String() string {
     o.LibraryIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Librarybatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Librarybatchrequest

    if LibrarybatchrequestMarshalled {
        return []byte("{}"), nil
    }
    LibrarybatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        LibraryIds []string `json:"libraryIds"`
        *Alias
    }{

        
        LibraryIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

