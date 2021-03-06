// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateTagDetails The representation of UpdateTagDetails
type UpdateTagDetails struct {

	// The description you assign to the tag during creation.
	Description *string `mandatory:"false" json:"description"`

	// Whether the tag is retired.
	// See Retiring Key Definitions and Namespace Definitions (https://docs.cloud.oracle.com/Content/Identity/Concepts/taggingoverview.htm#Retiring).
	IsRetired *bool `mandatory:"false" json:"isRetired"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Indicates whether the tag is enabled for cost tracking.
	IsCostTracking *bool `mandatory:"false" json:"isCostTracking"`

	// Additional validation rule for values specified for the tag definition.
	// If no validator is defined for a tag definition, then any (valid) value will be accepted.
	// The default value for `validator` is an empty map (no additional validation).
	Validator BaseTagDefinitionValidator `mandatory:"false" json:"validator"`
}

func (m UpdateTagDetails) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTagDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description    *string                           `json:"description"`
		IsRetired      *bool                             `json:"isRetired"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		IsCostTracking *bool                             `json:"isCostTracking"`
		Validator      basetagdefinitionvalidator        `json:"validator"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Description = model.Description
	m.IsRetired = model.IsRetired
	m.FreeformTags = model.FreeformTags
	m.DefinedTags = model.DefinedTags
	m.IsCostTracking = model.IsCostTracking
	nn, e := model.Validator.UnmarshalPolymorphicJSON(model.Validator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Validator = nn.(BaseTagDefinitionValidator)
	} else {
		m.Validator = nil
	}
	return
}
