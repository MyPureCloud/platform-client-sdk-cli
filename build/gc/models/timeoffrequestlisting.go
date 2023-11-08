package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestlistingDud struct { 
    


    

}

// Timeoffrequestlisting
type Timeoffrequestlisting struct { 
    // Entities - List of time off requests
    Entities []Timeoffrequest `json:"entities"`


    // DownloadUrl - URL from which to fetch results for requests with a large result set. If populated, the downloaded data will conform to the same schema as would normally be returned, excepting downloaded data will never itself contain a downloadUrl
    DownloadUrl string `json:"downloadUrl"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestlisting) String() string {
     o.Entities = []Timeoffrequest{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestlisting) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestlisting

    if TimeoffrequestlistingMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestlistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Timeoffrequest `json:"entities"`
        
        DownloadUrl string `json:"downloadUrl"`
        *Alias
    }{

        
        Entities: []Timeoffrequest{{}},
        


        

        Alias: (*Alias)(u),
    })
}

