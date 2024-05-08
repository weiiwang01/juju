// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/provider/oci (interfaces: ComputeClient)
//
// Generated by this command:
//
//	mockgen -typed -package testing -destination testing/mocks_compute.go -write_package_comment=false github.com/juju/juju/internal/provider/oci ComputeClient
package testing

import (
	context "context"
	reflect "reflect"

	core "github.com/oracle/oci-go-sdk/v65/core"
	gomock "go.uber.org/mock/gomock"
)

// MockComputeClient is a mock of ComputeClient interface.
type MockComputeClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeClientMockRecorder
}

// MockComputeClientMockRecorder is the mock recorder for MockComputeClient.
type MockComputeClientMockRecorder struct {
	mock *MockComputeClient
}

// NewMockComputeClient creates a new mock instance.
func NewMockComputeClient(ctrl *gomock.Controller) *MockComputeClient {
	mock := &MockComputeClient{ctrl: ctrl}
	mock.recorder = &MockComputeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockComputeClient) EXPECT() *MockComputeClientMockRecorder {
	return m.recorder
}

// AttachVolume mocks base method.
func (m *MockComputeClient) AttachVolume(arg0 context.Context, arg1 core.AttachVolumeRequest) (core.AttachVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachVolume", arg0, arg1)
	ret0, _ := ret[0].(core.AttachVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AttachVolume indicates an expected call of AttachVolume.
func (mr *MockComputeClientMockRecorder) AttachVolume(arg0, arg1 any) *MockComputeClientAttachVolumeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachVolume", reflect.TypeOf((*MockComputeClient)(nil).AttachVolume), arg0, arg1)
	return &MockComputeClientAttachVolumeCall{Call: call}
}

// MockComputeClientAttachVolumeCall wrap *gomock.Call
type MockComputeClientAttachVolumeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientAttachVolumeCall) Return(arg0 core.AttachVolumeResponse, arg1 error) *MockComputeClientAttachVolumeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientAttachVolumeCall) Do(f func(context.Context, core.AttachVolumeRequest) (core.AttachVolumeResponse, error)) *MockComputeClientAttachVolumeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientAttachVolumeCall) DoAndReturn(f func(context.Context, core.AttachVolumeRequest) (core.AttachVolumeResponse, error)) *MockComputeClientAttachVolumeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DetachVolume mocks base method.
func (m *MockComputeClient) DetachVolume(arg0 context.Context, arg1 core.DetachVolumeRequest) (core.DetachVolumeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachVolume", arg0, arg1)
	ret0, _ := ret[0].(core.DetachVolumeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachVolume indicates an expected call of DetachVolume.
func (mr *MockComputeClientMockRecorder) DetachVolume(arg0, arg1 any) *MockComputeClientDetachVolumeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachVolume", reflect.TypeOf((*MockComputeClient)(nil).DetachVolume), arg0, arg1)
	return &MockComputeClientDetachVolumeCall{Call: call}
}

// MockComputeClientDetachVolumeCall wrap *gomock.Call
type MockComputeClientDetachVolumeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientDetachVolumeCall) Return(arg0 core.DetachVolumeResponse, arg1 error) *MockComputeClientDetachVolumeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientDetachVolumeCall) Do(f func(context.Context, core.DetachVolumeRequest) (core.DetachVolumeResponse, error)) *MockComputeClientDetachVolumeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientDetachVolumeCall) DoAndReturn(f func(context.Context, core.DetachVolumeRequest) (core.DetachVolumeResponse, error)) *MockComputeClientDetachVolumeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInstance mocks base method.
func (m *MockComputeClient) GetInstance(arg0 context.Context, arg1 core.GetInstanceRequest) (core.GetInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstance", arg0, arg1)
	ret0, _ := ret[0].(core.GetInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockComputeClientMockRecorder) GetInstance(arg0, arg1 any) *MockComputeClientGetInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockComputeClient)(nil).GetInstance), arg0, arg1)
	return &MockComputeClientGetInstanceCall{Call: call}
}

