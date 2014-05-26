package gommd

/*
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <libMultiMarkdown.h>
#cgo CFLAGS: -IMultiMarkdown-4/ --std=c99
#cgo LDFLAGS: -LMultiMarkdown-4/ -lmmd
*/
import "C"

import (
	"unsafe"
)

const (
	EXT_COMPATIBILITY       = 1 << 0  /* Markdown compatibility mode */
	EXT_COMPLETE            = 1 << 1  /* Create complete document */
	EXT_SNIPPET             = 1 << 2  /* Create snippet only */
	EXT_HEAD_CLOSED         = 1 << 3  /* for use by parser */
	EXT_SMART               = 1 << 4  /* Enable Smart quotes */
	EXT_NOTES               = 1 << 5  /* Enable Footnotes */
	EXT_NO_LABELS           = 1 << 6  /* Don't add anchors to headers, etc. */
	EXT_FILTER_STYLES       = 1 << 7  /* Filter out style blocks */
	EXT_FILTER_HTML         = 1 << 8  /* Filter out raw HTML */
	EXT_PROCESS_HTML        = 1 << 9  /* Process Markdown inside HTML */
	EXT_NO_METADATA         = 1 << 10 /* Don't parse Metadata */
	EXT_OBFUSCATE           = 1 << 11 /* Mask email addresses */
	EXT_CRITIC              = 1 << 12 /* Critic Markup Support */
	EXT_CRITIC_ACCEPT       = 1 << 13 /* Accept all proposed changes */
	EXT_CRITIC_REJECT       = 1 << 14 /* Reject all proposed changes */
	EXT_RANDOM_FOOT         = 1 << 15 /* Use random numbers for footnote links */
	EXT_HEADINGSECTION      = 1 << 16 /* Group blocks under parent heading */
	EXT_ESCAPED_LINE_BREAKS = 1 << 17 /* Escaped line break */
	EXT_FAKE                = 1 << 31 /* 31 is highest number allowed */
)

const (
	FORMAT_HTML                  = 0
	FORMAT_TEXT                  = 1
	FORMAT_LATEX                 = 2
	FORMAT_MEMOIR                = 3
	FORMAT_BEAMER                = 4
	FORMAT_OPML                  = 5
	FORMAT_ODF                   = 6
	FORMAT_RTF                   = 7
	FORMAT_ORIGINAL              = 8 /* Not currently used */
	FORMAT_CRITIC_ACCEPT         = 9
	FORMAT_CRITIC_REJECT         = 10
	FORMAT_CRITIC_HTML_HIGHLIGHT = 11
	FORMAT_LYX                   = 12
)

func MarkdownToString(source string, exts uint64, format int) string {
	res := C.markdown_to_string(C.CString(source), C.ulong(exts), C.int(format))
	defer C.free(unsafe.Pointer(res))
	return C.GoString(res)
}
