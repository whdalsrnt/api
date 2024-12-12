# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/service.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n+spaceone/api/alert_manager/v1/service.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"2\n\x0eServiceMembers\x12\x0c\n\x04USER\x18\x01 \x03(\t\x12\x12\n\nUSER_GROUP\x18\x02 \x03(\t\"\xc3\x02\n\x0eServiceOptions\x12_\n\x14notification_urgency\x18\x01 \x01(\x0e\x32\x41.spaceone.api.alert_manager.v1.ServiceOptions.NotificationUrgency\x12Q\n\rrecovery_mode\x18\x02 \x01(\x0e\x32:.spaceone.api.alert_manager.v1.ServiceOptions.RecoveryMode\"?\n\x13NotificationUrgency\x12\x10\n\x0cURGENCY_NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\r\n\tHIGH_ONLY\x10\x02\"<\n\x0cRecoveryMode\x12\x16\n\x12RECOVERY_MODE_NONE\x10\x00\x12\n\n\x06MANUAL\x10\x01\x12\x08\n\x04\x41UTO\x10\x02\"\'\n\nAlertStats\x12\x0c\n\x04high\x18\x01 \x01(\x05\x12\x0b\n\x03low\x18\x02 \x01(\x05\"\xc1\x01\n\x06\x41lerts\x12\x38\n\x05TOTAL\x18\x01 \x01(\x0b\x32).spaceone.api.alert_manager.v1.AlertStats\x12<\n\tTRIGGERED\x18\x02 \x01(\x0b\x32).spaceone.api.alert_manager.v1.AlertStats\x12?\n\x0c\x41\x43KNOWLEDGED\x18\x03 \x01(\x0b\x32).spaceone.api.alert_manager.v1.AlertStats\"\xf5\x01\n\x14ServiceCreateRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0bservice_key\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12>\n\x07members\x18\x04 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceMembers\x12>\n\x07options\x18\x05 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceOptions\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\"\xd2\x01\n\x14ServiceUpdateRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12>\n\x07options\x18\x04 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceOptions\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x1c\n\x14\x65scalation_policy_id\x18\x15 \x01(\t\"q\n\x1bServiceChangeMembersRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12>\n\x07members\x18\x02 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceMembers\"$\n\x0eServiceRequest\x12\x12\n\nservice_id\x18\x01 \x01(\t\"\x80\x01\n\x12ServiceSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x12\n\nservice_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x1c\n\x14\x65scalation_policy_id\x18\x15 \x01(\t\"H\n\x10ServiceStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\xca\x03\n\x0bServiceInfo\x12\x12\n\nservice_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bservice_key\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12>\n\x07members\x18\x05 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceMembers\x12>\n\x07options\x18\x06 \x01(\x0b\x32-.spaceone.api.alert_manager.v1.ServiceOptions\x12\x10\n\x08\x63hannels\x18\x07 \x03(\t\x12\x10\n\x08webhooks\x18\x08 \x03(\t\x12\x35\n\x06\x61lerts\x18\t \x01(\x0b\x32%.spaceone.api.alert_manager.v1.Alerts\x12%\n\x04tags\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x1c\n\x14\x65scalation_policy_id\x18\x17 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\"`\n\x0cServicesInfo\x12;\n\x07results\x18\x01 \x03(\x0b\x32*.spaceone.api.alert_manager.v1.ServiceInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\x86\x08\n\x07Service\x12\x96\x01\n\x06\x63reate\x12\x33.spaceone.api.alert_manager.v1.ServiceCreateRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"+\x82\xd3\xe4\x93\x02%\" /alert-manager/v1/service/create:\x01*\x12\x96\x01\n\x06update\x12\x33.spaceone.api.alert_manager.v1.ServiceUpdateRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"+\x82\xd3\xe4\x93\x02%\" /alert-manager/v1/service/update:\x01*\x12\xac\x01\n\x0e\x63hange_members\x12:.spaceone.api.alert_manager.v1.ServiceChangeMembersRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"2\x82\xd3\xe4\x93\x02,\"\'/alert-manager/v1/service/change-member:\x01*\x12|\n\x06\x64\x65lete\x12-.spaceone.api.alert_manager.v1.ServiceRequest\x1a\x16.google.protobuf.Empty\"+\x82\xd3\xe4\x93\x02%\" /alert-manager/v1/service/delete:\x01*\x12\x8a\x01\n\x03get\x12-.spaceone.api.alert_manager.v1.ServiceRequest\x1a*.spaceone.api.alert_manager.v1.ServiceInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/alert-manager/v1/service/get:\x01*\x12\x91\x01\n\x04list\x12\x31.spaceone.api.alert_manager.v1.ServiceSearchQuery\x1a+.spaceone.api.alert_manager.v1.ServicesInfo\")\x82\xd3\xe4\x93\x02#\"\x1e/alert-manager/v1/service/list:\x01*\x12{\n\x04stat\x12/.spaceone.api.alert_manager.v1.ServiceStatQuery\x1a\x17.google.protobuf.Struct\")\x82\xd3\xe4\x93\x02#\"\x1e/alert-manager/v1/service/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.service_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_SERVICE'].methods_by_name['create']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002%\" /alert-manager/v1/service/create:\001*'
  _globals['_SERVICE'].methods_by_name['update']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002%\" /alert-manager/v1/service/update:\001*'
  _globals['_SERVICE'].methods_by_name['change_members']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['change_members']._serialized_options = b'\202\323\344\223\002,\"\'/alert-manager/v1/service/change-member:\001*'
  _globals['_SERVICE'].methods_by_name['delete']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002%\" /alert-manager/v1/service/delete:\001*'
  _globals['_SERVICE'].methods_by_name['get']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\"\"\035/alert-manager/v1/service/get:\001*'
  _globals['_SERVICE'].methods_by_name['list']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002#\"\036/alert-manager/v1/service/list:\001*'
  _globals['_SERVICE'].methods_by_name['stat']._loaded_options = None
  _globals['_SERVICE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002#\"\036/alert-manager/v1/service/stat:\001*'
  _globals['_SERVICEMEMBERS']._serialized_start=201
  _globals['_SERVICEMEMBERS']._serialized_end=251
  _globals['_SERVICEOPTIONS']._serialized_start=254
  _globals['_SERVICEOPTIONS']._serialized_end=577
  _globals['_SERVICEOPTIONS_NOTIFICATIONURGENCY']._serialized_start=452
  _globals['_SERVICEOPTIONS_NOTIFICATIONURGENCY']._serialized_end=515
  _globals['_SERVICEOPTIONS_RECOVERYMODE']._serialized_start=517
  _globals['_SERVICEOPTIONS_RECOVERYMODE']._serialized_end=577
  _globals['_ALERTSTATS']._serialized_start=579
  _globals['_ALERTSTATS']._serialized_end=618
  _globals['_ALERTS']._serialized_start=621
  _globals['_ALERTS']._serialized_end=814
  _globals['_SERVICECREATEREQUEST']._serialized_start=817
  _globals['_SERVICECREATEREQUEST']._serialized_end=1062
  _globals['_SERVICEUPDATEREQUEST']._serialized_start=1065
  _globals['_SERVICEUPDATEREQUEST']._serialized_end=1275
  _globals['_SERVICECHANGEMEMBERSREQUEST']._serialized_start=1277
  _globals['_SERVICECHANGEMEMBERSREQUEST']._serialized_end=1390
  _globals['_SERVICEREQUEST']._serialized_start=1392
  _globals['_SERVICEREQUEST']._serialized_end=1428
  _globals['_SERVICESEARCHQUERY']._serialized_start=1431
  _globals['_SERVICESEARCHQUERY']._serialized_end=1559
  _globals['_SERVICESTATQUERY']._serialized_start=1561
  _globals['_SERVICESTATQUERY']._serialized_end=1633
  _globals['_SERVICEINFO']._serialized_start=1636
  _globals['_SERVICEINFO']._serialized_end=2094
  _globals['_SERVICESINFO']._serialized_start=2096
  _globals['_SERVICESINFO']._serialized_end=2192
  _globals['_SERVICE']._serialized_start=2195
  _globals['_SERVICE']._serialized_end=3225
# @@protoc_insertion_point(module_scope)
