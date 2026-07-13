package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationfiledeletionrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationfiledeletionrequestDud struct { 
    


    

}

// V3synchronizationfiledeletionrequest
type V3synchronizationfiledeletionrequest struct { 
    // FileId - The identifier of the file to mark for deletion. Mutually exclusive with fileName.
    FileId string `json:"fileId"`


    // FileName - Name of the file to mark for deletion. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \\{^}%`]\">[~<#|. Mutually exclusive with fileId.
    FileName string `json:"fileName"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationfiledeletionrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationfiledeletionrequest) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationfiledeletionrequest

    if V3synchronizationfiledeletionrequestMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationfiledeletionrequestMarshalled = true

    return json.Marshal(&struct {
        
        FileId string `json:"fileId"`
        
        FileName string `json:"fileName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

