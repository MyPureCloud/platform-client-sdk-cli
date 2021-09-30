package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FaxdocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FaxdocumentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Faxdocument
type Faxdocument struct { 
    


    // Name
    Name string `json:"name"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ContentUri
    ContentUri string `json:"contentUri"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // CreatedBy
    CreatedBy Domainentityref `json:"createdBy"`


    // SharingUri
    SharingUri string `json:"sharingUri"`


    // ContentType
    ContentType string `json:"contentType"`


    // ContentLength
    ContentLength int `json:"contentLength"`


    // Filename
    Filename string `json:"filename"`


    // Read
    Read bool `json:"read"`


    // PageCount
    PageCount int `json:"pageCount"`


    // CallerAddress
    CallerAddress string `json:"callerAddress"`


    // ReceiverAddress
    ReceiverAddress string `json:"receiverAddress"`


    // Thumbnails
    Thumbnails []Documentthumbnail `json:"thumbnails"`


    // DownloadSharingUri
    DownloadSharingUri string `json:"downloadSharingUri"`


    

}

// String returns a JSON representation of the model
func (o *Faxdocument) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Thumbnails = []Documentthumbnail{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Faxdocument) MarshalJSON() ([]byte, error) {
    type Alias Faxdocument

    if FaxdocumentMarshalled {
        return []byte("{}"), nil
    }
    FaxdocumentMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ContentUri string `json:"contentUri"`
        
        Workspace Domainentityref `json:"workspace"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        SharingUri string `json:"sharingUri"`
        
        ContentType string `json:"contentType"`
        
        ContentLength int `json:"contentLength"`
        
        Filename string `json:"filename"`
        
        Read bool `json:"read"`
        
        PageCount int `json:"pageCount"`
        
        CallerAddress string `json:"callerAddress"`
        
        ReceiverAddress string `json:"receiverAddress"`
        
        Thumbnails []Documentthumbnail `json:"thumbnails"`
        
        DownloadSharingUri string `json:"downloadSharingUri"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Thumbnails: []Documentthumbnail{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

