# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/plugin/metric.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/monitoring/plugin/metric.proto\x12\x1espaceone.api.monitoring.plugin\x1a\x1cgoogle/protobuf/struct.proto\"\x9f\x01\n\rMetricRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12&\n\x05query\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06schema\x18\x04 \x01(\t\"\xf4\x01\n\x11MetricDataRequest\x12(\n\x07options\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct\x12,\n\x0bsecret_data\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0cmetric_query\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x0e\n\x06metric\x18\x04 \x01(\t\x12\r\n\x05start\x18\x05 \x01(\t\x12\x0b\n\x03\x65nd\x18\x06 \x01(\t\x12\x0e\n\x06period\x18\x07 \x01(\x05\x12\x0c\n\x04stat\x18\x08 \x01(\t\x12\x0e\n\x06schema\x18\t \x01(\t\"\x8c\x01\n\nMetricInfo\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05group\x18\x03 \x01(\t\x12%\n\x04unit\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12-\n\x0cmetric_query\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\"J\n\x0bMetricsInfo\x12;\n\x07metrics\x18\x01 \x03(\x0b\x32*.spaceone.api.monitoring.plugin.MetricInfo\"e\n\x0eMetricDataInfo\x12*\n\x06labels\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.ListValue\x12\'\n\x06values\x18\x02 \x01(\x0b\x32\x17.google.protobuf.Struct2\xdf\x01\n\x06Metric\x12\x64\n\x04list\x12-.spaceone.api.monitoring.plugin.MetricRequest\x1a+.spaceone.api.monitoring.plugin.MetricsInfo\"\x00\x12o\n\x08get_data\x12\x31.spaceone.api.monitoring.plugin.MetricDataRequest\x1a..spaceone.api.monitoring.plugin.MetricDataInfo\"\x00\x42\x45ZCgithub.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/pluginb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.monitoring.plugin.metric_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZCgithub.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/plugin'
  _globals['_METRICREQUEST']._serialized_start=110
  _globals['_METRICREQUEST']._serialized_end=269
  _globals['_METRICDATAREQUEST']._serialized_start=272
  _globals['_METRICDATAREQUEST']._serialized_end=516
  _globals['_METRICINFO']._serialized_start=519
  _globals['_METRICINFO']._serialized_end=659
  _globals['_METRICSINFO']._serialized_start=661
  _globals['_METRICSINFO']._serialized_end=735
  _globals['_METRICDATAINFO']._serialized_start=737
  _globals['_METRICDATAINFO']._serialized_end=838
  _globals['_METRIC']._serialized_start=841
  _globals['_METRIC']._serialized_end=1064
# @@protoc_insertion_point(module_scope)
