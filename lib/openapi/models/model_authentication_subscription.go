/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthenticationSubscription struct {
	AuthenticationMethod               AuthMethod      `json:"authenticationMethod" bson:"authenticationMethod" yaml:"authenticationMethod"`
	PermanentKey                       *PermanentKey   `json:"permanentKey" bson:"permanentKey" yaml:"permanentKey"`
	SequenceNumber                     string          `json:"sequenceNumber" bson:"sequenceNumber" yaml:"sequenceNumber"`
	AuthenticationManagementField      string          `json:"authenticationManagementField,omitempty" bson:"authenticationManagementField" yaml:"authenticationManagementField"`
	VectorAlgorithm                    VectorAlgorithm `json:"vectorAlgorithm,omitempty" bson:"vectorAlgorithm" yaml:"vectorAlgorithm"`
	Milenage                           *Milenage       `json:"milenage,omitempty" bson:"milenage" yaml:"milenage,omitempty"`
	Tuak                               *Tuak           `json:"tuak,omitempty" bson:"tuak" yaml:"tuak,omitempty"`
	Opc                                *Opc            `json:"opc,omitempty" bson:"opc" yaml:"opc,omitempty"`
	Topc                               *Topc           `json:"topc,omitempty" bson:"topc" yaml:"topc,omitempty"`
	SharedAuthenticationSubscriptionId *SharedData     `json:"sharedAuthenticationSubscriptionId,omitempty" bson:"sharedAuthenticationSubscriptionId" yaml:"sharedAuthenticationSubscriptionId,omitempty"`
}