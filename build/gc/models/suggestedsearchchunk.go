package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestedsearchchunkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestedsearchchunkDud struct { 
    Title string `json:"title"`


    Uri string `json:"uri"`


    ChunkId string `json:"chunkId"`


    Text string `json:"text"`


    Confidence float32 `json:"confidence"`


    Document Addressableentityref `json:"document"`

}

// Suggestedsearchchunk
type Suggestedsearchchunk struct { 
    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Suggestedsearchchunk) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestedsearchchunk) MarshalJSON() ([]byte, error) {
    type Alias Suggestedsearchchunk

    if SuggestedsearchchunkMarshalled {
        return []byte("{}"), nil
    }
    SuggestedsearchchunkMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

