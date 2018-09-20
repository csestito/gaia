package gaia

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// EnforcerReportIdentity represents the Identity of the object.
var EnforcerReportIdentity = elemental.Identity{
	Name:     "enforcerreport",
	Category: "enforcerreports",
	Private:  false,
}

// EnforcerReportsList represents a list of EnforcerReports
type EnforcerReportsList []*EnforcerReport

// Identity returns the identity of the objects in the list.
func (o EnforcerReportsList) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Copy returns a pointer to a copy the EnforcerReportsList.
func (o EnforcerReportsList) Copy() elemental.Identifiables {

	copy := append(EnforcerReportsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the EnforcerReportsList.
func (o EnforcerReportsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(EnforcerReportsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*EnforcerReport))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o EnforcerReportsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o EnforcerReportsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o EnforcerReportsList) Version() int {

	return 1
}

// EnforcerReport represents the model of a enforcerreport
type EnforcerReport struct {
	// ID of the enforcer to report.
	ID string `json:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// Name of the enforcer to report.
	Name string `json:"name" bson:"-" mapstructure:"name,omitempty"`

	// Namespace of the enforcer to report.
	Namespace string `json:"namespace" bson:"-" mapstructure:"namespace,omitempty"`

	// Date of the report.
	Timestamp time.Time `json:"timestamp" bson:"-" mapstructure:"timestamp,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewEnforcerReport returns a new *EnforcerReport
func NewEnforcerReport() *EnforcerReport {

	return &EnforcerReport{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *EnforcerReport) Identity() elemental.Identity {

	return EnforcerReportIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *EnforcerReport) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *EnforcerReport) SetIdentifier(id string) {

}

// Version returns the hardcoded version of the model.
func (o *EnforcerReport) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *EnforcerReport) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *EnforcerReport) Doc() string {
	return `Post a new enforcer statistics report.`
}

func (o *EnforcerReport) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// Validate valides the current information stored into the structure.
func (o *EnforcerReport) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("ID", o.ID); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("namespace", o.Namespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredTime("timestamp", o.Timestamp); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*EnforcerReport) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := EnforcerReportAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return EnforcerReportLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*EnforcerReport) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return EnforcerReportAttributesMap
}

// EnforcerReportAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID of the enforcer to report.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer to report.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer to report.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"Timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}

// EnforcerReportLowerCaseAttributesMap represents the map of attribute for EnforcerReport.
var EnforcerReportLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ID",
		Description:    `ID of the enforcer to report.`,
		Exposed:        true,
		Name:           "ID",
		Required:       true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `Name of the enforcer to report.`,
		Exposed:        true,
		Name:           "name",
		Required:       true,
		Type:           "string",
	},
	"namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Namespace",
		Description:    `Namespace of the enforcer to report.`,
		Exposed:        true,
		Name:           "namespace",
		Required:       true,
		Type:           "string",
	},
	"timestamp": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Timestamp",
		Description:    `Date of the report.`,
		Exposed:        true,
		Name:           "timestamp",
		Required:       true,
		Type:           "time",
	},
}