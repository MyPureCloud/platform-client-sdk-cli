package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContentfilerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContentfilerequestDud struct { 
    


    

}

// Contentfilerequest
type Contentfilerequest struct { 
    // UploadKey - Key that identifies the file in the storage including the file name
    UploadKey string `json:"uploadKey"`


    // Name - The name of the file
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Contentfilerequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contentfilerequest) MarshalJSON() ([]byte, error) {
    type Alias Contentfilerequest

    if ContentfilerequestMarshalled {
        return []byte("{}"), nil
    }
    ContentfilerequestMarshalled = true

    return json.Marshal(&struct {
        
        UploadKey string `json:"uploadKey"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

