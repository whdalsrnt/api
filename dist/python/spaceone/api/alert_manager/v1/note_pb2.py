# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: spaceone/api/alert_manager/v1/note.proto
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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(spaceone/api/alert_manager/v1/note.proto\x12\x1dspaceone.api.alert_manager.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1cgoogle/api/annotations.proto\x1a spaceone/api/core/v2/query.proto\"3\n\x11NoteCreateRequest\x12\x10\n\x08\x61lert_id\x18\x01 \x01(\t\x12\x0c\n\x04note\x18\x02 \x01(\t\"2\n\x11NoteUpdateRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\t\x12\x0c\n\x04note\x18\x02 \x01(\t\"\x1e\n\x0bNoteRequest\x12\x0f\n\x07note_id\x18\x01 \x01(\t\"\x9e\x01\n\x0fNoteSearchQuery\x12*\n\x05query\x18\x01 \x01(\x0b\x32\x1b.spaceone.api.core.v2.Query\x12\x14\n\x0cworkspace_id\x18\x02 \x01(\t\x12\x12\n\nservice_id\x18\x03 \x01(\t\x12\x10\n\x08\x61lert_id\x18\x04 \x01(\t\x12\x0f\n\x07note_id\x18\x05 \x01(\t\x12\x12\n\ncreated_by\x18\x1f \x01(\t\"E\n\rNoteStatQuery\x12\x34\n\x05query\x18\x01 \x01(\x0b\x32%.spaceone.api.core.v2.StatisticsQuery\"\xb4\x01\n\x08NoteInfo\x12\x0f\n\x07note_id\x18\x01 \x01(\t\x12\x0c\n\x04note\x18\x02 \x01(\t\x12\x11\n\tdomain_id\x18\x15 \x01(\t\x12\x14\n\x0cworkspace_id\x18\x16 \x01(\t\x12\x12\n\nservice_id\x18\x17 \x01(\t\x12\x10\n\x08\x61lert_id\x18\x18 \x01(\t\x12\x12\n\ncreated_at\x18\x1f \x01(\t\x12\x12\n\nupdated_at\x18  \x01(\t\x12\x12\n\ncreated_by\x18! \x01(\t\"Z\n\tNotesInfo\x12\x38\n\x07results\x18\x01 \x03(\x0b\x32\'.spaceone.api.alert_manager.v1.NoteInfo\x12\x13\n\x0btotal_count\x18\x02 \x01(\x05\x32\xa4\x06\n\x04Note\x12\x8d\x01\n\x06\x63reate\x12\x30.spaceone.api.alert_manager.v1.NoteCreateRequest\x1a\'.spaceone.api.alert_manager.v1.NoteInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/alert-manager/v1/note/create:\x01*\x12\x8d\x01\n\x06update\x12\x30.spaceone.api.alert_manager.v1.NoteUpdateRequest\x1a\'.spaceone.api.alert_manager.v1.NoteInfo\"(\x82\xd3\xe4\x93\x02\"\"\x1d/alert-manager/v1/note/update:\x01*\x12v\n\x06\x64\x65lete\x12*.spaceone.api.alert_manager.v1.NoteRequest\x1a\x16.google.protobuf.Empty\"(\x82\xd3\xe4\x93\x02\"\"\x1d/alert-manager/v1/note/delete:\x01*\x12\x81\x01\n\x03get\x12*.spaceone.api.alert_manager.v1.NoteRequest\x1a\'.spaceone.api.alert_manager.v1.NoteInfo\"%\x82\xd3\xe4\x93\x02\x1f\"\x1a/alert-manager/v1/note/get:\x01*\x12\x88\x01\n\x04list\x12..spaceone.api.alert_manager.v1.NoteSearchQuery\x1a(.spaceone.api.alert_manager.v1.NotesInfo\"&\x82\xd3\xe4\x93\x02 \"\x1b/alert-manager/v1/note/list:\x01*\x12u\n\x04stat\x12,.spaceone.api.alert_manager.v1.NoteStatQuery\x1a\x17.google.protobuf.Struct\"&\x82\xd3\xe4\x93\x02 \"\x1b/alert_manager/v1/note/stat:\x01*BDZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'spaceone.api.alert_manager.v1.note_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cloudforet-io/api/dist/go/spaceone/api/alert-manager/v1'
  _globals['_NOTE'].methods_by_name['create']._loaded_options = None
  _globals['_NOTE'].methods_by_name['create']._serialized_options = b'\202\323\344\223\002\"\"\035/alert-manager/v1/note/create:\001*'
  _globals['_NOTE'].methods_by_name['update']._loaded_options = None
  _globals['_NOTE'].methods_by_name['update']._serialized_options = b'\202\323\344\223\002\"\"\035/alert-manager/v1/note/update:\001*'
  _globals['_NOTE'].methods_by_name['delete']._loaded_options = None
  _globals['_NOTE'].methods_by_name['delete']._serialized_options = b'\202\323\344\223\002\"\"\035/alert-manager/v1/note/delete:\001*'
  _globals['_NOTE'].methods_by_name['get']._loaded_options = None
  _globals['_NOTE'].methods_by_name['get']._serialized_options = b'\202\323\344\223\002\037\"\032/alert-manager/v1/note/get:\001*'
  _globals['_NOTE'].methods_by_name['list']._loaded_options = None
  _globals['_NOTE'].methods_by_name['list']._serialized_options = b'\202\323\344\223\002 \"\033/alert-manager/v1/note/list:\001*'
  _globals['_NOTE'].methods_by_name['stat']._loaded_options = None
  _globals['_NOTE'].methods_by_name['stat']._serialized_options = b'\202\323\344\223\002 \"\033/alert_manager/v1/note/stat:\001*'
  _globals['_NOTECREATEREQUEST']._serialized_start=198
  _globals['_NOTECREATEREQUEST']._serialized_end=249
  _globals['_NOTEUPDATEREQUEST']._serialized_start=251
  _globals['_NOTEUPDATEREQUEST']._serialized_end=301
  _globals['_NOTEREQUEST']._serialized_start=303
  _globals['_NOTEREQUEST']._serialized_end=333
  _globals['_NOTESEARCHQUERY']._serialized_start=336
  _globals['_NOTESEARCHQUERY']._serialized_end=494
  _globals['_NOTESTATQUERY']._serialized_start=496
  _globals['_NOTESTATQUERY']._serialized_end=565
  _globals['_NOTEINFO']._serialized_start=568
  _globals['_NOTEINFO']._serialized_end=748
  _globals['_NOTESINFO']._serialized_start=750
  _globals['_NOTESINFO']._serialized_end=840
  _globals['_NOTE']._serialized_start=843
  _globals['_NOTE']._serialized_end=1647
# @@protoc_insertion_point(module_scope)
