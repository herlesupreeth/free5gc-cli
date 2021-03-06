/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type SharedData struct {
	SharedDataId            string                             `json:"sharedDataId" yaml:"sharedDataId" bson:"sharedDataId" mapstructure:"SharedDataId"`
	SharedAmData            *AccessAndMobilitySubscriptionData `json:"sharedAmData,omitempty" yaml:"sharedAmData,omitempty" bson:"sharedAmData" mapstructure:"SharedAmData"`
	SharedSmsSubsData       *SmsSubscriptionData               `json:"sharedSmsSubsData,omitempty" yaml:"sharedSmsSubsData,omitempty" bson:"sharedSmsSubsData" mapstructure:"SharedSmsSubsData"`
	SharedSmsMngSubsData    *SmsManagementSubscriptionData     `json:"sharedSmsMngSubsData,omitempty" yaml:"sharedSmsMngSubsData,omitempty" bson:"sharedSmsMngSubsData" mapstructure:"SharedSmsMngSubsData"`
	SharedDnnConfigurations map[string]DnnConfiguration        `json:"sharedDnnConfigurations,omitempty" yaml:"sharedDnnConfigurations,omitempty" bson:"sharedDnnConfigurations" mapstructure:"SharedDnnConfigurations"`
	SharedTraceData         *TraceData                         `json:"sharedTraceData,omitempty" yaml:"sharedTraceData,omitempty" bson:"sharedTraceData" mapstructure:"SharedTraceData"`
	SharedSnssaiInfos       map[string]SnssaiInfo              `json:"sharedSnssaiInfos,omitempty" yaml:"sharedSnssaiInfos,omitempty" bson:"sharedSnssaiInfos" mapstructure:"SharedSnssaiInfos"`
}
