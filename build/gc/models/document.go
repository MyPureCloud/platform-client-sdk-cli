package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Document
type Document struct { 
    


    // Name
    Name string `json:"name"`


    // ChangeNumber
    ChangeNumber int `json:"changeNumber"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateUploaded - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateUploaded time.Time `json:"dateUploaded"`


    // ContentUri
    ContentUri string `json:"contentUri"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // CreatedBy
    CreatedBy Domainentityref `json:"createdBy"`


    // UploadedBy
    UploadedBy Domainentityref `json:"uploadedBy"`


    // SharingUri
    SharingUri string `json:"sharingUri"`


    // ContentType
    ContentType string `json:"contentType"`


    // ContentLength
    ContentLength int `json:"contentLength"`


    // SystemType
    SystemType string `json:"systemType"`


    // Filename
    Filename string `json:"filename"`


    // PageCount
    PageCount int `json:"pageCount"`


    // Read
    Read bool `json:"read"`


    // CallerAddress
    CallerAddress string `json:"callerAddress"`


    // ReceiverAddress
    ReceiverAddress string `json:"receiverAddress"`


    // Tags
    Tags []string `json:"tags"`


    // TagValues
    TagValues []Tagvalue `json:"tagValues"`


    // Attributes
    Attributes []Documentattribute `json:"attributes"`


    // Thumbnails
    Thumbnails []Documentthumbnail `json:"thumbnails"`


    // UploadStatus
    UploadStatus Domainentityref `json:"uploadStatus"`


    // UploadDestinationUri
    UploadDestinationUri string `json:"uploadDestinationUri"`


    // UploadMethod
    UploadMethod string `json:"uploadMethod"`


    // LockInfo
    LockInfo Lockinfo `json:"lockInfo"`


    // Acl - A list of permitted action rights for the user making the request
    Acl []string `json:"acl"`


    // SharingStatus
    SharingStatus string `json:"sharingStatus"`


    // DownloadSharingUri
    DownloadSharingUri string `json:"downloadSharingUri"`


    

}

// String returns a JSON representation of the model
func (o *Document) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Tags = []string{""} 
     o.TagValues = []Tagvalue{{}} 
     o.Attributes = []Documentattribute{{}} 
     o.Thumbnails = []Documentthumbnail{{}} 
    
    
    
    
     o.Acl = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Document) MarshalJSON() ([]byte, error) {
    type Alias Document

    if DocumentMarshalled {
        return []byte("{}"), nil
    }
    DocumentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ChangeNumber int `json:"changeNumber"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateUploaded time.Time `json:"dateUploaded"`
        
        ContentUri string `json:"contentUri"`
        
        Workspace Domainentityref `json:"workspace"`
        
        CreatedBy Domainentityref `json:"createdBy"`
        
        UploadedBy Domainentityref `json:"uploadedBy"`
        
        SharingUri string `json:"sharingUri"`
        
        ContentType string `json:"contentType"`
        
        ContentLength int `json:"contentLength"`
        
        SystemType string `json:"systemType"`
        
        Filename string `json:"filename"`
        
        PageCount int `json:"pageCount"`
        
        Read bool `json:"read"`
        
        CallerAddress string `json:"callerAddress"`
        
        ReceiverAddress string `json:"receiverAddress"`
        
        Tags []string `json:"tags"`
        
        TagValues []Tagvalue `json:"tagValues"`
        
        Attributes []Documentattribute `json:"attributes"`
        
        Thumbnails []Documentthumbnail `json:"thumbnails"`
        
        UploadStatus Domainentityref `json:"uploadStatus"`
        
        UploadDestinationUri string `json:"uploadDestinationUri"`
        
        UploadMethod string `json:"uploadMethod"`
        
        LockInfo Lockinfo `json:"lockInfo"`
        
        Acl []string `json:"acl"`
        
        SharingStatus string `json:"sharingStatus"`
        
        DownloadSharingUri string `json:"downloadSharingUri"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Tags: []string{""},
        


        
        TagValues: []Tagvalue{{}},
        


        
        Attributes: []Documentattribute{{}},
        


        
        Thumbnails: []Documentthumbnail{{}},
        


        


        


        


        


        
        Acl: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

