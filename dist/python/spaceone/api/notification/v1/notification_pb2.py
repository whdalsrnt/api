# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/notification/v1/notification.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/spaceone/api/notification/v1/notification.proto\x12\x1cspaceone.api.notification.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"\xaf\x04\n\x19\x43reateNotificationRequest\x12\x15\n\rresource_type\x18\x01 \x01(\t\x12\x13\n\x0bresource_id\x18\x02 \x01(\t\x12\r\n\x05topic\x18\x03 \x01(\t\x12(\n\x07message\x18\x04 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x63\n\x11notification_type\x18\x05 \x01(\x0e\x32H.spaceone.api.notification.v1.CreateNotificationRequest.NotificationType\x12\x65\n\x12notification_level\x18\x07 \x01(\x0e\x32I.spaceone.api.notification.v1.CreateNotificationRequest.NotificationLEVEL\"]\n\x10NotificationType\x12\x1a\n\x16NOTIFICATION_TYPE_NONE\x10\x00\x12\x08\n\x04INFO\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\x12\x0b\n\x07WARNING\x10\x04\"\x81\x01\n\x11NotificationLEVEL\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03LV1\x10\x02\x12\x07\n\x03LV2\x10\x03\x12\x07\n\x03LV3\x10\x04\x12\x07\n\x03LV4\x10\x05\x12\x07\n\x03LV5\x10\x06\x12\x07\n\x03LV6\x10\x07\x12\x07\n\x03LV7\x10\x08\x12\x07\n\x03LV8\x10\t\x12\x07\n\x03LV9\x10\n\x12\x08\n\x04LV10\x10\x0b\"\xac\x01\n\x17PushNotificationRequest\x12\x13\n\x0bprotocol_id\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x02 \x01(\t\x12(\n\x07message\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x19\n\x11notification_type\x18\x04 \x01(\t\x12\x1a\n\x12notification_level\x18\x05 \x01(\t\x12\r\n\x05topic\x18\x06 \x01(\t\"-\n\x14NotificationsRequest\x12\x15\n\rnotifications\x18\x01 \x03(\t\"C\n\x16GetNotificationRequest\x12\x17\n\x0fnotification_id\x18\x01 \x01(\t\x12\x10\n\x08set_read\x18\x02 \x01(\x08\"\x9c\x04\n\x11NotificationQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x17\n\x0fnotification_id\x18\x02 \x01(\t\x12\r\n\x05topic\x18\x03 \x01(\t\x12[\n\x11notification_type\x18\x04 \x01(\x0e\x32@.spaceone.api.notification.v1.NotificationQuery.NotificationType\x12]\n\x12notification_level\x18\x05 \x01(\x0e\x32\x41.spaceone.api.notification.v1.NotificationQuery.NotificationLEVEL\x12\x0f\n\x07is_read\x18\x06 \x01(\x08\x12\x12\n\nproject_id\x18\x15 \x01(\t\x12\x1e\n\x16parent_notification_id\x18\x16 \x01(\t\"]\n\x10NotificationType\x12\x1a\n\x16NOTIFICATION_TYPE_NONE\x10\x00\x12\x08\n\x04INFO\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\x12\x0b\n\x07WARNING\x10\x04\"S\n\x11NotificationLEVEL\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03LV1\x10\x02\x12\x07\n\x03LV2\x10\x03\x12\x07\n\x03LV3\x10\x04\x12\x07\n\x03LV4\x10\x05\x12\x07\n\x03LV5\x10\x06\"\x9b\x04\n\x10NotificationInfo\x12\x17\n\x0fnotification_id\x18\x01 \x01(\t\x12\r\n\x05topic\x18\x02 \x01(\t\x12(\n\x07message\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\x12Z\n\x11notification_type\x18\x04 \x01(\x0e\x32?.spaceone.api.notification.v1.NotificationInfo.NotificationType\x12\\\n\x12notification_level\x18\x05 \x01(\x0e\x32@.spaceone.api.notification.v1.NotificationInfo.NotificationLEVEL\x12\x0f\n\x07is_read\x18\x06 \x01(\x08\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x0f\n\x07user_id\x18\x16 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\"]\n\x10NotificationType\x12\x1a\n\x16NOTIFICATION_TYPE_NONE\x10\x00\x12\x08\n\x04INFO\x10\x01\x12\t\n\x05\x45RROR\x10\x02\x12\x0b\n\x07SUCCESS\x10\x03\x12\x0b\n\x07WARNING\x10\x04\"S\n\x11NotificationLEVEL\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41LL\x10\x01\x12\x07\n\x03LV1\x10\x02\x12\x07\n\x03LV2\x10\x03\x12\x07\n\x03LV3\x10\x04\x12\x07\n\x03LV4\x10\x05\x12\x07\n\x03LV5\x10\x06\"i\n\x11NotificationsInfo\x12?\n\x07results\x18\x01 \x03(\x0b\x32..spaceone.api.notification.v1.NotificationInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"M\n\x15NotificationStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery2\x94\x07\n\x0cNotification\x12[\n\x06\x63reate\x12\x37.spaceone.api.notification.v1.CreateNotificationRequest\x1a\x16.google.protobuf.Empty\"\x00\x12W\n\x04push\x12\x35.spaceone.api.notification.v1.PushNotificationRequest\x1a\x16.google.protobuf.Empty\"\x00\x12\x85\x01\n\x06\x64\x65lete\x12\x32.spaceone.api.notification.v1.NotificationsRequest\x1a\x16.google.protobuf.Empty\"/\x82\xd3\xe4\x93\x02)\"$/notification/v1/notification/delete:\x01*\x12\x89\x01\n\x08set_read\x12\x32.spaceone.api.notification.v1.NotificationsRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\"&/notification/v1/notification/set-read:\x01*\x12\x99\x01\n\x03get\x12\x34.spaceone.api.notification.v1.GetNotificationRequest\x1a..spaceone.api.notification.v1.NotificationInfo\",\x82\xd3\xe4\x93\x02&\"!/notification/v1/notification/get:\x01*\x12\x97\x01\n\x04list\x12/.spaceone.api.notification.v1.NotificationQuery\x1a/.spaceone.api.notification.v1.NotificationsInfo\"-\x82\xd3\xe4\x93\x02\'\"\"/notification/v1/notification/list:\x01*\x12\x83\x01\n\x04stat\x12\x33.spaceone.api.notification.v1.NotificationStatQuery\x1a\x17.google.protobuf.Struct\"-\x82\xd3\xe4\x93\x02\'\"\"/notification/v1/notification/stat:\x01*BCZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.notification.v1.notification_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZAgithub.com/cloudforet-io/api/dist/go/spaceone/api/notification/v1'
  _globals['_NOTIFICATION'].methods_by_name['delete']._loaded_options = None
  _globals['_NOTIFICATION'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002)\"$/notification/v1/notification/delete:\001*'
  _globals['_NOTIFICATION'].methods_by_name['set_read']._loaded_options = None
  _globals['_NOTIFICATION'].methods_by_name['set_read']._serialized_options = b'\202\323\344\223\002+\"&/notification/v1/notification/set-read:\001*'
  _globals['_NOTIFICATION'].methods_by_name['get']._loaded_options = None
  _globals['_NOTIFICATION'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002&\"!/notification/v1/notification/get:\001*'
  _globals['_NOTIFICATION'].methods_by_name['list']._loaded_options = None
  _globals['_NOTIFICATION'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\'\"\"/notification/v1/notification/list:\001*'
  _globals['_NOTIFICATION'].methods_by_name['stat']._loaded_options = None
  _globals['_NOTIFICATION'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\'\"\"/notification/v1/notification/stat:\001*'
  _globals['_CREATENOTIFICATIONREQUEST']._serialized_start=205
  _globals['_CREATENOTIFICATIONREQUEST']._serialized_end=764
  _globals['_CREATENOTIFICATIONREQUEST_NOTIFICATIONTYPE']._serialized_start=539
  _globals['_CREATENOTIFICATIONREQUEST_NOTIFICATIONTYPE']._serialized_end=632
  _globals['_CREATENOTIFICATIONREQUEST_NOTIFICATIONLEVEL']._serialized_start=635
  _globals['_CREATENOTIFICATIONREQUEST_NOTIFICATIONLEVEL']._serialized_end=764
  _globals['_PUSHNOTIFICATIONREQUEST']._serialized_start=767
  _globals['_PUSHNOTIFICATIONREQUEST']._serialized_end=939
  _globals['_NOTIFICATIONSREQUEST']._serialized_start=941
  _globals['_NOTIFICATIONSREQUEST']._serialized_end=986
  _globals['_GETNOTIFICATIONREQUEST']._serialized_start=988
  _globals['_GETNOTIFICATIONREQUEST']._serialized_end=1055
  _globals['_NOTIFICATIONQUERY']._serialized_start=1058
  _globals['_NOTIFICATIONQUERY']._serialized_end=1598
  _globals['_NOTIFICATIONQUERY_NOTIFICATIONTYPE']._serialized_start=539
  _globals['_NOTIFICATIONQUERY_NOTIFICATIONTYPE']._serialized_end=632
  _globals['_NOTIFICATIONQUERY_NOTIFICATIONLEVEL']._serialized_start=635
  _globals['_NOTIFICATIONQUERY_NOTIFICATIONLEVEL']._serialized_end=718
  _globals['_NOTIFICATIONINFO']._serialized_start=1601
  _globals['_NOTIFICATIONINFO']._serialized_end=2140
  _globals['_NOTIFICATIONINFO_NOTIFICATIONTYPE']._serialized_start=539
  _globals['_NOTIFICATIONINFO_NOTIFICATIONTYPE']._serialized_end=632
  _globals['_NOTIFICATIONINFO_NOTIFICATIONLEVEL']._serialized_start=635
  _globals['_NOTIFICATIONINFO_NOTIFICATIONLEVEL']._serialized_end=718
  _globals['_NOTIFICATIONSINFO']._serialized_start=2142
  _globals['_NOTIFICATIONSINFO']._serialized_end=2247
  _globals['_NOTIFICATIONSTATQUERY']._serialized_start=2249
  _globals['_NOTIFICATIONSTATQUERY']._serialized_end=2326
  _globals['_NOTIFICATION']._serialized_start=2329
  _globals['_NOTIFICATION']._serialized_end=3245
# @@protoc_insertion_point(module_scope)
