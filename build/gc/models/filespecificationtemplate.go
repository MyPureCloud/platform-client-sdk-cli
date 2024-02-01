package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FilespecificationtemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FilespecificationtemplateDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Filespecificationtemplate
type Filespecificationtemplate struct { 
    


    // Name - The name of the File Specification template.
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Description - Description of the file specification template
    Description string `json:"description"`


    // Format - File format
    Format string `json:"format"`


    // NumberOfHeadingLinesSkipped - Number of heading lines to be skipped
    NumberOfHeadingLinesSkipped int `json:"numberOfHeadingLinesSkipped"`


    // NumberOfTrailingLinesSkipped - Number of trailing lines to be skipped
    NumberOfTrailingLinesSkipped int `json:"numberOfTrailingLinesSkipped"`


    // Header - If true indicates that delimited file has a header row, which can provide column names
    Header bool `json:"header"`


    // Delimiter - Kind of delimiter
    Delimiter string `json:"delimiter"`


    // DelimiterValue - Delimiter character, used only when delimiter=\"Custom\"
    DelimiterValue string `json:"delimiterValue"`


    // ColumnInformation - Columns specification
    ColumnInformation []Column `json:"columnInformation"`


    // PreprocessingRules - Preprocessing rules
    PreprocessingRules []Preprocessingrule `json:"preprocessingRules"`


    

}

// String returns a JSON representation of the model
func (o *Filespecificationtemplate) String() string {
    
    
    
    
    
    
    
    
    
     o.ColumnInformation = []Column{{}} 
     o.PreprocessingRules = []Preprocessingrule{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Filespecificationtemplate) MarshalJSON() ([]byte, error) {
    type Alias Filespecificationtemplate

    if FilespecificationtemplateMarshalled {
        return []byte("{}"), nil
    }
    FilespecificationtemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        Format string `json:"format"`
        
        NumberOfHeadingLinesSkipped int `json:"numberOfHeadingLinesSkipped"`
        
        NumberOfTrailingLinesSkipped int `json:"numberOfTrailingLinesSkipped"`
        
        Header bool `json:"header"`
        
        Delimiter string `json:"delimiter"`
        
        DelimiterValue string `json:"delimiterValue"`
        
        ColumnInformation []Column `json:"columnInformation"`
        
        PreprocessingRules []Preprocessingrule `json:"preprocessingRules"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        ColumnInformation: []Column{{}},
        


        
        PreprocessingRules: []Preprocessingrule{{}},
        


        

        Alias: (*Alias)(u),
    })
}

