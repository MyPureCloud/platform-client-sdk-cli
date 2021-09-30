package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SipdownloadresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SipdownloadresponseDud struct { 
    


    

}

// Sipdownloadresponse
type Sipdownloadresponse struct { 
    // DownloadId - unique id of the downloaded file
    DownloadId string `json:"downloadId"`


    // DocumentId - Document id of pcap file
    DocumentId string `json:"documentId"`

}

// String returns a JSON representation of the model
func (o *Sipdownloadresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sipdownloadresponse) MarshalJSON() ([]byte, error) {
    type Alias Sipdownloadresponse

    if SipdownloadresponseMarshalled {
        return []byte("{}"), nil
    }
    SipdownloadresponseMarshalled = true

    return json.Marshal(&struct { 
        DownloadId string `json:"downloadId"`
        
        DocumentId string `json:"documentId"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

