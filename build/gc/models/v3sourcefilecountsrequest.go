package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefilecountsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefilecountsrequestDud struct { 
    


    


    

}

// V3sourcefilecountsrequest
type V3sourcefilecountsrequest struct { 
    // SourceIds - Source IDs used to scope the file count response.
    SourceIds []string `json:"sourceIds"`


    // Filter - Composite tag filter.
    Filter V3sourcetagfilter `json:"filter"`


    // IncludeFileNames - When true, response may include fileNames per source.
    IncludeFileNames bool `json:"includeFileNames"`

}

// String returns a JSON representation of the model
func (o *V3sourcefilecountsrequest) String() string {
     o.SourceIds = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefilecountsrequest) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefilecountsrequest

    if V3sourcefilecountsrequestMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefilecountsrequestMarshalled = true

    return json.Marshal(&struct {
        
        SourceIds []string `json:"sourceIds"`
        
        Filter V3sourcetagfilter `json:"filter"`
        
        IncludeFileNames bool `json:"includeFileNames"`
        *Alias
    }{

        
        SourceIds: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

