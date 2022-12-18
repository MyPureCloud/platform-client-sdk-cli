package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnnotationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnnotationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Annotation
type Annotation struct { 
    


    // Name
    Name string `json:"name"`


    // VarType
    VarType string `json:"type"`


    // Location - Offset of annotation in milliseconds.
    Location int `json:"location"`


    // DurationMs - Duration of annotation in milliseconds.
    DurationMs int `json:"durationMs"`


    // AbsoluteLocation - Offset of annotation (milliseconds) from start of recording (after removing the cumulative duration of all pauses).
    AbsoluteLocation int `json:"absoluteLocation"`


    // AbsoluteDurationMs - Duration of annotation (milliseconds).
    AbsoluteDurationMs int `json:"absoluteDurationMs"`


    // RecordingLocation - Offset of annotation (milliseconds) from start of recording, adjusted for any recording cuts
    RecordingLocation int `json:"recordingLocation"`


    // RecordingDurationMs - Duration of annotation (milliseconds), adjusted for any recording cuts.
    RecordingDurationMs int `json:"recordingDurationMs"`


    // User - User that created this annotation (if any).
    User User `json:"user"`


    // Description - Text of annotation. Maximum character limit is 500.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Annotation) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Annotation) MarshalJSON() ([]byte, error) {
    type Alias Annotation

    if AnnotationMarshalled {
        return []byte("{}"), nil
    }
    AnnotationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Location int `json:"location"`
        
        DurationMs int `json:"durationMs"`
        
        AbsoluteLocation int `json:"absoluteLocation"`
        
        AbsoluteDurationMs int `json:"absoluteDurationMs"`
        
        RecordingLocation int `json:"recordingLocation"`
        
        RecordingDurationMs int `json:"recordingDurationMs"`
        
        User User `json:"user"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

