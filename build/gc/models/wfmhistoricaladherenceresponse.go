package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmhistoricaladherenceresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmhistoricaladherenceresponseDud struct { 
    


    


    


    


    

}

// Wfmhistoricaladherenceresponse
type Wfmhistoricaladherenceresponse struct { 
    // Id - The query ID to listen for
    Id string `json:"id"`


    // DownloadUrl - Deprecated. Use downloadUrls instead.
    DownloadUrl string `json:"downloadUrl"`


    // DownloadResult - Result will always come via downloadUrls; however the schema is included for documentation
    DownloadResult Wfmhistoricaladherenceresultwrapper `json:"downloadResult"`


    // DownloadUrls - The uri list to GET the results of the Historical Adherence query. For notification purposes only
    DownloadUrls []string `json:"downloadUrls"`


    // QueryState - The state of the adherence query
    QueryState string `json:"queryState"`

}

// String returns a JSON representation of the model
func (o *Wfmhistoricaladherenceresponse) String() string {
    
    
    
     o.DownloadUrls = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmhistoricaladherenceresponse) MarshalJSON() ([]byte, error) {
    type Alias Wfmhistoricaladherenceresponse

    if WfmhistoricaladherenceresponseMarshalled {
        return []byte("{}"), nil
    }
    WfmhistoricaladherenceresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        DownloadUrl string `json:"downloadUrl"`
        
        DownloadResult Wfmhistoricaladherenceresultwrapper `json:"downloadResult"`
        
        DownloadUrls []string `json:"downloadUrls"`
        
        QueryState string `json:"queryState"`
        *Alias
    }{

        


        


        


        
        DownloadUrls: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

