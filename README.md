# GoWinHttp

GoWinHttp provides a Goish interface to
the
[WinHTTP API](https://msdn.microsoft.com/en-us/library/windows/desktop/aa382925(v=vs.85).aspx). This
aims to expose an API that matches or is quite close to the original
API, while cleaning up the interface where possible to help it fit in
with Go. Currently as much of WinHTTP as required for SpiderOak
Semaphor is implemented, but we aim for eventual completeness. Because
Win32 (and thus WinHTTP) has a reasonably consistent pattern of use
that happens to not be what's considered to be idiomatic Go, there
will still be some rough edges. However, this package should enable
someone with a basic understanding of WinHTTP and knowledge of Go to
produce software using winhttp.dll.

It should be obvious but to be clear: this will build and run only on
Windows.

We make very heavy use of
the
[expanded Windows sys package](https://godoc.org/golang.org/x/sys/windows). 

## Installing:

```
go get github.com/SpiderOak/GoWinHttp
```

## Use:

### Basics

The assumption we are making is that you can look up function calls in
MSDN and you have some knowledge of Go. To help mesh understanding, we
use argument and field names from MSDN, stripping off Hungarian
notation when it doesn't reduce clarity. We also try to expose Go
types to the caller and convert to Win32-friendly types
internally. Finally, arguments and struct fields that are noted as
reserved in MSDN we try to hide here- reserved arguments aren't in
functions, and while we need to keep reserved struct fields to retain
memory layouts, we don't export them.

**NOTE:** We currently do not support async WinHTTP operation. This
will be rectified soon.

### Types

As above, we try to hold hard to using unspecial Go types for function
arguments, and these types will either map to their Win32 versions or
we will do a little magic to make it work (such as working with
`[]byte` instead of forcing you to deal with `*byte` for data). There
three exceptions to this:

1. We define an HInternet type so the compiler can help ensure you're
   only passing around internet handles from us.
1. We have a variety of struct types not available from the system.
1. Generic interfaces in WinHTTP requiring use of pointers. See "Pointers", below.

### Pointers

The Win32 API operates in an OO-ish manner via a C API. This means we
will be sometimes passing data via pointer, and sometimes we aren't
going to be able to hide this from
you. The
[unsafe.Pointer documentation](https://golang.org/pkg/unsafe/#Pointer)
notes special use cases for making syscalls. Functions where we cannot
avoid using C-compatible pointers via the
`uintptr(unsafe.Pointer(&x))` pattern are specifically marked with
compiler pragmas to operate safely with that pattern. Write to your
local congressperson or MP about generics in Go if this is upsetting.

### Example:

Get a session handle:

```go
hsession, err := winhttp.Open(
    "winhttp",
    winhttp.WINHTTP_ACCESS_TYPE_DEFAULT_PROXY,
    "",
    "",
    0,
)
if err != nil {
    return "", err
}
```

Set options on a request (also an example of passing in pointers):

```go
dwFlags := uint32(winhttp.SECURITY_FLAG_IGNORE_UNKNOWN_CA)

err = winhttp.SetOption(
    request,
    winhttp.WINHTTP_OPTION_SECURITY_FLAGS,
    uintptr(unsafe.Pointer(&dwFlags)),
    unsafe.Sizeof(dwFlags),
)
if err != nil {
    fmt.Printf("error in WinHttpSetOption: %s", err)
}
```

## Contact

Feel free to reachout to matt at spideroak dot com for questions or
concerns about this library!

## Legalese

This work is covered by the Mozilla Public License, v2. See LICENSE.md
for more information.

&copy; 2017 SpiderOak, Inc.
