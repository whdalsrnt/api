# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from spaceone.api.plugin.v1 import plugin_pb2 as spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2

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
        + f' but the generated code in spaceone/api/plugin/v1/plugin_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class PluginStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.get_plugin_endpoint = channel.unary_unary(
                '/spaceone.api.plugin.v1.Plugin/get_plugin_endpoint',
                request_serializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpointRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpoint.FromString,
                _registered_method=True)
        self.get_plugin_metadata = channel.unary_unary(
                '/spaceone.api.plugin.v1.Plugin/get_plugin_metadata',
                request_serializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadataRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadata.FromString,
                _registered_method=True)
        self.notify_failure = channel.unary_unary(
                '/spaceone.api.plugin.v1.Plugin/notify_failure',
                request_serializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginFailureRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)


class PluginServicer(object):
    """Missing associated documentation comment in .proto file."""

    def get_plugin_endpoint(self, request, context):
        """Gets the `endpoint` of a specific plugin instance. A Plugin returns only a single `endpoint` by determining `labels` and `priority`. If the requested plugin instance is already deployed, the `endpoint` is returned. If not, the `endpoint` is returned after deploying the plugin instance.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_plugin_metadata(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def notify_failure(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PluginServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'get_plugin_endpoint': grpc.unary_unary_rpc_method_handler(
                    servicer.get_plugin_endpoint,
                    request_deserializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpointRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpoint.SerializeToString,
            ),
            'get_plugin_metadata': grpc.unary_unary_rpc_method_handler(
                    servicer.get_plugin_metadata,
                    request_deserializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadataRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadata.SerializeToString,
            ),
            'notify_failure': grpc.unary_unary_rpc_method_handler(
                    servicer.notify_failure,
                    request_deserializer=spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginFailureRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.plugin.v1.Plugin', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('spaceone.api.plugin.v1.Plugin', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Plugin(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def get_plugin_endpoint(request,
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
            '/spaceone.api.plugin.v1.Plugin/get_plugin_endpoint',
            spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpointRequest.SerializeToString,
            spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginEndpoint.FromString,
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
    def get_plugin_metadata(request,
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
            '/spaceone.api.plugin.v1.Plugin/get_plugin_metadata',
            spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadataRequest.SerializeToString,
            spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginMetadata.FromString,
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
    def notify_failure(request,
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
            '/spaceone.api.plugin.v1.Plugin/notify_failure',
            spaceone_dot_api_dot_plugin_dot_v1_dot_plugin__pb2.PluginFailureRequest.SerializeToString,
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
