package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationdeletionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationdeletionDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// V3synchronizationdeletion
type V3synchronizationdeletion struct { 
    // FileId - The unique identifier for the deletion object.
    FileId string `json:"fileId"`


    // FileName - Name of the file marked for deletion.
    FileName string `json:"fileName"`


    // Synchronization - The synchronization of the file deletion.
    Synchronization V3synchronizationref `json:"synchronization"`


    

}

// String returns a JSON representation of the model
func (o *V3synchronizationdeletion) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationdeletion) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationdeletion

    if V3synchronizationdeletionMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationdeletionMarshalled = true

    return json.Marshal(&struct {
        
        FileId string `json:"fileId"`
        
        FileName string `json:"fileName"`
        
        Synchronization V3synchronizationref `json:"synchronization"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

