from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Issue(_message.Message):
    __slots__ = ["id", "solution_attempts", "open", "year", "day", "part", "last_error_code"]
    ID_FIELD_NUMBER: _ClassVar[int]
    SOLUTION_ATTEMPTS_FIELD_NUMBER: _ClassVar[int]
    OPEN_FIELD_NUMBER: _ClassVar[int]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    LAST_ERROR_CODE_FIELD_NUMBER: _ClassVar[int]
    id: int
    solution_attempts: _containers.RepeatedCompositeFieldContainer[Solution]
    open: bool
    year: int
    day: int
    part: int
    last_error_code: str
    def __init__(self, id: _Optional[int] = ..., solution_attempts: _Optional[_Iterable[_Union[Solution, _Mapping]]] = ..., open: bool = ..., year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ..., last_error_code: _Optional[str] = ...) -> None: ...

class Solutions(_message.Message):
    __slots__ = ["solutions"]
    SOLUTIONS_FIELD_NUMBER: _ClassVar[int]
    solutions: _containers.RepeatedCompositeFieldContainer[Solution]
    def __init__(self, solutions: _Optional[_Iterable[_Union[Solution, _Mapping]]] = ...) -> None: ...

class Solution(_message.Message):
    __slots__ = ["year", "day", "part", "big_answer", "string_answer", "answer", "solution_made"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    BIG_ANSWER_FIELD_NUMBER: _ClassVar[int]
    STRING_ANSWER_FIELD_NUMBER: _ClassVar[int]
    ANSWER_FIELD_NUMBER: _ClassVar[int]
    SOLUTION_MADE_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    part: int
    big_answer: int
    string_answer: str
    answer: int
    solution_made: int
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ..., big_answer: _Optional[int] = ..., string_answer: _Optional[str] = ..., answer: _Optional[int] = ..., solution_made: _Optional[int] = ...) -> None: ...

class SolveRequest(_message.Message):
    __slots__ = ["year", "day", "part", "data"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    DATA_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    part: int
    data: str
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ..., data: _Optional[str] = ...) -> None: ...

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
    __slots__ = ["callback", "year"]
    CALLBACK_FIELD_NUMBER: _ClassVar[int]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    callback: str
    year: int
    def __init__(self, callback: _Optional[str] = ..., year: _Optional[int] = ...) -> None: ...

class RegisterResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class AddSolutionRequest(_message.Message):
    __slots__ = ["solution"]
    SOLUTION_FIELD_NUMBER: _ClassVar[int]
    solution: Solution
    def __init__(self, solution: _Optional[_Union[Solution, _Mapping]] = ...) -> None: ...

class AddSolutionResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...

class GetSolutionRequest(_message.Message):
    __slots__ = ["year", "day", "part"]
    YEAR_FIELD_NUMBER: _ClassVar[int]
    DAY_FIELD_NUMBER: _ClassVar[int]
    PART_FIELD_NUMBER: _ClassVar[int]
    year: int
    day: int
    part: int
    def __init__(self, year: _Optional[int] = ..., day: _Optional[int] = ..., part: _Optional[int] = ...) -> None: ...

class GetSolutionResponse(_message.Message):
    __slots__ = ["solution"]
    SOLUTION_FIELD_NUMBER: _ClassVar[int]
    solution: Solution
    def __init__(self, solution: _Optional[_Union[Solution, _Mapping]] = ...) -> None: ...

class SetCookieRequest(_message.Message):
    __slots__ = ["cookie"]
    COOKIE_FIELD_NUMBER: _ClassVar[int]
    cookie: str
    def __init__(self, cookie: _Optional[str] = ...) -> None: ...

class SetCookieResponse(_message.Message):
    __slots__ = []
    def __init__(self) -> None: ...
