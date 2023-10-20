# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/plugin/v1/plugin.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/plugin/v1/plugin.proto\x12\x16spaceone.api.plugin.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\"\xed\x01\n\x15PluginEndpointRequest\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12\'\n\x06labels\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12O\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32\x39.spaceone.api.plugin.v1.PluginEndpointRequest.UpgradeMode\"#\n\x0bUpgradeMode\x12\n\n\x06MANUAL\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\"\xee\x01\n\x15PluginMetadataRequest\x12\x11\n\tplugin_id\x18\x01 \x01(\t\x12\x0f\n\x07version\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x04 \x01(\t\x12O\n\x0cupgrade_mode\x18\x05 \x01(\x0e\x32\x39.spaceone.api.plugin.v1.PluginMetadataRequest.UpgradeMode\"#\n\x0bUpgradeMode\x12\n\n\x06MANUAL\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\"d\n\x14PluginFailureRequest\x12\x15\n\rsupervisor_id\x18\x01 \x01(\t\x12\x11\n\tplugin_id\x18\x02 \x01(\t\x12\x0f\n\x07version\x18\x03 \x01(\t\x12\x11\n\tdomain_id\x18\x04 \x01(\t\"Q\n\x0ePluginEndpoint\x12\x10\n\x08\x65ndpoint\x18\x01 \x01(\t\x12\x14\n\x0c\x61\x63\x63\x65ss_token\x18\x02 \x01(\t\x12\x17\n\x0fupdated_version\x18\x03 \x01(\t\";\n\x0ePluginMetadata\x12)\n\x08metadata\x18\x01 \x01(\x0b\x32\x17.google.protobuf.Struct2\xd0\x03\n\x06Plugin\x12\x9e\x01\n\x13get_plugin_endpoint\x12-.spaceone.api.plugin.v1.PluginEndpointRequest\x1a&.spaceone.api.plugin.v1.PluginEndpoint\"0\x82\xd3\xe4\x93\x02*\"%/plugin/v1/plugin/get-plugin-endpoint:\x01*\x12\x9e\x01\n\x13get_plugin_metadata\x12-.spaceone.api.plugin.v1.PluginMetadataRequest\x1a&.spaceone.api.plugin.v1.PluginMetadata\"0\x82\xd3\xe4\x93\x02*\"%/plugin/v1/plugin/get-plugin-metadata:\x01*\x12\x83\x01\n\x0enotify_failure\x12,.spaceone.api.plugin.v1.PluginFailureRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /plugin/v1/plugin/notify-failure:\x01*B=Z;github.com/cloudforet-io/api/dist/go/spaceone/api/plugin/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.plugin.v1.plugin_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z;github.com/cloudforet-io/api/dist/go/spaceone/api/plugin/v1'
  _PLUGIN.methods_by_name['get_plugin_endpoint']._options = None
  _PLUGIN.methods_by_name['get_plugin_endpoint']._serialized_options = b'\202\323\344\223\002*\"%/plugin/v1/plugin/get-plugin-endpoint:\001*'
  _PLUGIN.methods_by_name['get_plugin_metadata']._options = None
  _PLUGIN.methods_by_name['get_plugin_metadata']._serialized_options = b'\202\323\344\223\002*\"%/plugin/v1/plugin/get-plugin-metadata:\001*'
  _PLUGIN.methods_by_name['notify_failure']._options = None
  _PLUGIN.methods_by_name['notify_failure']._serialized_options = b'\202\323\344\223\002%\" /plugin/v1/plugin/notify-failure:\001*'
  _globals['_PLUGINENDPOINTREQUEST']._serialized_start=153
  _globals['_PLUGINENDPOINTREQUEST']._serialized_end=390
  _globals['_PLUGINENDPOINTREQUEST_UPGRADEMODE']._serialized_start=355
  _globals['_PLUGINENDPOINTREQUEST_UPGRADEMODE']._serialized_end=390
  _globals['_PLUGINMETADATAREQUEST']._serialized_start=393
  _globals['_PLUGINMETADATAREQUEST']._serialized_end=631
  _globals['_PLUGINMETADATAREQUEST_UPGRADEMODE']._serialized_start=355
  _globals['_PLUGINMETADATAREQUEST_UPGRADEMODE']._serialized_end=390
  _globals['_PLUGINFAILUREREQUEST']._serialized_start=633
  _globals['_PLUGINFAILUREREQUEST']._serialized_end=733
  _globals['_PLUGINENDPOINT']._serialized_start=735
  _globals['_PLUGINENDPOINT']._serialized_end=816
  _globals['_PLUGINMETADATA']._serialized_start=818
  _globals['_PLUGINMETADATA']._serialized_end=877
  _globals['_PLUGIN']._serialized_start=880
  _globals['_PLUGIN']._serialized_end=1344
# @@protoc_insertion_point(module_scope)
