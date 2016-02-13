// Code generated by protoc-gen-go.
// source: google.golang.org/cloud/bigtable/internal/cluster_service_proto/bigtable_cluster_service.proto
// DO NOT EDIT!

package google_bigtable_admin_cluster_v1

import proto "github.com/nildev/account/Godeps/_workspace/src/github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_admin_cluster_v11 "github.com/nildev/account/Godeps/_workspace/src/google.golang.org/cloud/bigtable/internal/cluster_data_proto"
import google_protobuf "github.com/nildev/account/Godeps/_workspace/src/google.golang.org/cloud/bigtable/internal/empty"

import (
	context "github.com/nildev/account/Godeps/_workspace/src/golang.org/x/net/context"
	grpc "github.com/nildev/account/Godeps/_workspace/src/google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BigtableClusterService service

type BigtableClusterServiceClient interface {
	// Lists the supported zones for the given project.
	ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(ctx context.Context, in *google_bigtable_admin_cluster_v11.Cluster, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type bigtableClusterServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableClusterServiceClient(cc *grpc.ClientConn) BigtableClusterServiceClient {
	return &bigtableClusterServiceClient{cc}
}

func (c *bigtableClusterServiceClient) ListZones(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error) {
	out := new(ListZonesResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListZones", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) GetCluster(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/GetCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/ListClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/CreateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) UpdateCluster(ctx context.Context, in *google_bigtable_admin_cluster_v11.Cluster, opts ...grpc.CallOption) (*google_bigtable_admin_cluster_v11.Cluster, error) {
	out := new(google_bigtable_admin_cluster_v11.Cluster)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/UpdateCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableClusterServiceClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.admin.cluster.v1.BigtableClusterService/DeleteCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableClusterService service

type BigtableClusterServiceServer interface {
	// Lists the supported zones for the given project.
	ListZones(context.Context, *ListZonesRequest) (*ListZonesResponse, error)
	// Gets information about a particular cluster.
	GetCluster(context.Context, *GetClusterRequest) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Lists all clusters in the given project, along with any zones for which
	// cluster information could not be retrieved.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a cluster and begins preparing it to begin serving. The returned
	// cluster embeds as its "current_operation" a long-running operation which
	// can be used to track the progress of turning up the new cluster.
	// Immediately upon completion of this request:
	//  * The cluster will be readable via the API, with all requested attributes
	//    but no allocated resources.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will render the cluster immediately unreadable
	//    via the API.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// Upon completion of the embedded operation:
	//  * Billing for all successfully-allocated resources will begin (some types
	//    may have lower than the requested levels).
	//  * New tables can be created in the cluster.
	//  * The cluster's allocated resource levels will be readable via the API.
	// The embedded operation's "metadata" field type is
	// [CreateClusterMetadata][google.bigtable.admin.cluster.v1.CreateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	CreateCluster(context.Context, *CreateClusterRequest) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Updates a cluster, and begins allocating or releasing resources as
	// requested. The returned cluster embeds as its "current_operation" a
	// long-running operation which can be used to track the progress of updating
	// the cluster.
	// Immediately upon completion of this request:
	//  * For resource types where a decrease in the cluster's allocation has been
	//    requested, billing will be based on the newly-requested level.
	// Until completion of the embedded operation:
	//  * Cancelling the operation will set its metadata's "cancelled_at_time",
	//    and begin restoring resources to their pre-request values. The operation
	//    is guaranteed to succeed at undoing all resource changes, after which
	//    point it will terminate with a CANCELLED status.
	//  * All other attempts to modify or delete the cluster will be rejected.
	//  * Reading the cluster via the API will continue to give the pre-request
	//    resource levels.
	// Upon completion of the embedded operation:
	//  * Billing will begin for all successfully-allocated resources (some types
	//    may have lower than the requested levels).
	//  * All newly-reserved resources will be available for serving the cluster's
	//    tables.
	//  * The cluster's new resource levels will be readable via the API.
	// [UpdateClusterMetadata][google.bigtable.admin.cluster.v1.UpdateClusterMetadata] The embedded operation's "response" field type is
	// [Cluster][google.bigtable.admin.cluster.v1.Cluster], if successful.
	UpdateCluster(context.Context, *google_bigtable_admin_cluster_v11.Cluster) (*google_bigtable_admin_cluster_v11.Cluster, error)
	// Marks a cluster and all of its tables for permanent deletion in 7 days.
	// Immediately upon completion of the request:
	//  * Billing will cease for all of the cluster's reserved resources.
	//  * The cluster's "delete_time" field will be set 7 days in the future.
	// Soon afterward:
	//  * All tables within the cluster will become unavailable.
	// Prior to the cluster's "delete_time":
	//  * The cluster can be recovered with a call to UndeleteCluster.
	//  * All other attempts to modify or delete the cluster will be rejected.
	// At the cluster's "delete_time":
	//  * The cluster and *all of its tables* will immediately and irrevocably
	//    disappear from the API, and their data will be permanently deleted.
	DeleteCluster(context.Context, *DeleteClusterRequest) (*google_protobuf.Empty, error)
}

func RegisterBigtableClusterServiceServer(s *grpc.Server, srv BigtableClusterServiceServer) {
	s.RegisterService(&_BigtableClusterService_serviceDesc, srv)
}

func _BigtableClusterService_ListZones_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).ListZones(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).GetCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).ListClusters(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).CreateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_UpdateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_bigtable_admin_cluster_v11.Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).UpdateCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableClusterService_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableClusterServiceServer).DeleteCluster(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BigtableClusterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.admin.cluster.v1.BigtableClusterService",
	HandlerType: (*BigtableClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListZones",
			Handler:    _BigtableClusterService_ListZones_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _BigtableClusterService_GetCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _BigtableClusterService_ListClusters_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _BigtableClusterService_CreateCluster_Handler,
		},
		{
			MethodName: "UpdateCluster",
			Handler:    _BigtableClusterService_UpdateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _BigtableClusterService_DeleteCluster_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0x3d, 0x4b, 0xf4, 0x40,
	0x14, 0x85, 0xf7, 0x2d, 0x5e, 0xc1, 0xc1, 0x6d, 0xa6, 0xd8, 0x62, 0xb1, 0x90, 0xc5, 0xc6, 0x66,
	0x82, 0xbb, 0x68, 0x63, 0x97, 0xf5, 0xa3, 0xb1, 0x58, 0x14, 0x41, 0x2c, 0x0c, 0x93, 0xe4, 0x3a,
	0x04, 0x26, 0x33, 0x31, 0x77, 0x12, 0xb0, 0xf2, 0xc7, 0xf9, 0xc7, 0xcc, 0xc7, 0x4c, 0x74, 0x45,
	0x49, 0x22, 0xd8, 0x84, 0x30, 0x73, 0xce, 0x79, 0xee, 0x49, 0x2e, 0x79, 0x14, 0x5a, 0x0b, 0x09,
	0x4c, 0x68, 0xc9, 0x95, 0x60, 0x3a, 0x17, 0x5e, 0x24, 0x75, 0x11, 0x7b, 0x61, 0x22, 0x0c, 0x0f,
	0x25, 0x78, 0x89, 0x32, 0x90, 0x2b, 0x2e, 0xab, 0xf3, 0x02, 0xab, 0xd7, 0x00, 0x21, 0x2f, 0x93,
	0x08, 0x82, 0x2c, 0xd7, 0x46, 0x77, 0xba, 0xe0, 0xcb, 0x35, 0x6b, 0xae, 0xe9, 0x81, 0xcd, 0x77,
	0x32, 0xc6, 0xe3, 0x34, 0x51, 0xcc, 0x8a, 0x59, 0x79, 0x3c, 0xbf, 0x1f, 0x3f, 0x41, 0xcc, 0x0d,
	0xff, 0x09, 0x5f, 0xdf, 0xb5, 0xec, 0xb9, 0xf8, 0xab, 0x6e, 0x41, 0x0a, 0x88, 0x5c, 0x00, 0x5a,
	0xd0, 0xd9, 0x70, 0x10, 0xa4, 0x99, 0x79, 0x69, 0x9f, 0xad, 0x79, 0xf9, 0xf6, 0x9f, 0xcc, 0x7c,
	0xab, 0x5b, 0xb7, 0x9c, 0xdb, 0x16, 0x43, 0x4b, 0xb2, 0x7b, 0x9d, 0xa0, 0x79, 0xd0, 0x0a, 0x90,
	0x2e, 0x59, 0xdf, 0xa7, 0x64, 0x9d, 0xf8, 0x06, 0x9e, 0x0b, 0x40, 0x33, 0x5f, 0x8d, 0xf2, 0x60,
	0xa6, 0x15, 0xc2, 0x62, 0x42, 0x15, 0x21, 0x57, 0x60, 0xec, 0x30, 0x74, 0x40, 0xc8, 0x87, 0xda,
	0x91, 0x8f, 0xfa, 0x4d, 0xd6, 0x51, 0xf1, 0x5e, 0xc9, 0x5e, 0x3d, 0x86, 0x3d, 0x40, 0x7a, 0x32,
	0x6c, 0x6c, 0xa7, 0x77, 0xcc, 0xd3, 0xb1, 0xb6, 0xae, 0xb0, 0x21, 0xd3, 0x75, 0x0e, 0xdc, 0xb8,
	0x1f, 0x40, 0x07, 0x44, 0x6d, 0x19, 0x7e, 0x55, 0x5b, 0x90, 0xe9, 0x5d, 0x16, 0x7f, 0xa2, 0x0e,
	0x77, 0x8f, 0x03, 0x71, 0x32, 0x3d, 0x07, 0x09, 0xa3, 0xea, 0x6d, 0x19, 0x5c, 0xbd, 0x99, 0xf3,
	0x35, 0xab, 0x1b, 0x16, 0x4f, 0xec, 0xa2, 0xde, 0xe4, 0xc5, 0xc4, 0xbf, 0x24, 0x87, 0x91, 0x4e,
	0x7b, 0x63, 0xfd, 0xfd, 0xef, 0x57, 0x1d, 0x37, 0x75, 0xe0, 0xe6, 0x5f, 0xb8, 0xd3, 0x24, 0xaf,
	0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0x37, 0xee, 0x22, 0xa1, 0x98, 0x04, 0x00, 0x00,
}
