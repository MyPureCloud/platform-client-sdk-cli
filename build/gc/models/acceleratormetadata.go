package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcceleratormetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcceleratormetadataDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Description string `json:"description"`


    Origin string `json:"origin"`


    VarType string `json:"type"`


    Classification string `json:"classification"`


    Tags []string `json:"tags"`


    SelfUri string `json:"selfUri"`

}

// Acceleratormetadata - Metadata for a CX infrastructure as code accelerator
type Acceleratormetadata struct { 
    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Acceleratormetadata) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acceleratormetadata) MarshalJSON() ([]byte, error) {
    type Alias Acceleratormetadata

    if AcceleratormetadataMarshalled {
        return []byte("{}"), nil
    }
    AcceleratormetadataMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

