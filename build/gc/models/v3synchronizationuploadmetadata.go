package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3synchronizationuploadmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3synchronizationuploadmetadataDud struct { 
    


    

}

// V3synchronizationuploadmetadata
type V3synchronizationuploadmetadata struct { 
    // OriginUri - The origin URI of the file to upload
    OriginUri string `json:"originUri"`


    // Tags - List of tags that is used for filtering
    Tags []Fabrictag `json:"tags"`

}

// String returns a JSON representation of the model
func (o *V3synchronizationuploadmetadata) String() string {
    
     o.Tags = []Fabrictag{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3synchronizationuploadmetadata) MarshalJSON() ([]byte, error) {
    type Alias V3synchronizationuploadmetadata

    if V3synchronizationuploadmetadataMarshalled {
        return []byte("{}"), nil
    }
    V3synchronizationuploadmetadataMarshalled = true

    return json.Marshal(&struct {
        
        OriginUri string `json:"originUri"`
        
        Tags []Fabrictag `json:"tags"`
        *Alias
    }{

        


        
        Tags: []Fabrictag{{}},
        

        Alias: (*Alias)(u),
    })
}

