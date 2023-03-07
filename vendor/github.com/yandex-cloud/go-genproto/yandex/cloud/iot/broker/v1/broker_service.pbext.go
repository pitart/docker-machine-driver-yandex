// Code generated by protoc-gen-goext. DO NOT EDIT.

package broker

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetBrokerRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *ListBrokersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListBrokersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBrokersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBrokersResponse) SetBrokers(v []*Broker) {
	m.Brokers = v
}

func (m *ListBrokersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateBrokerRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateBrokerRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateBrokerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateBrokerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateBrokerRequest) SetCertificates(v []*CreateBrokerRequest_Certificate) {
	m.Certificates = v
}

func (m *CreateBrokerRequest) SetPassword(v string) {
	m.Password = v
}

func (m *CreateBrokerRequest_Certificate) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *CreateBrokerMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *UpdateBrokerRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *UpdateBrokerRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateBrokerRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateBrokerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateBrokerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateBrokerMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *ListBrokerCertificatesRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *ListBrokerCertificatesResponse) SetCertificates(v []*BrokerCertificate) {
	m.Certificates = v
}

func (m *AddBrokerCertificateRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *AddBrokerCertificateRequest) SetCertificateData(v string) {
	m.CertificateData = v
}

func (m *AddBrokerCertificateMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *AddBrokerCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteBrokerCertificateRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerCertificateRequest) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *DeleteBrokerCertificateMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerCertificateMetadata) SetFingerprint(v string) {
	m.Fingerprint = v
}

func (m *ListBrokerPasswordsRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *ListBrokerPasswordsResponse) SetPasswords(v []*BrokerPassword) {
	m.Passwords = v
}

func (m *AddBrokerPasswordRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *AddBrokerPasswordRequest) SetPassword(v string) {
	m.Password = v
}

func (m *AddBrokerPasswordMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *AddBrokerPasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteBrokerPasswordRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerPasswordRequest) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *DeleteBrokerPasswordMetadata) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *DeleteBrokerPasswordMetadata) SetPasswordId(v string) {
	m.PasswordId = v
}

func (m *ListBrokerOperationsRequest) SetBrokerId(v string) {
	m.BrokerId = v
}

func (m *ListBrokerOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBrokerOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBrokerOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListBrokerOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListBrokerOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}