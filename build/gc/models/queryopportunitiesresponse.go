package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryopportunitiesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryopportunitiesresponseDud struct { 
    


    

}

// Queryopportunitiesresponse
type Queryopportunitiesresponse struct { 
    // Result - The query result. Null if downloadUrl is populated
    Result Opportunitiesresultwithpagination `json:"result"`


    // DownloadUrl - The URL used to retrieve large datasets. If present, the response conforms to the schema for the result field
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Queryopportunitiesresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryopportunitiesresponse) MarshalJSON() ([]byte, error) {
    type Alias Queryopportunitiesresponse

    if QueryopportunitiesresponseMarshalled {
        return []byte("{}"), nil
    }
    QueryopportunitiesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Result Opportunitiesresultwithpagination `json:"result"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

