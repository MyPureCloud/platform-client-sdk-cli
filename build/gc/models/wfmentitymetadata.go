package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmentitymetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmentitymetadataDud struct { 
    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`

}

// Wfmentitymetadata
type Wfmentitymetadata struct { 
    


    


    


    

}

// String returns a JSON representation of the model
func (o *Wfmentitymetadata) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmentitymetadata) MarshalJSON() ([]byte, error) {
    type Alias Wfmentitymetadata

    if WfmentitymetadataMarshalled {
        return []byte("{}"), nil
    }
    WfmentitymetadataMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