// MockComputeClientGetInstanceCall wrap *gomock.Call
type MockComputeClientGetInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientGetInstanceCall) Return(arg0 core.GetInstanceResponse, arg1 error) *MockComputeClientGetInstanceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientGetInstanceCall) Do(f func(context.Context, core.GetInstanceRequest) (core.GetInstanceResponse, error)) *MockComputeClientGetInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientGetInstanceCall) DoAndReturn(f func(context.Context, core.GetInstanceRequest) (core.GetInstanceResponse, error)) *MockComputeClientGetInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetVolumeAttachment mocks base method.
func (m *MockComputeClient) GetVolumeAttachment(arg0 context.Context, arg1 core.GetVolumeAttachmentRequest) (core.GetVolumeAttachmentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeAttachment", arg0, arg1)
	ret0, _ := ret[0].(core.GetVolumeAttachmentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeAttachment indicates an expected call of GetVolumeAttachment.
func (mr *MockComputeClientMockRecorder) GetVolumeAttachment(arg0, arg1 any) *MockComputeClientGetVolumeAttachmentCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeAttachment", reflect.TypeOf((*MockComputeClient)(nil).GetVolumeAttachment), arg0, arg1)
	return &MockComputeClientGetVolumeAttachmentCall{Call: call}
}

// MockComputeClientGetVolumeAttachmentCall wrap *gomock.Call
type MockComputeClientGetVolumeAttachmentCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientGetVolumeAttachmentCall) Return(arg0 core.GetVolumeAttachmentResponse, arg1 error) *MockComputeClientGetVolumeAttachmentCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientGetVolumeAttachmentCall) Do(f func(context.Context, core.GetVolumeAttachmentRequest) (core.GetVolumeAttachmentResponse, error)) *MockComputeClientGetVolumeAttachmentCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientGetVolumeAttachmentCall) DoAndReturn(f func(context.Context, core.GetVolumeAttachmentRequest) (core.GetVolumeAttachmentResponse, error)) *MockComputeClientGetVolumeAttachmentCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LaunchInstance mocks base method.
func (m *MockComputeClient) LaunchInstance(arg0 context.Context, arg1 core.LaunchInstanceRequest) (core.LaunchInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LaunchInstance", arg0, arg1)
	ret0, _ := ret[0].(core.LaunchInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LaunchInstance indicates an expected call of LaunchInstance.
func (mr *MockComputeClientMockRecorder) LaunchInstance(arg0, arg1 any) *MockComputeClientLaunchInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LaunchInstance", reflect.TypeOf((*MockComputeClient)(nil).LaunchInstance), arg0, arg1)
	return &MockComputeClientLaunchInstanceCall{Call: call}
}

// MockComputeClientLaunchInstanceCall wrap *gomock.Call
type MockComputeClientLaunchInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientLaunchInstanceCall) Return(arg0 core.LaunchInstanceResponse, arg1 error) *MockComputeClientLaunchInstanceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientLaunchInstanceCall) Do(f func(context.Context, core.LaunchInstanceRequest) (core.LaunchInstanceResponse, error)) *MockComputeClientLaunchInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientLaunchInstanceCall) DoAndReturn(f func(context.Context, core.LaunchInstanceRequest) (core.LaunchInstanceResponse, error)) *MockComputeClientLaunchInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListImages mocks base method.
func (m *MockComputeClient) ListImages(arg0 context.Context, arg1 *string) ([]core.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]core.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockComputeClientMockRecorder) ListImages(arg0, arg1 any) *MockComputeClientListImagesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockComputeClient)(nil).ListImages), arg0, arg1)
	return &MockComputeClientListImagesCall{Call: call}
}

// MockComputeClientListImagesCall wrap *gomock.Call
type MockComputeClientListImagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientListImagesCall) Return(arg0 []core.Image, arg1 error) *MockComputeClientListImagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientListImagesCall) Do(f func(context.Context, *string) ([]core.Image, error)) *MockComputeClientListImagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientListImagesCall) DoAndReturn(f func(context.Context, *string) ([]core.Image, error)) *MockComputeClientListImagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListInstances mocks base method.
func (m *MockComputeClient) ListInstances(arg0 context.Context, arg1 *string) ([]core.Instance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInstances", arg0, arg1)
	ret0, _ := ret[0].([]core.Instance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockComputeClientMockRecorder) ListInstances(arg0, arg1 any) *MockComputeClientListInstancesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockComputeClient)(nil).ListInstances), arg0, arg1)
	return &MockComputeClientListInstancesCall{Call: call}
}

// MockComputeClientListInstancesCall wrap *gomock.Call
type MockComputeClientListInstancesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientListInstancesCall) Return(arg0 []core.Instance, arg1 error) *MockComputeClientListInstancesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientListInstancesCall) Do(f func(context.Context, *string) ([]core.Instance, error)) *MockComputeClientListInstancesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientListInstancesCall) DoAndReturn(f func(context.Context, *string) ([]core.Instance, error)) *MockComputeClientListInstancesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListShapes mocks base method.
func (m *MockComputeClient) ListShapes(arg0 context.Context, arg1, arg2 *string) ([]core.Shape, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShapes", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.Shape)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListShapes indicates an expected call of ListShapes.
func (mr *MockComputeClientMockRecorder) ListShapes(arg0, arg1, arg2 any) *MockComputeClientListShapesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShapes", reflect.TypeOf((*MockComputeClient)(nil).ListShapes), arg0, arg1, arg2)
	return &MockComputeClientListShapesCall{Call: call}
}

