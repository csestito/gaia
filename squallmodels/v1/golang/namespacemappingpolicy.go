package squallmodels

import "fmt"
import "github.com/aporeto-inc/elemental"

import "time"

// NamespaceMappingPolicyIdentity represents the Identity of the object
var NamespaceMappingPolicyIdentity = elemental.Identity{
	Name:     "namespacemappingpolicy",
	Category: "namespacemappingpolicies",
}

// NamespaceMappingPoliciesList represents a list of NamespaceMappingPolicies
type NamespaceMappingPoliciesList []*NamespaceMappingPolicy

// ContentIdentity returns the identity of the objects in the list.
func (o NamespaceMappingPoliciesList) ContentIdentity() elemental.Identity {
	return NamespaceMappingPolicyIdentity
}

// List convert the object to and elemental.IdentifiablesList.
func (o NamespaceMappingPoliciesList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// NamespaceMappingPolicy represents the model of a namespacemappingpolicy
type NamespaceMappingPolicy struct {
	// ID is the identifier of the object.
	ID string `json:"ID" bson:"-"`

	// Annotation stores additional information about an entity
	Annotation map[string]string `json:"annotation" bson:"annotation"`

	// AssociatedTags are the list of tags attached to an entity
	AssociatedTags []string `json:"associatedTags" bson:"associatedtags"`

	// CreatedTime is the time at which the object was created
	CreateTime time.Time `json:"createTime" bson:"createtime"`

	// Description is the description of the object.
	Description string `json:"description" bson:"description"`

	// Disabled defines if the propert is disabled.
	Disabled bool `json:"disabled" bson:"disabled"`

	// mappedNamespace is the mapped namespace
	MappedNamespace string `json:"mappedNamespace" bson:"mappednamespace"`

	// Name is the name of the entity
	Name string `json:"name" bson:"name"`

	// Namespace tag attached to an entity
	Namespace string `json:"namespace" bson:"namespace"`

	// NormalizedTags contains the list of normalized tags of the entities
	NormalizedTags []string `json:"normalizedTags" bson:"normalizedtags"`

	// Protected defines if the object is protected.
	Protected bool `json:"protected" bson:"protected"`

	// Subject is the subject.
	Subject [][]string `json:"subject" bson:"-"`

	// UpdateTime is the time at which an entity was updated.
	UpdateTime time.Time `json:"updateTime" bson:"updatetime"`

	ModelVersion float64 `json:"-" bson:"_modelversion"`
}

// NewNamespaceMappingPolicy returns a new *NamespaceMappingPolicy
func NewNamespaceMappingPolicy() *NamespaceMappingPolicy {

	return &NamespaceMappingPolicy{
		ModelVersion:   1.0,
		AssociatedTags: []string{},
		NormalizedTags: []string{},
	}
}

// Identity returns the Identity of the object.
func (o *NamespaceMappingPolicy) Identity() elemental.Identity {

	return NamespaceMappingPolicyIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *NamespaceMappingPolicy) SetIdentifier(ID string) {

	o.ID = ID
}

// Version returns the hardcoded version of the model
func (o *NamespaceMappingPolicy) Version() float64 {

	return 1.0
}

func (o *NamespaceMappingPolicy) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetAssociatedTags returns the associatedTags of the receiver
func (o *NamespaceMappingPolicy) GetAssociatedTags() []string {
	return o.AssociatedTags
}

// SetAssociatedTags set the given associatedTags of the receiver
func (o *NamespaceMappingPolicy) SetAssociatedTags(associatedTags []string) {
	o.AssociatedTags = associatedTags
}

// SetCreateTime set the given createTime of the receiver
func (o *NamespaceMappingPolicy) SetCreateTime(createTime time.Time) {
	o.CreateTime = createTime
}

// GetDisabled returns the disabled of the receiver
func (o *NamespaceMappingPolicy) GetDisabled() bool {
	return o.Disabled
}

// SetDisabled set the given disabled of the receiver
func (o *NamespaceMappingPolicy) SetDisabled(disabled bool) {
	o.Disabled = disabled
}

// GetName returns the name of the receiver
func (o *NamespaceMappingPolicy) GetName() string {
	return o.Name
}

// SetName set the given name of the receiver
func (o *NamespaceMappingPolicy) SetName(name string) {
	o.Name = name
}

// GetNamespace returns the namespace of the receiver
func (o *NamespaceMappingPolicy) GetNamespace() string {
	return o.Namespace
}

// SetNamespace set the given namespace of the receiver
func (o *NamespaceMappingPolicy) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// GetNormalizedTags returns the normalizedTags of the receiver
func (o *NamespaceMappingPolicy) GetNormalizedTags() []string {
	return o.NormalizedTags
}

// SetNormalizedTags set the given normalizedTags of the receiver
func (o *NamespaceMappingPolicy) SetNormalizedTags(normalizedTags []string) {
	o.NormalizedTags = normalizedTags
}

// GetProtected returns the protected of the receiver
func (o *NamespaceMappingPolicy) GetProtected() bool {
	return o.Protected
}

// SetUpdateTime set the given updateTime of the receiver
func (o *NamespaceMappingPolicy) SetUpdateTime(updateTime time.Time) {
	o.UpdateTime = updateTime
}

// Validate valides the current information stored into the structure.
func (o *NamespaceMappingPolicy) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("mappedNamespace", o.MappedNamespace); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("mappedNamespace", o.MappedNamespace); err != nil {
		errors = append(errors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		errors = append(errors, err)
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
func (NamespaceMappingPolicy) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	return NamespaceMappingPolicyAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (NamespaceMappingPolicy) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return NamespaceMappingPolicyAttributesMap
}

// NamespaceMappingPolicyAttributesMap represents the map of attribute for NamespaceMappingPolicy.
var NamespaceMappingPolicyAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `ID is the identifier of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		ReadOnly:       true,
		Type:           "string",
		Unique:         true,
	},
	"Annotation": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Annotation stores additional information about an entity`,
		Exposed:        true,
		Name:           "annotation",
		Stored:         true,
		SubType:        "annotation",
		Type:           "external",
	},
	"AssociatedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `AssociatedTags are the list of tags attached to an entity`,
		Exposed:        true,
		Getter:         true,
		Name:           "associatedTags",
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Type:           "external",
	},
	"CreateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `CreatedTime is the time at which the object was created`,
		Exposed:        true,
		Name:           "createTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Description is the description of the object.`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Disabled": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Disabled defines if the propert is disabled.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "disabled",
		Orderable:      true,
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"MappedNamespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `mappedNamespace is the mapped namespace`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Name:           "mappedNamespace",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Name is the name of the entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
		Unique:         true,
	},
	"Namespace": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		CreationOnly:   true,
		Description:    `Namespace tag attached to an entity`,
		Exposed:        true,
		Filterable:     true,
		Format:         "free",
		Getter:         true,
		Index:          true,
		Name:           "namespace",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"NormalizedTags": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `NormalizedTags contains the list of normalized tags of the entities`,
		Exposed:        true,
		Getter:         true,
		Name:           "normalizedTags",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		SubType:        "tags_list",
		Transient:      true,
		Type:           "external",
	},
	"Protected": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Protected defines if the object is protected.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "protected",
		Orderable:      true,
		Stored:         true,
		Type:           "boolean",
	},
	"Subject": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Description:    `Subject is the subject.`,
		Exposed:        true,
		Name:           "subject",
		Orderable:      true,
		SubType:        "policies_list",
		Type:           "external",
	},
	"UpdateTime": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		Description:    `UpdateTime is the time at which an entity was updated.`,
		Exposed:        true,
		Name:           "updateTime",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "time",
	},
}