package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalmetricdatawriteresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalmetricdatawriteresponseDud struct { 
    


    

}

// Externalmetricdatawriteresponse - External metric data write response
type Externalmetricdatawriteresponse struct { 
    // ProcessedEntities - The list of processed entities
    ProcessedEntities []Externalmetricdataprocesseditem `json:"processedEntities"`


    // UnprocessedEntities - The list of unprocessed entities
    UnprocessedEntities []Externalmetricdataunprocesseditem `json:"unprocessedEntities"`

}

// String returns a JSON representation of the model
func (o *Externalmetricdatawriteresponse) String() string {
     o.ProcessedEntities = []Externalmetricdataprocesseditem{{}} 
     o.UnprocessedEntities = []Externalmetricdataunprocesseditem{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalmetricdatawriteresponse) MarshalJSON() ([]byte, error) {
    type Alias Externalmetricdatawriteresponse

    if ExternalmetricdatawriteresponseMarshalled {
        return []byte("{}"), nil
    }
    ExternalmetricdatawriteresponseMarshalled = true

    return json.Marshal(&struct {
        
        ProcessedEntities []Externalmetricdataprocesseditem `json:"processedEntities"`
        
        UnprocessedEntities []Externalmetricdataunprocesseditem `json:"unprocessedEntities"`
        *Alias
    }{

        
        ProcessedEntities: []Externalmetricdataprocesseditem{{}},
        


        
        UnprocessedEntities: []Externalmetricdataunprocesseditem{{}},
        

        Alias: (*Alias)(u),
    })
}

