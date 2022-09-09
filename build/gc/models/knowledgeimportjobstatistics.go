package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeimportjobstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeimportjobstatisticsDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Knowledgeimportjobstatistics
type Knowledgeimportjobstatistics struct { 
    // CountDocumentImportActivityCreate - Number of documents will be created by the import.
    CountDocumentImportActivityCreate int `json:"countDocumentImportActivityCreate"`


    // CountDocumentImportActivityUpdate - Number of documents will be updated by the import.
    CountDocumentImportActivityUpdate int `json:"countDocumentImportActivityUpdate"`


    // CountDocumentStateDraft - Number of documents will be imported as draft.
    CountDocumentStateDraft int `json:"countDocumentStateDraft"`


    // CountDocumentStatePublished - Number of documents will be imported as published.
    CountDocumentStatePublished int `json:"countDocumentStatePublished"`


    // CountDocumentValidationSuccess - Number of documents that validated successfully for import.
    CountDocumentValidationSuccess int `json:"countDocumentValidationSuccess"`


    // CountDocumentValidationFailure - Number of documents that failed validation for import.
    CountDocumentValidationFailure int `json:"countDocumentValidationFailure"`


    // CountDocumentImportSuccess - Number of imported documents.
    CountDocumentImportSuccess int `json:"countDocumentImportSuccess"`


    // CountDocumentImportFailure - Number of documents failed to import.
    CountDocumentImportFailure int `json:"countDocumentImportFailure"`


    // CountCategoryValidationSuccess - Number of categories that validated successfully for import.
    CountCategoryValidationSuccess int `json:"countCategoryValidationSuccess"`


    // CountCategoryValidationFailure - Number of categories that failed validation for import.
    CountCategoryValidationFailure int `json:"countCategoryValidationFailure"`


    // CountCategoryImportSuccess - Number of imported categories.
    CountCategoryImportSuccess int `json:"countCategoryImportSuccess"`


    // CountCategoryImportFailure - Number of categories failed to import.
    CountCategoryImportFailure int `json:"countCategoryImportFailure"`


    // CountLabelValidationSuccess - Number of labels that validated successfully for import.
    CountLabelValidationSuccess int `json:"countLabelValidationSuccess"`


    // CountLabelValidationFailure - Number of labels that failed validation for import.
    CountLabelValidationFailure int `json:"countLabelValidationFailure"`


    // CountLabelImportSuccess - Number of imported labels.
    CountLabelImportSuccess int `json:"countLabelImportSuccess"`


    // CountLabelImportFailure - Number of labels failed to import.
    CountLabelImportFailure int `json:"countLabelImportFailure"`


    // MigrationDetected - Shows whether the import treated as migration or not.
    MigrationDetected bool `json:"migrationDetected"`

}

// String returns a JSON representation of the model
func (o *Knowledgeimportjobstatistics) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeimportjobstatistics) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeimportjobstatistics

    if KnowledgeimportjobstatisticsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeimportjobstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        CountDocumentImportActivityCreate int `json:"countDocumentImportActivityCreate"`
        
        CountDocumentImportActivityUpdate int `json:"countDocumentImportActivityUpdate"`
        
        CountDocumentStateDraft int `json:"countDocumentStateDraft"`
        
        CountDocumentStatePublished int `json:"countDocumentStatePublished"`
        
        CountDocumentValidationSuccess int `json:"countDocumentValidationSuccess"`
        
        CountDocumentValidationFailure int `json:"countDocumentValidationFailure"`
        
        CountDocumentImportSuccess int `json:"countDocumentImportSuccess"`
        
        CountDocumentImportFailure int `json:"countDocumentImportFailure"`
        
        CountCategoryValidationSuccess int `json:"countCategoryValidationSuccess"`
        
        CountCategoryValidationFailure int `json:"countCategoryValidationFailure"`
        
        CountCategoryImportSuccess int `json:"countCategoryImportSuccess"`
        
        CountCategoryImportFailure int `json:"countCategoryImportFailure"`
        
        CountLabelValidationSuccess int `json:"countLabelValidationSuccess"`
        
        CountLabelValidationFailure int `json:"countLabelValidationFailure"`
        
        CountLabelImportSuccess int `json:"countLabelImportSuccess"`
        
        CountLabelImportFailure int `json:"countLabelImportFailure"`
        
        MigrationDetected bool `json:"migrationDetected"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

