# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: advent.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0c\x61\x64vent.proto\x12\x0c\x61\x64ventofcode\"7\n\x0cSolveRequest\x12\x0c\n\x04year\x18\x01 \x01(\x05\x12\x0b\n\x03\x64\x61y\x18\x02 \x01(\x05\x12\x0c\n\x04part\x18\x03 \x01(\x05\"J\n\rSolveResponse\x12\x0e\n\x06\x61nswer\x18\x01 \x01(\x05\x12\x15\n\rstring_answer\x18\x02 \x01(\t\x12\x12\n\nbig_answer\x18\x03 \x01(\x03\"X\n\rUploadRequest\x12\x0c\n\x04year\x18\x01 \x01(\x05\x12\x0b\n\x03\x64\x61y\x18\x02 \x01(\x05\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\t\x12\x10\n\x08\x64ual_day\x18\x04 \x01(\x08\x12\x0c\n\x04part\x18\x05 \x01(\x05\"\x10\n\x0eUploadResponse\"9\n\x0eGetDataRequest\x12\x0c\n\x04year\x18\x01 \x01(\x05\x12\x0b\n\x03\x64\x61y\x18\x02 \x01(\x05\x12\x0c\n\x04part\x18\x03 \x01(\x05\"\x1f\n\x0fGetDataResponse\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\t\"1\n\x0fRegisterRequest\x12\x10\n\x08\x63\x61llback\x18\x01 \x01(\t\x12\x0c\n\x04year\x18\x02 \x01(\x05\"\x12\n\x10RegisterResponse2\xed\x01\n\x13\x41\x64ventOfCodeService\x12\x42\n\x05Solve\x12\x1a.adventofcode.SolveRequest\x1a\x1b.adventofcode.SolveResponse\"\x00\x12\x45\n\x06Upload\x12\x1b.adventofcode.UploadRequest\x1a\x1c.adventofcode.UploadResponse\"\x00\x12K\n\x08Register\x12\x1d.adventofcode.RegisterRequest\x1a\x1e.adventofcode.RegisterResponse\"\x00\x32S\n\rSolverService\x12\x42\n\x05Solve\x12\x1a.adventofcode.SolveRequest\x1a\x1b.adventofcode.SolveResponse\"\x00\x42,Z*github.com/brotherlogic/adventofcode/protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'advent_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z*github.com/brotherlogic/adventofcode/proto'
  _globals['_SOLVEREQUEST']._serialized_start=30
  _globals['_SOLVEREQUEST']._serialized_end=85
  _globals['_SOLVERESPONSE']._serialized_start=87
  _globals['_SOLVERESPONSE']._serialized_end=161
  _globals['_UPLOADREQUEST']._serialized_start=163
  _globals['_UPLOADREQUEST']._serialized_end=251
  _globals['_UPLOADRESPONSE']._serialized_start=253
  _globals['_UPLOADRESPONSE']._serialized_end=269
  _globals['_GETDATAREQUEST']._serialized_start=271
  _globals['_GETDATAREQUEST']._serialized_end=328
  _globals['_GETDATARESPONSE']._serialized_start=330
  _globals['_GETDATARESPONSE']._serialized_end=361
  _globals['_REGISTERREQUEST']._serialized_start=363
  _globals['_REGISTERREQUEST']._serialized_end=412
  _globals['_REGISTERRESPONSE']._serialized_start=414
  _globals['_REGISTERRESPONSE']._serialized_end=432
  _globals['_ADVENTOFCODESERVICE']._serialized_start=435
  _globals['_ADVENTOFCODESERVICE']._serialized_end=672
  _globals['_SOLVERSERVICE']._serialized_start=674
  _globals['_SOLVERSERVICE']._serialized_end=757
# @@protoc_insertion_point(module_scope)
