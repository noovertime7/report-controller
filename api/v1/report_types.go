/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ReportSpec defines the desired state of Report
type ReportSpec struct {
	Pull     Pull   `json:"pull"`
	Schedule string `json:"schedule"`
	Save     Save   `json:"save"`
	Send     Send   `json:"send"`
}

type Pull struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
}

type Save struct {
	Type      SaveType `json:"type"`
	Endpoint  string   `json:"endpoint"`
	AccessKey string   `json:"accessKey"`
	AccessID  string   `json:"accessID"`
	Region    string   `json:"region,omitempty"`
}

type Send struct {
	Type SendType `json:"type"`
}

type SaveType string

const (
	MinioSaveType SaveType = "minio"
)

type SendType string

const (
	Email SendType = "email"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Report is the Schema for the reports API
type Report struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ReportSpec   `json:"spec,omitempty"`
	Status ReportStatus `json:"status,omitempty"`
}

// ReportStatus defines the observed state of Report
type ReportStatus struct {
	// Phase defines the current operation that the backup process is taking.
	Phase ReportPhase `json:"phase,omitempty"`
	// StartTime is the times that this backup entered the `BackingUp' phase.
	// +optional
	StartTime *metav1.Time `json:"startTime,omitempty"`
	// CompletionTime is the time that this backup entered the `Completed' phase.
	// +optional
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
}

//+kubebuilder:object:root=true

// ReportList contains a list of Report
type ReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Report `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Report{}, &ReportList{})
}
