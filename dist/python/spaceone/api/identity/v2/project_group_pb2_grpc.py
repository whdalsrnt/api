# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.identity.v2 import project_group_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2

GRPC_GENERATED_VERSION = '1.65.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.66.0'
SCHEDULED_RELEASE_DATE = 'August 6, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in spaceone/api/identity/v2/project_group_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class ProjectGroupStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.create = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/create',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.update = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/update',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.change_parent_group = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/change_parent_group',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.delete = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/delete',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.add_users = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/add_users',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.remove_users = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/remove_users',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.get = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/get',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
                _registered_method=True)
        self.list = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/list',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.FromString,
                _registered_method=True)
        self.stat = channel.unary_unary(
                '/spaceone.api.identity.v2.ProjectGroup/stat',
                request_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                _registered_method=True)


class ProjectGroupServicer(object):
    """Missing associated documentation comment in .proto file."""

    def create(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def change_parent_group(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def add_users(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def remove_users(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ProjectGroupServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'create': grpc.unary_unary_rpc_method_handler(
                    servicer.create,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'change_parent_group': grpc.unary_unary_rpc_method_handler(
                    servicer.change_parent_group,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'add_users': grpc.unary_unary_rpc_method_handler(
                    servicer.add_users,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'remove_users': grpc.unary_unary_rpc_method_handler(
                    servicer.remove_users,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.identity.v2.ProjectGroup', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.identity.v2.ProjectGroup', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class ProjectGroup(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/create',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.CreateProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/update',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UpdateProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def change_parent_group(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/change_parent_group',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ChangeParentGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/delete',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def add_users(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/add_users',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def remove_users(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/remove_users',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.UsersProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/get',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupRequest.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def list(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/list',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupSearchQuery.SerializeToString,
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupsInfo.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def stat(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/spaceone.api.identity.v2.ProjectGroup/stat',
            spaceone_dot_api_dot_identity_dot_v2_dot_project__group__pb2.ProjectGroupStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
