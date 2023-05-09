# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from spaceone.api.inventory.v1 import collector_pb2 as spaceone_dot_api_dot_inventory_dot_v1_dot_collector__pb2
from spaceone.api.inventory.v1 import job_pb2 as spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2


class JobStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.delete = channel.unary_unary(
                '/spaceone.api.inventory.v1.Job/delete',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.get = channel.unary_unary(
                '/spaceone.api.inventory.v1.Job/get',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.GetJobRequest.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_collector__pb2.JobInfo.FromString,
                )
        self.list = channel.unary_unary(
                '/spaceone.api.inventory.v1.Job/list',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsQuery.SerializeToString,
                response_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsInfo.FromString,
                )
        self.stat = channel.unary_unary(
                '/spaceone.api.inventory.v1.Job/stat',
                request_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobStatQuery.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_struct__pb2.Struct.FromString,
                )


class JobServicer(object):
    """Missing associated documentation comment in .proto file."""

    def delete(self, request, context):
        """
        desc: Deletes a specific Job. You must specify the `job_id` of the Job to delete.
        request_example: >-
        {
        "job_id": "job-123456789012",
        "domain_id": "domain-123456789012"
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def get(self, request, context):
        """
        desc: Gets a specific Job. Prints detailed information about the Job, including its state, total tasks, and collector info.
        request_example: >-
        {
        "job_id": "job-123456789012",
        "domain_id": "domain-123456789012"
        }
        response_example: >-
        {
        "job_id": "job-123456789012",
        "status": "ERROR",
        "filter": {},
        "total_tasks": 2,
        "collector_info": {
        "collector_id": "collector-123456789012",
        "name": "Jenkins Collector",
        "state": "ENABLED",
        "plugin_info": {
        "plugin_id": "plugin-jenkins-inven-collector",
        "version": "0.1.1"
        },
        "provider": "jenkins",
        "capability": {},
        "is_public": true
        },
        "domain_id": "domain-123456789012",
        "created_at": "2022-01-01T10:00:01.389Z",
        "updated_at": "2022-01-01T10:00:01.389Z",
        "finished_at": "2022-01-01T10:02:11.270Z"
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def list(self, request, context):
        """
        desc: Gets a list of all Jobs. You can use a query to get a filtered list of Jobs.
        request_example: >-
        {
        "query": {}
        }
        response_example: >-
        {
        "results": [
        {
        "job_id": "job-3b124006c2d2",
        "status": "SUCCESS",
        "filter": {},
        "total_tasks": 2,
        "collector_info": {
        "collector_id": "collector-accd02663b3d",
        "name": "openstack-collector",
        "state": "ENABLED",
        "plugin_info": {
        "plugin_id": "plugin-openstack-inven-collector",
        "version": "0.4.2.20220616.134758"
        },
        "provider": "openstack",
        "capability": {
        "supported_schema": [
        "openstack_credentials"
        ]
        },
        "is_public": true
        },
        "domain_id": "domain-58010aa2e451",
        "created_at": "2022-06-17T08:00:01.225Z",
        "updated_at": "2022-06-17T08:00:01.225Z",
        "finished_at": "2022-06-17T08:00:15.197Z"
        },
        {
        "job_id": "job-587a3d3b4db3",
        "status": "SUCCESS",
        "filter": {},
        "total_tasks": 3,
        "collector_info": {
        "collector_id": "collector-2c0847644f39",
        "name": "AWS stat-kwon Collector",
        "state": "ENABLED",
        "plugin_info": {
        "plugin_id": "plugin-30d21ef75a5d",
        "version": "1.13.13.20220610.143142"
        },
        "provider": "aws",
        "capability": {
        "supported_schema": [
        "aws_access_key"
        ]
        },
        "is_public": true
        },
        "domain_id": "domain-58010aa2e451",
        "created_at": "2022-06-17T08:00:00.407Z",
        "updated_at": "2022-06-17T08:00:00.407Z",
        "finished_at": "2022-06-17T08:07:32.023Z"
        }
        ],
        "total_count": 2
        }
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def stat(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_JobServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'delete': grpc.unary_unary_rpc_method_handler(
                    servicer.delete,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'get': grpc.unary_unary_rpc_method_handler(
                    servicer.get,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.GetJobRequest.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_collector__pb2.JobInfo.SerializeToString,
            ),
            'list': grpc.unary_unary_rpc_method_handler(
                    servicer.list,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsQuery.FromString,
                    response_serializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsInfo.SerializeToString,
            ),
            'stat': grpc.unary_unary_rpc_method_handler(
                    servicer.stat,
                    request_deserializer=spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobStatQuery.FromString,
                    response_serializer=google_dot_protobuf_dot_struct__pb2.Struct.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'spaceone.api.inventory.v1.Job', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Job(object):
    """Missing associated documentation comment in .proto file."""

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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.Job/delete',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.Job/get',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.GetJobRequest.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_collector__pb2.JobInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.Job/list',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsQuery.SerializeToString,
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobsInfo.FromString,
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
        return grpc.experimental.unary_unary(request, target, '/spaceone.api.inventory.v1.Job/stat',
            spaceone_dot_api_dot_inventory_dot_v1_dot_job__pb2.JobStatQuery.SerializeToString,
            google_dot_protobuf_dot_struct__pb2.Struct.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)