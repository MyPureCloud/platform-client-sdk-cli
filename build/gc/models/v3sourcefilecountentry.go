package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    V3sourcefilecountentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type V3sourcefilecountentryDud struct { 
    


    


    


    

}

// V3sourcefilecountentry
type V3sourcefilecountentry struct { 
    // Source - Knowledge source reference for this entry.
    Source V3sourceref `json:"source"`


    // SelectedTagsFileCount - Unique file count matching the tag filter for this source.
    SelectedTagsFileCount int `json:"selectedTagsFileCount"`


    // TotalFileCount - Total unique file count for this source without tag filters.
    TotalFileCount int `json:"totalFileCount"`


    // FileNames - File names when includeFileNames is enabled.
    FileNames []string `json:"fileNames"`

}

// String returns a JSON representation of the model
func (o *V3sourcefilecountentry) String() string {
    
    
    
     o.FileNames = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *V3sourcefilecountentry) MarshalJSON() ([]byte, error) {
    type Alias V3sourcefilecountentry

    if V3sourcefilecountentryMarshalled {
        return []byte("{}"), nil
    }
    V3sourcefilecountentryMarshalled = true

    return json.Marshal(&struct {
        
        Source V3sourceref `json:"source"`
        
        SelectedTagsFileCount int `json:"selectedTagsFileCount"`
        
        TotalFileCount int `json:"totalFileCount"`
        
        FileNames []string `json:"fileNames"`
        *Alias
    }{

        


        


        


        
        FileNames: []string{""},
        

        Alias: (*Alias)(u),
    })
}

