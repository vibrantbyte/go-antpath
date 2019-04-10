/* Created by "go tool cgo" - DO NOT EDIT. */

/* package command-line-arguments */

/* Start of preamble from import "C" comments.  */


#line 3 "/Users/xiaoyueya/projects/gopro/src/github.com/vibrantbyte/go-antpath/antpath.go"
#include <stdlib.h>

#line 1 "cgo-generated-wrapper"


/* End of preamble from import "C" comments.  */


/* Start of boilerplate cgo prologue.  */
#line 1 "cgo-gcc-export-header-prolog"

#ifndef GO_CGO_PROLOGUE_H
#define GO_CGO_PROLOGUE_H

typedef signed char GoInt8;
typedef unsigned char GoUint8;
typedef short GoInt16;
typedef unsigned short GoUint16;
typedef int GoInt32;
typedef unsigned int GoUint32;
typedef long long GoInt64;
typedef unsigned long long GoUint64;
typedef GoInt64 GoInt;
typedef GoUint64 GoUint;
typedef __SIZE_TYPE__ GoUintptr;
typedef float GoFloat32;
typedef double GoFloat64;
typedef float _Complex GoComplex64;
typedef double _Complex GoComplex128;

/*
  static assertion to make sure the file is being used on architecture
  at least with matching size of GoInt.
*/
typedef char _check_for_64_bit_pointer_matching_GoInt[sizeof(void*)==64/8 ? 1:-1];

typedef struct { const char *p; GoInt n; } GoString;
typedef void *GoMap;
typedef void *GoChan;
typedef struct { void *t; void *v; } GoInterface;
typedef struct { void *data; GoInt len; GoInt cap; } GoSlice;

#endif

/* End of boilerplate cgo prologue.  */

#ifdef __cplusplus
extern "C" {
#endif


extern char* Version();

extern void Increment(GoInt* p0);

extern GoUint8 IsPattern(GoString p0);

extern GoUint8 Match(GoString p0, GoString p1);

extern GoUint8 MatchStart(GoString p0, GoString p1);

extern char* ExtractPathWithinPattern(GoString p0, GoString p1);

extern char* ExtractUriTemplateVariables(GoString p0, GoString p1);

extern char* Combine(GoString p0, GoString p1);

extern void SetPathSeparator(GoString p0);

extern void SetCaseSensitive(GoUint8 p0);

extern void SetTrimTokens(GoUint8 p0);

extern void SetCachePatterns(GoUint8 p0);

#ifdef __cplusplus
}
#endif
