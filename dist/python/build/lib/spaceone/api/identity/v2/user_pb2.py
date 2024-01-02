# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/identity/v2/user.proto
# Protobuf Python Version: 4.25.0
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
from spaceone.api.identity.v2 import workspace_pb2 as spaceone_dot_api_dot_identity_dot_v2_dot_workspace__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n#spaceone/api/identity/v2/user.proto\x12\x18spaceone.api.identity.v2\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\x1a(spaceone/api/identity/v2/workspace.proto\"\xa3\x01\n\x03MFA\x12\x32\n\x05state\x18\x01 \x01(\x0e\x32#.spaceone.api.identity.v2.MFA.State\x12\x10\n\x08mfa_type\x18\x02 \x01(\t\x12(\n\x07options\x18\x03 \x01(\x0b\x32\x17.google.protobuf.Struct\",\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\"\xed\x01\n\x11\x43reateUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x35\n\tauth_type\x18\x05 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\x12\x10\n\x08language\x18\x06 \x01(\t\x12\x10\n\x08timezone\x18\x07 \x01(\t\x12%\n\x04tags\x18\x08 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0ereset_password\x18\t \x01(\x08\"\xb6\x01\n\x11UpdateUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x10\n\x08language\x18\x05 \x01(\t\x12\x10\n\x08timezone\x18\x06 \x01(\t\x12%\n\x04tags\x18\x07 \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x16\n\x0ereset_password\x18\x08 \x01(\x08\"8\n\x16VerifyUserEmailRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\r\n\x05\x65mail\x18\x02 \x01(\t\"(\n\x15\x44isableUserMFARequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"t\n\x19SetRequiredActionsRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x46\n\x10required_actions\x18\x02 \x03(\x0e\x32,.spaceone.api.identity.v2.UserRequiredAction\"\x1e\n\x0bUserRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"\x9d\x02\n\x0fUserSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x0f\n\x07user_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12>\n\x05state\x18\x04 \x01(\x0e\x32/.spaceone.api.identity.v2.UserSearchQuery.State\x12\r\n\x05\x65mail\x18\x05 \x01(\t\x12\x35\n\tauth_type\x18\x06 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\"9\n\x05State\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x0b\n\x07PENDING\x10\x03\"\x8e\x05\n\x08UserInfo\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x37\n\x05state\x18\x03 \x01(\x0e\x32(.spaceone.api.identity.v2.UserInfo.State\x12\r\n\x05\x65mail\x18\x04 \x01(\t\x12\x16\n\x0e\x65mail_verified\x18\x05 \x01(\x08\x12\x35\n\tauth_type\x18\x06 \x01(\x0e\x32\".spaceone.api.identity.v2.AuthType\x12\x0f\n\x07role_id\x18\x07 \x01(\t\x12>\n\trole_type\x18\x08 \x01(\x0e\x32+.spaceone.api.identity.v2.UserInfo.RoleType\x12*\n\x03mfa\x18\t \x01(\x0b\x32\x1d.spaceone.api.identity.v2.MFA\x12\x10\n\x08language\x18\x0b \x01(\t\x12\x10\n\x08timezone\x18\x0c \x01(\t\x12\x46\n\x10required_actions\x18\r \x03(\x0e\x32,.spaceone.api.identity.v2.UserRequiredAction\x12%\n\x04tags\x18\x0e \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x18\n\x10last_accessed_at\x18  \x01(\t\"?\n\x05State\x12\x0e\n\nSTATE_NONE\x10\x00\x12\x0b\n\x07\x45NABLED\x10\x01\x12\x0c\n\x08\x44ISABLED\x10\x02\x12\x0b\n\x07PENDING\x10\x03\":\n\x08RoleType\x12\x12\n\x0eROLE_TYPE_NONE\x10\x00\x12\x10\n\x0c\x44OMAIN_ADMIN\x10\x01\x12\x08\n\x04USER\x10\x02\"U\n\tUsersInfo\x12\x33\n\x07results\x18\x01 \x03(\x0b\x32\".spaceone.api.identity.v2.UserInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\"E\n\rUserStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery*5\n\x08\x41uthType\x12\x10\n\x0cNONE_BACKEND\x10\x00\x12\t\n\x05LOCAL\x10\x01\x12\x0c\n\x08\x45XTERNAL\x10\x02*)\n\x12UserRequiredAction\x12\x13\n\x0fUPDATE_PASSWORD\x10\x00\x32\x80\x0b\n\x04User\x12~\n\x06\x63reate\x12+.spaceone.api.identity.v2.CreateUserRequest\x1a\".spaceone.api.identity.v2.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/user/create:\x01*\x12~\n\x06update\x12+.spaceone.api.identity.v2.UpdateUserRequest\x1a\".spaceone.api.identity.v2.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/user/update:\x01*\x12\x83\x01\n\x0cverify_email\x12\x30.spaceone.api.identity.v2.VerifyUserEmailRequest\x1a\x16.google.protobuf.Empty\")\x82\xd3\xe4\x93\x02#\"\x1e/identity/v2/user/verify-email:\x01*\x12\x8c\x01\n\x0b\x64isable_mfa\x12/.spaceone.api.identity.v2.DisableUserMFARequest\x1a\".spaceone.api.identity.v2.UserInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/identity/v2/user/disable-mfa:\x01*\x12\xa2\x01\n\x14set_required_actions\x12\x33.spaceone.api.identity.v2.SetRequiredActionsRequest\x1a\".spaceone.api.identity.v2.UserInfo\"1\x82\xd3\xe4\x93\x02+\"&/identity/v2/user/set-required-actions:\x01*\x12x\n\x06\x65nable\x12%.spaceone.api.identity.v2.UserRequest\x1a\".spaceone.api.identity.v2.UserInfo\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/user/enable:\x01*\x12z\n\x07\x64isable\x12%.spaceone.api.identity.v2.UserRequest\x1a\".spaceone.api.identity.v2.UserInfo\"$\x82\xd3\xe4\x93\x02\x1e\"\x19/identity/v2/user/disable:\x01*\x12l\n\x06\x64\x65lete\x12%.spaceone.api.identity.v2.UserRequest\x1a\x16.google.protobuf.Empty\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/identity/v2/user/delete:\x01*\x12r\n\x03get\x12%.spaceone.api.identity.v2.UserRequest\x1a\".spaceone.api.identity.v2.UserInfo\" \x82\xd3\xe4\x93\x02\x1a\"\x15/identity/v2/user/get:\x01*\x12y\n\x04list\x12).spaceone.api.identity.v2.UserSearchQuery\x1a#.spaceone.api.identity.v2.UsersInfo\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v2/user/list:\x01*\x12k\n\x04stat\x12\'.spaceone.api.identity.v2.UserStatQuery\x1a\x17.google.protobuf.Struct\"!\x82\xd3\xe4\x93\x02\x1b\"\x16/identity/v1/user/stat:\x01*B?Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.identity.v2.user_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z=github.com/cloudforet-io/api/dist/go/spaceone/api/identity/v2'
  _globals['_USER'].methods_by_name['create']._options = None
  _globals['_USER'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/user/create:\001*'
  _globals['_USER'].methods_by_name['update']._options = None
  _globals['_USER'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/user/update:\001*'
  _globals['_USER'].methods_by_name['verify_email']._options = None
  _globals['_USER'].methods_by_name['verify_email']._serialized_options = b'\202\323\344\223\002#\"\036/identity/v2/user/verify-email:\001*'
  _globals['_USER'].methods_by_name['disable_mfa']._options = None
  _globals['_USER'].methods_by_name['disable_mfa']._serialized_options = b'\202\323\344\223\002\"\"\035/identity/v2/user/disable-mfa:\001*'
  _globals['_USER'].methods_by_name['set_required_actions']._options = None
  _globals['_USER'].methods_by_name['set_required_actions']._serialized_options = b'\202\323\344\223\002+\"&/identity/v2/user/set-required-actions:\001*'
  _globals['_USER'].methods_by_name['enable']._options = None
  _globals['_USER'].methods_by_name['enable']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/user/enable:\001*'
  _globals['_USER'].methods_by_name['disable']._options = None
  _globals['_USER'].methods_by_name['disable']._serialized_options = b'\202\323\344\223\002\036\"\031/identity/v2/user/disable:\001*'
  _globals['_USER'].methods_by_name['delete']._options = None
  _globals['_USER'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\035\"\030/identity/v2/user/delete:\001*'
  _globals['_USER'].methods_by_name['get']._options = None
  _globals['_USER'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\032\"\025/identity/v2/user/get:\001*'
  _globals['_USER'].methods_by_name['list']._options = None
  _globals['_USER'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v2/user/list:\001*'
  _globals['_USER'].methods_by_name['stat']._options = None
  _globals['_USER'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002\033\"\026/identity/v1/user/stat:\001*'
  _globals['_AUTHTYPE']._serialized_start=2174
  _globals['_AUTHTYPE']._serialized_end=2227
  _globals['_USERREQUIREDACTION']._serialized_start=2229
  _globals['_USERREQUIREDACTION']._serialized_end=2270
  _globals['_MFA']._serialized_start=231
  _globals['_MFA']._serialized_end=394
  _globals['_MFA_STATE']._serialized_start=350
  _globals['_MFA_STATE']._serialized_end=394
  _globals['_CREATEUSERREQUEST']._serialized_start=397
  _globals['_CREATEUSERREQUEST']._serialized_end=634
  _globals['_UPDATEUSERREQUEST']._serialized_start=637
  _globals['_UPDATEUSERREQUEST']._serialized_end=819
  _globals['_VERIFYUSEREMAILREQUEST']._serialized_start=821
  _globals['_VERIFYUSEREMAILREQUEST']._serialized_end=877
  _globals['_DISABLEUSERMFAREQUEST']._serialized_start=879
  _globals['_DISABLEUSERMFAREQUEST']._serialized_end=919
  _globals['_SETREQUIREDACTIONSREQUEST']._serialized_start=921
  _globals['_SETREQUIREDACTIONSREQUEST']._serialized_end=1037
  _globals['_USERREQUEST']._serialized_start=1039
  _globals['_USERREQUEST']._serialized_end=1069
  _globals['_USERSEARCHQUERY']._serialized_start=1072
  _globals['_USERSEARCHQUERY']._serialized_end=1357
  _globals['_USERSEARCHQUERY_STATE']._serialized_start=1300
  _globals['_USERSEARCHQUERY_STATE']._serialized_end=1357
  _globals['_USERINFO']._serialized_start=1360
  _globals['_USERINFO']._serialized_end=2014
  _globals['_USERINFO_STATE']._serialized_start=1891
  _globals['_USERINFO_STATE']._serialized_end=1954
  _globals['_USERINFO_ROLETYPE']._serialized_start=1956
  _globals['_USERINFO_ROLETYPE']._serialized_end=2014
  _globals['_USERSINFO']._serialized_start=2016
  _globals['_USERSINFO']._serialized_end=2101
  _globals['_USERSTATQUERY']._serialized_start=2103
  _globals['_USERSTATQUERY']._serialized_end=2172
  _globals['_USER']._serialized_start=2273
  _globals['_USER']._serialized_end=3681
# @@protoc_insertion_point(module_scope)
