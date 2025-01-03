# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/inventory_v2/v1/metric_example.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from spaceone.api.core.v2 import query_pb2 as spaceone_dot_api_dot_core_dot_v2_dot_query__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n1spaceone/api/inventory_v2/v1/metric_example.proto\x12\x1cspaceone.api.inventory_v2.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x8e\x01\n\x1a\x43reateMetricExampleRequest\x12\x11\n\tmetric_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"\x8f\x01\n\x1aUpdateMetricExampleRequest\x12\x12\n\nexample_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\"*\n\x14MetricExampleRequest\x12\x12\n\nexample_id\x18\x01 \x01(\t\"\x8b\x01\n\x12MetricExampleQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x12\n\nexample_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x11\n\tmetric_id\x18\x04 \x01(\t\x12\x14\n\x0cnamespace_id\x18\x05 \x01(\t\"\xfb\x01\n\x11MetricExampleInfo\x12\x12\n\nexample_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12%\n\x04tags\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07user_id\x18\x16 \x01(\t\x12\x11\n\tmetric_id\x18\x17 \x01(\t\x12\x14\n\x0cnamespace_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"k\n\x12MetricExamplesInfo\x12@\n\x07results\x18\x01 \x03(\x0b\x32/.spaceone.api.inventory_v2.v1.MetricExampleInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"N\n\x16MetricExampleStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xaf\x07\n\rMetricExample\x12\xa6\x01\n\x06\x63reate\x12\x38.spaceone.api.inventory_v2.v1.CreateMetricExampleRequest\x1a/.spaceone.api.inventory_v2.v1.MetricExampleInfo\"1\x82\xd3\xe4\x93\x02+\"&/inventory-v2/v1/metric-example/create:\x01*\x12\xa6\x01\n\x06update\x12\x38.spaceone.api.inventory_v2.v1.UpdateMetricExampleRequest\x1a/.spaceone.api.inventory_v2.v1.MetricExampleInfo\"1\x82\xd3\xe4\x93\x02+\"&/inventory-v2/v1/metric-example/update:\x01*\x12\x87\x01\n\x06\x64\x65lete\x12\x32.spaceone.api.inventory_v2.v1.MetricExampleRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/inventory-v2/v1/metric-example/delete:\x01*\x12\x9a\x01\n\x03get\x12\x32.spaceone.api.inventory_v2.v1.MetricExampleRequest\x1a/.spaceone.api.inventory_v2.v1.MetricExampleInfo\".\x82\xd3\xe4\x93\x02(\"#/inventory-v2/v1/metric-example/get:\x01*\x12\x9b\x01\n\x04list\x12\x30.spaceone.api.inventory_v2.v1.MetricExampleQuery\x1a\x30.spaceone.api.inventory_v2.v1.MetricExamplesInfo\"/\x82\xd3\xe4\x93\x02)\"$/inventory-v2/v1/metric-example/list:\x01*\x12\x86\x01\n\x04stat\x12\x34.spaceone.api.inventory_v2.v1.MetricExampleStatQuery\x1a\x17.google.protobuf.Struct\"/\x82\xd3\xe4\x93\x02)\"$/inventory-v2/v1/metric-example/stat:\x01*BCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/inventory_v2/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.inventory_v2.v1.metric_example_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/inventory_v2/v1'
  _globals['_METRICEXAMPLE'].methods_by_name['create']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002+\"&/inventory-v2/v1/metric-example/create:\001*'
  _globals['_METRICEXAMPLE'].methods_by_name['update']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002+\"&/inventory-v2/v1/metric-example/update:\001*'
  _globals['_METRICEXAMPLE'].methods_by_name['delete']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002+\"&/inventory-v2/v1/metric-example/delete:\001*'
  _globals['_METRICEXAMPLE'].methods_by_name['get']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002(\"#/inventory-v2/v1/metric-example/get:\001*'
  _globals['_METRICEXAMPLE'].methods_by_name['list']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002)\"$/inventory-v2/v1/metric-example/list:\001*'
  _globals['_METRICEXAMPLE'].methods_by_name['stat']._loaded_options = None
  _globals['_METRICEXAMPLE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002)\"$/inventory-v2/v1/metric-example/stat:\001*'
  _globals['_CREATEMETRICEXAMPLEREQUEST']._serialized_start=207
  _globals['_CREATEMETRICEXAMPLEREQUEST']._serialized_end=349
  _globals['_UPDATEMETRICEXAMPLEREQUEST']._serialized_start=352
  _globals['_UPDATEMETRICEXAMPLEREQUEST']._serialized_end=495
  _globals['_METRICEXAMPLEREQUEST']._serialized_start=497
  _globals['_METRICEXAMPLEREQUEST']._serialized_end=539
  _globals['_METRICEXAMPLEQUERY']._serialized_start=542
  _globals['_METRICEXAMPLEQUERY']._serialized_end=681
  _globals['_METRICEXAMPLEINFO']._serialized_start=684
  _globals['_METRICEXAMPLEINFO']._serialized_end=935
  _globals['_METRICEXAMPLESINFO']._serialized_start=937
  _globals['_METRICEXAMPLESINFO']._serialized_end=1044
  _globals['_METRICEXAMPLESTATQUERY']._serialized_start=1046
  _globals['_METRICEXAMPLESTATQUERY']._serialized_end=1124
  _globals['_METRICEXAMPLE']._serialized_start=1127
  _globals['_METRICEXAMPLE']._serialized_end=2070
# @@protoc_insertion_point(module_scope)
