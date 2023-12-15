# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.file_manager.v1 import public_file_pb2 as spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2


class PublicFileStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.add = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/add',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.CreatePublicFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
                )
        self.update = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/update',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.UpdatePublicFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
                )
        self.delete = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/delete',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get_download_url = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/get_download_url',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/get',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/list',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileSearchQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFilesInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.file_manager.v1.PublicFile/stat',
                request_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class PublicFileServicer(object):
    """Missing associated documentation comment in .proto file."""

    def add(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def update(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def delete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get_download_url(self, request, context):
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


def add_PublicFileServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'add': grpc.unary_unary_rpc_method_handler(
                    servicer.add,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.CreatePublicFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.SerializeToString,
            ),
            'update': grpc.unary_unary_rpc_method_handler(
                    servicer.update,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.UpdatePublicFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.SerializeToString,
            ),
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get_download_url': grpc.unary_unary_rpc_method_handler(
                    servicer.get_download_url,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileSearchQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFilesInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.file_manager.v1.PublicFile', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PublicFile(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def add(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/add',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.CreatePublicFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/update',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.UpdatePublicFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/delete',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def get_download_url(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/get_download_url',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/get',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileRequest.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/list',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileSearchQuery.SerializeToString,
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFilesInfo.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.file_manager.v1.PublicFile/stat',
            spaceone_dot_api_dot_file__manager_dot_v1_dot_public__file__pb2.PublicFileStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)