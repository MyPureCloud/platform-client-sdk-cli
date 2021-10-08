package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelogsjobfileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelogsjobfileDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Edgelogsjobfile
type Edgelogsjobfile struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // TimeCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimeCreated time.Time `json:"timeCreated"`


    // TimeModified - The time this log file was last modified on the Edge. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    TimeModified time.Time `json:"timeModified"`


    // SizeBytes - The size of this file in bytes.
    SizeBytes float64 `json:"sizeBytes"`


    // UploadStatus - The status of the upload of this file from the Edge to the cloud.  Use /upload to start an upload.
    UploadStatus string `json:"uploadStatus"`


    // EdgePath - The path of this file on the Edge.
    EdgePath string `json:"edgePath"`


    // DownloadId - The download ID to use with the downloads API.
    DownloadId string `json:"downloadId"`


    

}

// String returns a JSON representation of the model
func (o *Edgelogsjobfile) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgelogsjobfile) MarshalJSON() ([]byte, error) {
    type Alias Edgelogsjobfile

    if EdgelogsjobfileMarshalled {
        return []byte("{}"), nil
    }
    EdgelogsjobfileMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        TimeCreated time.Time `json:"timeCreated"`
        
        TimeModified time.Time `json:"timeModified"`
        
        SizeBytes float64 `json:"sizeBytes"`
        
        UploadStatus string `json:"uploadStatus"`
        
        EdgePath string `json:"edgePath"`
        
        DownloadId string `json:"downloadId"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

