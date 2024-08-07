# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/monitoring/v1/escalation_policy.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n2spaceone/api/monitoring/v1/escalation_policy.proto\x12\x1aspaceone.api.monitoring.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\x94\x02\n\x14\x45scalationPolicyRule\x12^\n\x12notification_level\x18\x01 \x01(\x0e\x32\x42.spaceone.api.monitoring.v1.EscalationPolicyRule.NotificationLevel\x12\x18\n\x10\x65scalate_minutes\x18\x02 \x01(\x05\"\x81\x01\n\x11NotificationLevel\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03LV1\x10\x02\x12\x07\n\x03LV2\x10\x03\x12\x07\n\x03LV3\x10\x04\x12\x07\n\x03LV4\x10\x05\x12\x07\n\x03LV5\x10\x06\x12\x07\n\x03LV6\x10\x07\x12\x07\n\x03LV7\x10\x08\x12\x07\n\x03LV8\x10\t\x12\x07\n\x03LV9\x10\n\x12\x08\n\x04LV10\x10\x0b\"\x9c\x04\n\x1d\x43reateEscalationPolicyRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12?\n\x05rules\x18\x02 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x03 \x01(\x05\x12m\n\x10\x66inish_condition\x18\x04 \x01(\x0e\x32S.spaceone.api.monitoring.v1.CreateEscalationPolicyRequest.EscalationFinishCondition\x12%\n\x04tags\x18\x05 \x01(\x0b\x32\x17.google.protobuf.Struct\x12_\n\x0eresource_group\x18\x14 \x01(\x0e\x32G.spaceone.api.monitoring.v1.CreateEscalationPolicyRequest.ResourceGroup\x12\x12\n\nproject_id\x18\x15 \x01(\t\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"D\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\r\n\tWORKSPACE\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"\xff\x02\n\x1dUpdateEscalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12?\n\x05rules\x18\x03 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x04 \x01(\x05\x12m\n\x10\x66inish_condition\x18\x05 \x01(\x0e\x32S.spaceone.api.monitoring.v1.UpdateEscalationPolicyRequest.EscalationFinishCondition\x12%\n\x04tags\x18\x06 \x01(\x0b\x32\x17.google.protobuf.Struct\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"7\n\x17\x45scalationPolicyRequest\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\"\xe4\x03\n\x15\x45scalationPolicyQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x1c\n\x14\x65scalation_policy_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x12\n\nis_default\x18\x04 \x01(\x08\x12\x65\n\x10\x66inish_condition\x18\x05 \x01(\x0e\x32K.spaceone.api.monitoring.v1.EscalationPolicyQuery.EscalationFinishCondition\x12W\n\x0eresource_group\x18\x06 \x01(\x0e\x32?.spaceone.api.monitoring.v1.EscalationPolicyQuery.ResourceGroup\x12\x12\n\nproject_id\x18\x07 \x01(\t\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"D\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\r\n\tWORKSPACE\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"\xf0\x04\n\x14\x45scalationPolicyInfo\x12\x1c\n\x14\x65scalation_policy_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nis_default\x18\x03 \x01(\x08\x12?\n\x05rules\x18\x04 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyRule\x12\x14\n\x0crepeat_count\x18\x05 \x01(\x05\x12\x64\n\x10\x66inish_condition\x18\x06 \x01(\x0e\x32J.spaceone.api.monitoring.v1.EscalationPolicyInfo.EscalationFinishCondition\x12V\n\x0eresource_group\x18\x07 \x01(\x0e\x32>.spaceone.api.monitoring.v1.EscalationPolicyInfo.ResourceGroup\x12\x12\n\nproject_id\x18\x08 \x01(\t\x12%\n\x04tags\x18\t \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x14\n\x0cworkspace_id\x18\n \x01(\t\x12\x11\n\tdomain_id\x18\x0b \x01(\t\x12\x12\n\ncreated_at\x18\x15 \x01(\t\"E\n\x19\x45scalationFinishCondition\x12\x08\n\x04NONE\x10\x00\x12\x10\n\x0c\x41\x43KNOWLEDGED\x10\x01\x12\x0c\n\x08RESOLVED\x10\x02\"D\n\rResourceGroup\x12\x17\n\x13RESOURCE_GROUP_NONE\x10\x00\x12\r\n\tWORKSPACE\x10\x01\x12\x0b\n\x07PROJECT\x10\x02\"p\n\x16\x45scalationPoliciesInfo\x12\x41\n\x07results\x18\x01 \x03(\x0b\x32\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"Q\n\x19\x45scalationPolicyStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\xf3\x08\n\x10\x45scalationPolicy\x12\xa9\x01\n\x06\x63reate\x12\x39.spaceone.api.monitoring.v1.CreateEscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/create:\x01*\x12\xa9\x01\n\x06update\x12\x39.spaceone.api.monitoring.v1.UpdateEscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/update:\x01*\x12\xad\x01\n\x0bset_default\x12\x33.spaceone.api.monitoring.v1.EscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"7\x82\xd3\xe4\x93\x02\x31\",/monitoring/v1/escalation-policy/set-default:\x01*\x12\x89\x01\n\x06\x64\x65lete\x12\x33.spaceone.api.monitoring.v1.EscalationPolicyRequest\x1a\x16.google.protobuf.Empty\"2\x82\xd3\xe4\x93\x02,\"\'/monitoring/v1/escalation-policy/delete:\x01*\x12\x9d\x01\n\x03get\x12\x33.spaceone.api.monitoring.v1.EscalationPolicyRequest\x1a\x30.spaceone.api.monitoring.v1.EscalationPolicyInfo\"/\x82\xd3\xe4\x93\x02)\"$/monitoring/v1/escalation-policy/get:\x01*\x12\x9f\x01\n\x04list\x12\x31.spaceone.api.monitoring.v1.EscalationPolicyQuery\x1a\x32.spaceone.api.monitoring.v1.EscalationPoliciesInfo\"0\x82\xd3\xe4\x93\x02*\"%/monitoring/v1/escalation-policy/list:\x01*\x12\x88\x01\n\x04stat\x12\x35.spaceone.api.monitoring.v1.EscalationPolicyStatQuery\x1a\x17.google.protobuf.Struct\"0\x82\xd3\xe4\x93\x02*\"%/monitoring/v1/escalation-policy/stat:\x01*BAZ?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.monitoring.v1.escalation_policy_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z?github.com/cloudforet-io/api/dist/go/spaceone/api/monitoring/v1'
  _globals['_ESCALATIONPOLICY'].methods_by_name['create']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/create:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['update']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/update:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['set_default']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['set_default']._serialized_options = b'\202\323\344\223\0021\",/monitoring/v1/escalation-policy/set-default:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['delete']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002,\"\'/monitoring/v1/escalation-policy/delete:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['get']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002)\"$/monitoring/v1/escalation-policy/get:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['list']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002*\"%/monitoring/v1/escalation-policy/list:\001*'
  _globals['_ESCALATIONPOLICY'].methods_by_name['stat']._loaded_options = None
  _globals['_ESCALATIONPOLICY'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002*\"%/monitoring/v1/escalation-policy/stat:\001*'
  _globals['_ESCALATIONPOLICYRULE']._serialized_start=206
  _globals['_ESCALATIONPOLICYRULE']._serialized_end=482
  _globals['_ESCALATIONPOLICYRULE_NOTIFICATIONLEVEL']._serialized_start=353
  _globals['_ESCALATIONPOLICYRULE_NOTIFICATIONLEVEL']._serialized_end=482
  _globals['_CREATEESCALATIONPOLICYREQUEST']._serialized_start=485
  _globals['_CREATEESCALATIONPOLICYREQUEST']._serialized_end=1025
  _globals['_CREATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION']._serialized_start=886
  _globals['_CREATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION']._serialized_end=955
  _globals['_CREATEESCALATIONPOLICYREQUEST_RESOURCEGROUP']._serialized_start=957
  _globals['_CREATEESCALATIONPOLICYREQUEST_RESOURCEGROUP']._serialized_end=1025
  _globals['_UPDATEESCALATIONPOLICYREQUEST']._serialized_start=1028
  _globals['_UPDATEESCALATIONPOLICYREQUEST']._serialized_end=1411
  _globals['_UPDATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION']._serialized_start=886
  _globals['_UPDATEESCALATIONPOLICYREQUEST_ESCALATIONFINISHCONDITION']._serialized_end=955
  _globals['_ESCALATIONPOLICYREQUEST']._serialized_start=1413
  _globals['_ESCALATIONPOLICYREQUEST']._serialized_end=1468
  _globals['_ESCALATIONPOLICYQUERY']._serialized_start=1471
  _globals['_ESCALATIONPOLICYQUERY']._serialized_end=1955
  _globals['_ESCALATIONPOLICYQUERY_ESCALATIONFINISHCONDITION']._serialized_start=886
  _globals['_ESCALATIONPOLICYQUERY_ESCALATIONFINISHCONDITION']._serialized_end=955
  _globals['_ESCALATIONPOLICYQUERY_RESOURCEGROUP']._serialized_start=957
  _globals['_ESCALATIONPOLICYQUERY_RESOURCEGROUP']._serialized_end=1025
  _globals['_ESCALATIONPOLICYINFO']._serialized_start=1958
  _globals['_ESCALATIONPOLICYINFO']._serialized_end=2582
  _globals['_ESCALATIONPOLICYINFO_ESCALATIONFINISHCONDITION']._serialized_start=886
  _globals['_ESCALATIONPOLICYINFO_ESCALATIONFINISHCONDITION']._serialized_end=955
  _globals['_ESCALATIONPOLICYINFO_RESOURCEGROUP']._serialized_start=957
  _globals['_ESCALATIONPOLICYINFO_RESOURCEGROUP']._serialized_end=1025
  _globals['_ESCALATIONPOLICIESINFO']._serialized_start=2584
  _globals['_ESCALATIONPOLICIESINFO']._serialized_end=2696
  _globals['_ESCALATIONPOLICYSTATQUERY']._serialized_start=2698
  _globals['_ESCALATIONPOLICYSTATQUERY']._serialized_end=2779
  _globals['_ESCALATIONPOLICY']._serialized_start=2782
  _globals['_ESCALATIONPOLICY']._serialized_end=3921
# @@protoc_insertion_point(module_scope)
