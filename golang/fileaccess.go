package gaia

import "fmt"
import "github.com/aporeto-inc/elemental"

// FileAccessModeValue represents the possible values for attribute "mode".
type FileAccessModeValue string

const (
	// FileAccessModeRead represents the value Read.
	FileAccessModeRead FileAccessModeValue = "Read"

	// FileAccessModeReadwrite represents the value ReadWrite.
	FileAccessModeReadwrite FileAccessModeValue = "ReadWrite"

	// FileAccessModeWrite represents the value Write.
	FileAccessModeWrite FileAccessModeValue = "Write"
)

// FileAccessIdentity represents the Identity of the object
var FileAccessIdentity = elemental.Identity{
	Name:     "fileaccess",
	Category: "fileaccesses",
}

// FileAccessList represents a list of FileAccess
type FileAccessList []*FileAccess

// FileAccess represents the model of a fileaccess
type FileAccess struct {
	// Action tells if the access has been allowed or not.
	Action string `json:"action" cql:"-" bson:"-"`

	// Count tells how many times the file has been accessed.
	Count int `json:"count" cql:"-" bson:"-"`

	// Host is the host that served the accessed file.
	Host string `json:"host" cql:"-" bson:"-"`

	// Mode is the mode of the accessed file.
	Mode FileAccessModeValue `json:"mode" cql:"-" bson:"-"`

	// Path is the path of the accessed file.
	Path string `json:"path" cql:"-" bson:"-"`

	// Protocol is the protocol used to access the file.
	Protocol string `json:"protocol" cql:"-" bson:"-"`
}

// NewFileAccess returns a new *FileAccess
func NewFileAccess() *FileAccess {

	return &FileAccess{}
}

// Identity returns the Identity of the object.
func (o *FileAccess) Identity() elemental.Identity {

	return FileAccessIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *FileAccess) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *FileAccess) SetIdentifier(ID string) {

}

func (o *FileAccess) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *FileAccess) Validate() error {

	errors := elemental.Errors{}

	if err := elemental.ValidateStringInList("mode", string(o.Mode), []string{"Read", "ReadWrite", "Write"}, true); err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (FileAccess) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return FileAccessAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (FileAccess) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return FileAccessAttributesMap
}

// FileAccessAttributesMap represents the map of attribute for FileAccess.
var FileAccessAttributesMap = map[string]elemental.AttributeSpecification{
	"Action": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    "Action tells if the access has been allowed or not.",
		Exposed:        true,
		Format:         "free",
		Name:           "action",
		ReadOnly:       true,
		Type:           "string",
	},
	"Count": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    "Count tells how many times the file has been accessed.",
		Exposed:        true,
		Name:           "count",
		ReadOnly:       true,
		Type:           "integer",
	},
	"Host": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    "Host is the host that served the accessed file.",
		Exposed:        true,
		Format:         "free",
		Name:           "host",
		ReadOnly:       true,
		Type:           "string",
	},
	"Mode": elemental.AttributeSpecification{
		AllowedChoices: []string{"Read", "ReadWrite", "Write"},
		Autogenerated:  true,
		Description:    "Mode is the mode of the accessed file.",
		Exposed:        true,
		Name:           "mode",
		ReadOnly:       true,
		Type:           "enum",
	},
	"Path": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    "Path is the path of the accessed file.",
		Exposed:        true,
		Format:         "free",
		Name:           "path",
		ReadOnly:       true,
		Type:           "string",
	},
	"Protocol": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    "Protocol is the protocol used to access the file.",
		Exposed:        true,
		Format:         "free",
		Name:           "protocol",
		ReadOnly:       true,
		Type:           "string",
	},
}
