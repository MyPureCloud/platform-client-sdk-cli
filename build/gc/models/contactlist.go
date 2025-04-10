package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactlistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactlistDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    


    


    


    


    


    ImportStatus Importstatus `json:"importStatus"`


    


    


    Size int `json:"size"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Contactlist
type Contactlist struct { 
    


    // Name
    Name string `json:"name"`


    


    


    // Version - Required for updates, must match the version number of the most recent update
    Version int `json:"version"`


    // Division - The division this entity belongs to.
    Division Domainentityref `json:"division"`


    // ColumnNames - The names of the contact data columns.
    ColumnNames []string `json:"columnNames"`


    // PhoneColumns - Indicates which columns are phone numbers.
    PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`


    // EmailColumns - Indicates which columns are email addresses
    EmailColumns []Emailcolumn `json:"emailColumns"`


    // WhatsAppColumns - Indicates which columns are whatsApp contacts
    WhatsAppColumns []Whatsappcolumn `json:"whatsAppColumns"`


    


    // PreviewModeColumnName - A column to check if a contact should always be dialed in preview mode.
    PreviewModeColumnName string `json:"previewModeColumnName"`


    // PreviewModeAcceptedValues - The values in the previewModeColumnName column that indicate a contact should always be dialed in preview mode.
    PreviewModeAcceptedValues []string `json:"previewModeAcceptedValues"`


    


    // AttemptLimits - AttemptLimits for this ContactList.
    AttemptLimits Domainentityref `json:"attemptLimits"`


    // AutomaticTimeZoneMapping - Indicates if automatic time zone mapping is to be used for this ContactList.
    AutomaticTimeZoneMapping bool `json:"automaticTimeZoneMapping"`


    // ZipCodeColumnName - The name of contact list column containing the zip code for use with automatic time zone mapping. Only allowed if 'automaticTimeZoneMapping' is set to true.
    ZipCodeColumnName string `json:"zipCodeColumnName"`


    // ColumnDataTypeSpecifications - The settings of the columns selected for dynamic queueing
    ColumnDataTypeSpecifications []Columndatatypespecification `json:"columnDataTypeSpecifications"`


    // TrimWhitespace - Whether to trim white space when importing a contactlist csv file, default value = true
    TrimWhitespace bool `json:"trimWhitespace"`


    

}

// String returns a JSON representation of the model
func (o *Contactlist) String() string {
    
    
    
     o.ColumnNames = []string{""} 
     o.PhoneColumns = []Contactphonenumbercolumn{{}} 
     o.EmailColumns = []Emailcolumn{{}} 
     o.WhatsAppColumns = []Whatsappcolumn{{}} 
    
     o.PreviewModeAcceptedValues = []string{""} 
    
    
    
     o.ColumnDataTypeSpecifications = []Columndatatypespecification{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactlist) MarshalJSON() ([]byte, error) {
    type Alias Contactlist

    if ContactlistMarshalled {
        return []byte("{}"), nil
    }
    ContactlistMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Version int `json:"version"`
        
        Division Domainentityref `json:"division"`
        
        ColumnNames []string `json:"columnNames"`
        
        PhoneColumns []Contactphonenumbercolumn `json:"phoneColumns"`
        
        EmailColumns []Emailcolumn `json:"emailColumns"`
        
        WhatsAppColumns []Whatsappcolumn `json:"whatsAppColumns"`
        
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
        


        
        WhatsAppColumns: []Whatsappcolumn{{}},
        


        


        


        
        PreviewModeAcceptedValues: []string{""},
        


        


        


        


        


        
        ColumnDataTypeSpecifications: []Columndatatypespecification{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

