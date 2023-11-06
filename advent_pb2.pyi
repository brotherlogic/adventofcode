from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class SolveRequest(_message.Message):
    __slots__ = ["year", "day", "part"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    part: int
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ...) -> None: ...

class SolveResponse(_message.Message):
    __slots__ = ["answer", "string_answer", "big_answer"]
    ANSWER_FIELD_NUMBER: _ClassVar[int]
    STRING_ANSWER_FIELD_NUMBER: _ClassVar[int]
    BIG_ANSWER_FIELD_NUMBER: _ClassVar[int]
    answer: int
    string_answer: str
    big_answer: int
    def __init__(self, answer: _Optional[int] = ..., string_answer: _Optional[str] = ..., big_answer: _Optional[int] = ...) -> None: ...

class UploadRequest(_message.Message):
    __slots__ = ["year", "day", "data", "dual_day", "part"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    DUAL_DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    data: str
    dual_day: bool
    part: int
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., data: _Optional[str] = ..., dual_day: bool = ..., part: _Optional[int] = ...) -> None: ...

class UploadResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetDataRequest(_message.Message):
    __slots__ = ["year", "day", "part"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    part: int
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ...) -> None: ...

class GetDataResponse(_message.Message):
    __slots__ = ["data"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    data: str
    def __init__(self, data: _Optional[str] = ...) -> None: ...

class RegisterRequest(_message.Message):
    __slots__ = ["callback"]
    CALLBACK_FIELD_NUMBER: _ClassVar[int]
    callback: str
    def __init__(self, callback: _Optional[str] = ...) -> None: ...

class RegisterResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...
