package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "sync"

// TabulationIdentity represents the Identity of the object
var TabulationIdentity = elemental.Identity{
	Name:     "tabulation",
	Category: "tabulations",
}

// TabulationsList represents a list of Tabulations
type TabulationsList []*Tabulation

// ContentIdentity returns the identity of the objects in the list.
func (o TabulationsList) ContentIdentity() elemental.Identity {

	return TabulationIdentity
}

// List converts the object to and elemental.IdentifiablesList.
func (o TabulationsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TabulationsList) DefaultOrder() []string {

	return []string{}
}

// Tabulation represents the model of a tabulation
type Tabulation struct {
	// Headers contains the requests headers that matched.
	Headers []string `json:"headers" bson:"-"`

	// Rows contains the tabulated data.
	Rows [][]interface{} `json:"rows" bson:"-"`

	// TargetIdentity contains the requested target identity.
	TargetIdentity string `json:"targetIdentity" bson:"-"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`

	sync.Mutex
}

// NewTabulation returns a new *Tabulation
func NewTabulation() *Tabulation {

	return &Tabulation{
		ModelVersion: 1.0,
		Rows:         [][]interface{}{},
	}
}

// Identity returns the Identity of the object.
func (o *Tabulation) Identity() elemental.Identity {

	return TabulationIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Tabulation) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Tabulation) SetIdentifier(ID string) {

}

// Version returns the hardcoded version of the model
func (o *Tabulation) Version() float64 {

	return 1.0
}

// DefaultOrder returns the list of default ordering fields.
func (o *Tabulation) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Tabulation) Doc() string {
	return `Tabulate API allows you to retrieve a custom table view for any identity using any tags you like as columns.`
}

func (o *Tabulation) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *Tabulation) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Tabulation) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return TabulationAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Tabulation) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return TabulationAttributesMap
}

// TabulationAttributesMap represents the map of attribute for Tabulation.
var TabulationAttributesMap = map[string]elemental.AttributeSpecification{
	"Headers": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Headers contains the requests headers that matched.`,
		Exposed:        true,
		Name:           "headers",
		ReadOnly:       true,
		SubType:        "string",
		Type:           "list",
	},
	"Rows": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `Rows contains the tabulated data.`,
		Exposed:        true,
		Name:           "rows",
		ReadOnly:       true,
		SubType:        "tabulated_data",
		Type:           "external",
	},
	"TargetIdentity": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `TargetIdentity contains the requested target identity.`,
		Exposed:        true,
		Format:         "free",
		Name:           "targetIdentity",
		ReadOnly:       true,
		Type:           "string",
	},
}