package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelogsjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelogsjobrequestDud struct { 
    


    


    

}

// Edgelogsjobrequest
type Edgelogsjobrequest struct { 
    // Path - A relative directory to the root Edge log folder to query from.
    Path string `json:"path"`


    // Query - The pattern to use when searching for logs, which may include the wildcards {*, ?}.  Multiple search patterns may be combined using a pipe '|' as a delimiter.
    Query string `json:"query"`


    // Recurse - Boolean whether or not to recurse into directories.
    Recurse bool `json:"recurse"`

}

// String returns a JSON representation of the model
func (o *Edgelogsjobrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgelogsjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Edgelogsjobrequest

    if EdgelogsjobrequestMarshalled {
        return []byte("{}"), nil
    }
    EdgelogsjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Path string `json:"path"`
        
        Query string `json:"query"`
        
        Recurse bool `json:"recurse"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

