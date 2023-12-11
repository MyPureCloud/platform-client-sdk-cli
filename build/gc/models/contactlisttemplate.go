package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlisttemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlisttemplateDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contactlisttemplate
type Contactlisttemplate struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // ColumnNames - The names of the contact data columns.
    ColumnNames []string `json:"columnNames"`


    // PhoneColumns - Indicates which columns are phone numbers.
    PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`


    // EmailColumns - Indicates which columns are email addresses
    EmailColumns []Emailcolumn `json:"emailColumns"`


    // PreviewModeColumnName - A column to check if a contact should always be dialed in preview mode.
    PreviewModeColumnName string `json:"previewModeColumnName"`


    // PreviewModeAcceptedValues - The values in the previewModeColumnName column that indicate a contact should always be dialed in preview mode.
    PreviewModeAcceptedValues []string `json:"previewModeAcceptedValues"`


    // AttemptLimits - AttemptLimits for this ContactListTemplate.
    AttemptLimits Domainentityref `json:"attemptLimits"`


    // AutomaticTimeZoneMapping - Indicates if automatic time zone mapping is to be used for this ContactListTemplate.
    AutomaticTimeZoneMapping bool `json:"automaticTimeZoneMapping"`


    // ZipCodeColumnName - The name of ContactListTemplate column containing the zip code for use with automatic time zone mapping. Only allowed if 'automaticTimeZoneMapping' is set to true.
    ZipCodeColumnName string `json:"zipCodeColumnName"`


    // ColumnDataTypeSpecifications - The settings of the columns selected for dynamic queueing
    ColumnDataTypeSpecifications []Columndatatypespecification `json:"columnDataTypeSpecifications"`


    // TrimWhitespace - Whether to trim white space when importing a ContactListTemplate csv file, default value = true
    TrimWhitespace bool `json:"trimWhitespace"`


    

}

// String returns a JSON representation of the model
func (o *Contactlisttemplate) String() string {
    
    
     o.ColumnNames = []string{""} 
     o.PhoneColumns = []Contactphonenumbercolumn{{}} 
     o.EmailColumns = []Emailcolumn{{}} 
    
     o.PreviewModeAcceptedValues = []string{""} 
    
    
    
     o.ColumnDataTypeSpecifications = []Columndatatypespecification{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlisttemplate) MarshalJSON() ([]byte, error) {
    type Alias Contactlisttemplate

    if ContactlisttemplateMarshalled {
        return []byte("{}"), nil
    }
    ContactlisttemplateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        ColumnNames []string `json:"columnNames"`
        
        PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`
        
        EmailColumns []Emailcolumn `json:"emailColumns"`
        
        PreviewModeColumnName string `json:"previewModeColumnName"`
        
        PreviewModeAcceptedValues []string `json:"previewModeAcceptedValues"`
        
        AttemptLimits Domainentityref `json:"attemptLimits"`
        
        AutomaticTimeZoneMapping bool `json:"automaticTimeZoneMapping"`
        
        ZipCodeColumnName string `json:"zipCodeColumnName"`
        
        ColumnDataTypeSpecifications []Columndatatypespecification `json:"columnDataTypeSpecifications"`
        
        TrimWhitespace bool `json:"trimWhitespace"`
        *Alias
    }{

        


        


        


        


        


        
        ColumnNames: []string{""},
        


        
        PhoneColumns: []Contactphonenumbercolumn{{}},
        


        
        EmailColumns: []Emailcolumn{{}},
        


        


        
        PreviewModeAcceptedValues: []string{""},
        


        


        


        


        
        ColumnDataTypeSpecifications: []Columndatatypespecification{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