// MockComputeClientListShapesCall wrap *gomock.Call
type MockComputeClientListShapesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientListShapesCall) Return(arg0 []core.Shape, arg1 error) *MockComputeClientListShapesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientListShapesCall) Do(f func(context.Context, *string, *string) ([]core.Shape, error)) *MockComputeClientListShapesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientListShapesCall) DoAndReturn(f func(context.Context, *string, *string) ([]core.Shape, error)) *MockComputeClientListShapesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListVnicAttachments mocks base method.
func (m *MockComputeClient) ListVnicAttachments(arg0 context.Context, arg1, arg2 *string) ([]core.VnicAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVnicAttachments", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.VnicAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVnicAttachments indicates an expected call of ListVnicAttachments.
func (mr *MockComputeClientMockRecorder) ListVnicAttachments(arg0, arg1, arg2 any) *MockComputeClientListVnicAttachmentsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVnicAttachments", reflect.TypeOf((*MockComputeClient)(nil).ListVnicAttachments), arg0, arg1, arg2)
	return &MockComputeClientListVnicAttachmentsCall{Call: call}
}

// MockComputeClientListVnicAttachmentsCall wrap *gomock.Call
type MockComputeClientListVnicAttachmentsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientListVnicAttachmentsCall) Return(arg0 []core.VnicAttachment, arg1 error) *MockComputeClientListVnicAttachmentsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientListVnicAttachmentsCall) Do(f func(context.Context, *string, *string) ([]core.VnicAttachment, error)) *MockComputeClientListVnicAttachmentsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientListVnicAttachmentsCall) DoAndReturn(f func(context.Context, *string, *string) ([]core.VnicAttachment, error)) *MockComputeClientListVnicAttachmentsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListVolumeAttachments mocks base method.
func (m *MockComputeClient) ListVolumeAttachments(arg0 context.Context, arg1, arg2 *string) ([]core.VolumeAttachment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumeAttachments", arg0, arg1, arg2)
	ret0, _ := ret[0].([]core.VolumeAttachment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumeAttachments indicates an expected call of ListVolumeAttachments.
func (mr *MockComputeClientMockRecorder) ListVolumeAttachments(arg0, arg1, arg2 any) *MockComputeClientListVolumeAttachmentsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumeAttachments", reflect.TypeOf((*MockComputeClient)(nil).ListVolumeAttachments), arg0, arg1, arg2)
	return &MockComputeClientListVolumeAttachmentsCall{Call: call}
}

// MockComputeClientListVolumeAttachmentsCall wrap *gomock.Call
type MockComputeClientListVolumeAttachmentsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientListVolumeAttachmentsCall) Return(arg0 []core.VolumeAttachment, arg1 error) *MockComputeClientListVolumeAttachmentsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientListVolumeAttachmentsCall) Do(f func(context.Context, *string, *string) ([]core.VolumeAttachment, error)) *MockComputeClientListVolumeAttachmentsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientListVolumeAttachmentsCall) DoAndReturn(f func(context.Context, *string, *string) ([]core.VolumeAttachment, error)) *MockComputeClientListVolumeAttachmentsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TerminateInstance mocks base method.
func (m *MockComputeClient) TerminateInstance(arg0 context.Context, arg1 core.TerminateInstanceRequest) (core.TerminateInstanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TerminateInstance", arg0, arg1)
	ret0, _ := ret[0].(core.TerminateInstanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TerminateInstance indicates an expected call of TerminateInstance.
func (mr *MockComputeClientMockRecorder) TerminateInstance(arg0, arg1 any) *MockComputeClientTerminateInstanceCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateInstance", reflect.TypeOf((*MockComputeClient)(nil).TerminateInstance), arg0, arg1)
	return &MockComputeClientTerminateInstanceCall{Call: call}
}

// MockComputeClientTerminateInstanceCall wrap *gomock.Call
type MockComputeClientTerminateInstanceCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockComputeClientTerminateInstanceCall) Return(arg0 core.TerminateInstanceResponse, arg1 error) *MockComputeClientTerminateInstanceCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockComputeClientTerminateInstanceCall) Do(f func(context.Context, core.TerminateInstanceRequest) (core.TerminateInstanceResponse, error)) *MockComputeClientTerminateInstanceCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockComputeClientTerminateInstanceCall) DoAndReturn(f func(context.Context, core.TerminateInstanceRequest) (core.TerminateInstanceResponse, error)) *MockComputeClientTerminateInstanceCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
