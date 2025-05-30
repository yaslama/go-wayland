// Generated by go-wayland-scanner
// https://github.com/yaslama/go-wayland/cmd/go-wayland-scanner
// XML file : https://gitlab.freedesktop.org/wayland/wayland-protocols/-/raw/1.41/unstable/text-input/text-input-unstable-v3.xml?ref_type=tags
//
// text_input_unstable_v3 Protocol Copyright:
//
// Copyright © 2012, 2013 Intel Corporation
// Copyright © 2015, 2016 Jan Arne Petersen
// Copyright © 2017, 2018 Red Hat, Inc.
// Copyright © 2018       Purism SPC
//
// Permission to use, copy, modify, distribute, and sell this
// software and its documentation for any purpose is hereby granted
// without fee, provided that the above copyright notice appear in
// all copies and that both that copyright notice and this permission
// notice appear in supporting documentation, and that the name of
// the copyright holders not be used in advertising or publicity
// pertaining to distribution of the software without specific,
// written prior permission.  The copyright holders make no
// representations about the suitability of this software for any
// purpose.  It is provided "as is" without express or implied
// warranty.
//
// THE COPYRIGHT HOLDERS DISCLAIM ALL WARRANTIES WITH REGARD TO THIS
// SOFTWARE, INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS, IN NO EVENT SHALL THE COPYRIGHT HOLDERS BE LIABLE FOR ANY
// SPECIAL, INDIRECT OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN
// AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
// ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.

package text_input

import "github.com/yaslama/go-wayland/wayland/client"

// TextInputInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const TextInputInterfaceName = "zwp_text_input_v3"

// TextInput : text input
//
// The zwp_text_input_v3 interface represents text input and input methods
// associated with a seat. It provides enter/leave events to follow the
// text input focus for a seat.
//
// Requests are used to enable/disable the text-input object and set
// state information like surrounding and selected text or the content type.
// The information about the entered text is sent to the text-input object
// via the preedit_string and commit_string events.
//
// Text is valid UTF-8 encoded, indices and lengths are in bytes. Indices
// must not point to middle bytes inside a code point: they must either
// point to the first byte of a code point or to the end of the buffer.
// Lengths must be measured between two valid indices.
//
// Focus moving throughout surfaces will result in the emission of
// zwp_text_input_v3.enter and zwp_text_input_v3.leave events. The focused
// surface must commit zwp_text_input_v3.enable and
// zwp_text_input_v3.disable requests as the keyboard focus moves across
// editable and non-editable elements of the UI. Those two requests are not
// expected to be paired with each other, the compositor must be able to
// handle consecutive series of the same request.
//
// State is sent by the state requests (set_surrounding_text,
// set_content_type and set_cursor_rectangle) and a commit request. After an
// enter event or disable request all state information is invalidated and
// needs to be resent by the client.
type TextInput struct {
	client.BaseProxy
	enterHandler                 TextInputEnterHandlerFunc
	leaveHandler                 TextInputLeaveHandlerFunc
	preeditStringHandler         TextInputPreeditStringHandlerFunc
	commitStringHandler          TextInputCommitStringHandlerFunc
	deleteSurroundingTextHandler TextInputDeleteSurroundingTextHandlerFunc
	doneHandler                  TextInputDoneHandlerFunc
}

// NewTextInput : text input
//
// The zwp_text_input_v3 interface represents text input and input methods
// associated with a seat. It provides enter/leave events to follow the
// text input focus for a seat.
//
// Requests are used to enable/disable the text-input object and set
// state information like surrounding and selected text or the content type.
// The information about the entered text is sent to the text-input object
// via the preedit_string and commit_string events.
//
// Text is valid UTF-8 encoded, indices and lengths are in bytes. Indices
// must not point to middle bytes inside a code point: they must either
// point to the first byte of a code point or to the end of the buffer.
// Lengths must be measured between two valid indices.
//
// Focus moving throughout surfaces will result in the emission of
// zwp_text_input_v3.enter and zwp_text_input_v3.leave events. The focused
// surface must commit zwp_text_input_v3.enable and
// zwp_text_input_v3.disable requests as the keyboard focus moves across
// editable and non-editable elements of the UI. Those two requests are not
// expected to be paired with each other, the compositor must be able to
// handle consecutive series of the same request.
//
// State is sent by the state requests (set_surrounding_text,
// set_content_type and set_cursor_rectangle) and a commit request. After an
// enter event or disable request all state information is invalidated and
// needs to be resent by the client.
func NewTextInput(ctx *client.Context) *TextInput {
	zwpTextInputV3 := &TextInput{}
	ctx.Register(zwpTextInputV3)
	return zwpTextInputV3
}

// Destroy : Destroy the wp_text_input
//
// Destroy the wp_text_input object. Also disables all surfaces enabled
// through this wp_text_input object.
func (i *TextInput) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Enable : Request text input to be enabled
//
// Requests text input on the surface previously obtained from the enter
// event.
//
// This request must be issued every time the active text input changes
// to a new one, including within the current surface. Use
// zwp_text_input_v3.disable when there is no longer any input focus on
// the current surface.
//
// Clients must not enable more than one text input on the single seat
// and should disable the current text input before enabling the new one.
// At most one instance of text input may be in enabled state per instance,
// Requests to enable the another text input when some text input is active
// must be ignored by compositor.
//
// This request resets all state associated with previous enable, disable,
// set_surrounding_text, set_text_change_cause, set_content_type, and
// set_cursor_rectangle requests, as well as the state associated with
// preedit_string, commit_string, and delete_surrounding_text events.
//
// The set_surrounding_text, set_content_type and set_cursor_rectangle
// requests must follow if the text input supports the necessary
// functionality.
//
// State set with this request is double-buffered. It will get applied on
// the next zwp_text_input_v3.commit request, and stay valid until the
// next committed enable or disable request.
//
// The changes must be applied by the compositor after issuing a
// zwp_text_input_v3.commit request.
func (i *TextInput) Enable() error {
	const opcode = 1
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Disable : Disable text input on a surface
//
// Explicitly disable text input on the current surface (typically when
// there is no focus on any text entry inside the surface).
//
// State set with this request is double-buffered. It will get applied on
// the next zwp_text_input_v3.commit request.
func (i *TextInput) Disable() error {
	const opcode = 2
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetSurroundingText : sets the surrounding text
//
// Sets the surrounding plain text around the input, excluding the preedit
// text.
//
// The client should notify the compositor of any changes in any of the
// values carried with this request, including changes caused by handling
// incoming text-input events as well as changes caused by other
// mechanisms like keyboard typing.
//
// If the client is unaware of the text around the cursor, it should not
// issue this request, to signify lack of support to the compositor.
//
// Text is UTF-8 encoded, and should include the cursor position, the
// complete selection and additional characters before and after them.
// There is a maximum length of wayland messages, so text can not be
// longer than 4000 bytes.
//
// Cursor is the byte offset of the cursor within text buffer.
//
// Anchor is the byte offset of the selection anchor within text buffer.
// If there is no selected text, anchor is the same as cursor.
//
// If any preedit text is present, it is replaced with a cursor for the
// purpose of this event.
//
// Values set with this request are double-buffered. They will get applied
// on the next zwp_text_input_v3.commit request, and stay valid until the
// next committed enable or disable request.
//
// The initial state for affected fields is empty, meaning that the text
// input does not support sending surrounding text. If the empty values
// get applied, subsequent attempts to change them may have no effect.
func (i *TextInput) SetSurroundingText(text string, cursor, anchor int32) error {
	const opcode = 3
	textLen := client.PaddedLen(len(text) + 1)
	_reqBufLen := 8 + (4 + textLen) + 4 + 4
	_reqBuf := make([]byte, _reqBufLen)
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutString(_reqBuf[l:l+(4+textLen)], text, textLen)
	l += (4 + textLen)
	client.PutUint32(_reqBuf[l:l+4], uint32(cursor))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(anchor))
	l += 4
	err := i.Context().WriteMsg(_reqBuf, nil)
	return err
}

// SetTextChangeCause : indicates the cause of surrounding text change
//
// Tells the compositor why the text surrounding the cursor changed.
//
// Whenever the client detects an external change in text, cursor, or
// anchor posision, it must issue this request to the compositor. This
// request is intended to give the input method a chance to update the
// preedit text in an appropriate way, e.g. by removing it when the user
// starts typing with a keyboard.
//
// cause describes the source of the change.
//
// The value set with this request is double-buffered. It must be applied
// and reset to initial at the next zwp_text_input_v3.commit request.
//
// The initial value of cause is input_method.
func (i *TextInput) SetTextChangeCause(cause uint32) error {
	const opcode = 4
	const _reqBufLen = 8 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(cause))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetContentType : set content purpose and hint
//
// Sets the content purpose and content hint. While the purpose is the
// basic purpose of an input field, the hint flags allow to modify some of
// the behavior.
//
// Values set with this request are double-buffered. They will get applied
// on the next zwp_text_input_v3.commit request.
// Subsequent attempts to update them may have no effect. The values
// remain valid until the next committed enable or disable request.
//
// The initial value for hint is none, and the initial value for purpose
// is normal.
func (i *TextInput) SetContentType(hint, purpose uint32) error {
	const opcode = 5
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(hint))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(purpose))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// SetCursorRectangle : set cursor position
//
// Marks an area around the cursor as a x, y, width, height rectangle in
// surface local coordinates.
//
// Allows the compositor to put a window with word suggestions near the
// cursor, without obstructing the text being input.
//
// If the client is unaware of the position of edited text, it should not
// issue this request, to signify lack of support to the compositor.
//
// Values set with this request are double-buffered. They will get applied
// on the next zwp_text_input_v3.commit request, and stay valid until the
// next committed enable or disable request.
//
// The initial values describing a cursor rectangle are empty. That means
// the text input does not support describing the cursor area. If the
// empty values get applied, subsequent attempts to change them may have
// no effect.
func (i *TextInput) SetCursorRectangle(x, y, width, height int32) error {
	const opcode = 6
	const _reqBufLen = 8 + 4 + 4 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(x))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(y))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(width))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(height))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// Commit : commit state
//
// Atomically applies state changes recently sent to the compositor.
//
// The commit request establishes and updates the state of the client, and
// must be issued after any changes to apply them.
//
// Text input state (enabled status, content purpose, content hint,
// surrounding text and change cause, cursor rectangle) is conceptually
// double-buffered within the context of a text input, i.e. between a
// committed enable request and the following committed enable or disable
// request.
//
// Protocol requests modify the pending state, as opposed to the current
// state in use by the input method. A commit request atomically applies
// all pending state, replacing the current state. After commit, the new
// pending state is as documented for each related request.
//
// Requests are applied in the order of arrival.
//
// Neither current nor pending state are modified unless noted otherwise.
//
// The compositor must count the number of commit requests coming from
// each zwp_text_input_v3 object and use the count as the serial in done
// events.
func (i *TextInput) Commit() error {
	const opcode = 7
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

type TextInputChangeCause uint32

// TextInputChangeCause : text change reason
//
// Reason for the change of surrounding text or cursor posision.
const (
	// TextInputChangeCauseInputMethod : input method caused the change
	TextInputChangeCauseInputMethod TextInputChangeCause = 0
	// TextInputChangeCauseOther : something else than the input method caused the change
	TextInputChangeCauseOther TextInputChangeCause = 1
)

func (e TextInputChangeCause) Name() string {
	switch e {
	case TextInputChangeCauseInputMethod:
		return "input_method"
	case TextInputChangeCauseOther:
		return "other"
	default:
		return ""
	}
}

func (e TextInputChangeCause) Value() string {
	switch e {
	case TextInputChangeCauseInputMethod:
		return "0"
	case TextInputChangeCauseOther:
		return "1"
	default:
		return ""
	}
}

func (e TextInputChangeCause) String() string {
	return e.Name() + "=" + e.Value()
}

type TextInputContentHint uint32

// TextInputContentHint : content hint
//
// Content hint is a bitmask to allow to modify the behavior of the text
// input.
const (
	// TextInputContentHintNone : no special behavior
	TextInputContentHintNone TextInputContentHint = 0x0
	// TextInputContentHintCompletion : suggest word completions
	TextInputContentHintCompletion TextInputContentHint = 0x1
	// TextInputContentHintSpellcheck : suggest word corrections
	TextInputContentHintSpellcheck TextInputContentHint = 0x2
	// TextInputContentHintAutoCapitalization : switch to uppercase letters at the start of a sentence
	TextInputContentHintAutoCapitalization TextInputContentHint = 0x4
	// TextInputContentHintLowercase : prefer lowercase letters
	TextInputContentHintLowercase TextInputContentHint = 0x8
	// TextInputContentHintUppercase : prefer uppercase letters
	TextInputContentHintUppercase TextInputContentHint = 0x10
	// TextInputContentHintTitlecase : prefer casing for titles and headings (can be language dependent)
	TextInputContentHintTitlecase TextInputContentHint = 0x20
	// TextInputContentHintHiddenText : characters should be hidden
	TextInputContentHintHiddenText TextInputContentHint = 0x40
	// TextInputContentHintSensitiveData : typed text should not be stored
	TextInputContentHintSensitiveData TextInputContentHint = 0x80
	// TextInputContentHintLatin : just Latin characters should be entered
	TextInputContentHintLatin TextInputContentHint = 0x100
	// TextInputContentHintMultiline : the text input is multiline
	TextInputContentHintMultiline TextInputContentHint = 0x200
)

func (e TextInputContentHint) Name() string {
	switch e {
	case TextInputContentHintNone:
		return "none"
	case TextInputContentHintCompletion:
		return "completion"
	case TextInputContentHintSpellcheck:
		return "spellcheck"
	case TextInputContentHintAutoCapitalization:
		return "auto_capitalization"
	case TextInputContentHintLowercase:
		return "lowercase"
	case TextInputContentHintUppercase:
		return "uppercase"
	case TextInputContentHintTitlecase:
		return "titlecase"
	case TextInputContentHintHiddenText:
		return "hidden_text"
	case TextInputContentHintSensitiveData:
		return "sensitive_data"
	case TextInputContentHintLatin:
		return "latin"
	case TextInputContentHintMultiline:
		return "multiline"
	default:
		return ""
	}
}

func (e TextInputContentHint) Value() string {
	switch e {
	case TextInputContentHintNone:
		return "0x0"
	case TextInputContentHintCompletion:
		return "0x1"
	case TextInputContentHintSpellcheck:
		return "0x2"
	case TextInputContentHintAutoCapitalization:
		return "0x4"
	case TextInputContentHintLowercase:
		return "0x8"
	case TextInputContentHintUppercase:
		return "0x10"
	case TextInputContentHintTitlecase:
		return "0x20"
	case TextInputContentHintHiddenText:
		return "0x40"
	case TextInputContentHintSensitiveData:
		return "0x80"
	case TextInputContentHintLatin:
		return "0x100"
	case TextInputContentHintMultiline:
		return "0x200"
	default:
		return ""
	}
}

func (e TextInputContentHint) String() string {
	return e.Name() + "=" + e.Value()
}

type TextInputContentPurpose uint32

// TextInputContentPurpose : content purpose
//
// The content purpose allows to specify the primary purpose of a text
// input.
//
// This allows an input method to show special purpose input panels with
// extra characters or to disallow some characters.
const (
	// TextInputContentPurposeNormal : default input, allowing all characters
	TextInputContentPurposeNormal TextInputContentPurpose = 0
	// TextInputContentPurposeAlpha : allow only alphabetic characters
	TextInputContentPurposeAlpha TextInputContentPurpose = 1
	// TextInputContentPurposeDigits : allow only digits
	TextInputContentPurposeDigits TextInputContentPurpose = 2
	// TextInputContentPurposeNumber : input a number (including decimal separator and sign)
	TextInputContentPurposeNumber TextInputContentPurpose = 3
	// TextInputContentPurposePhone : input a phone number
	TextInputContentPurposePhone TextInputContentPurpose = 4
	// TextInputContentPurposeUrl : input an URL
	TextInputContentPurposeUrl TextInputContentPurpose = 5
	// TextInputContentPurposeEmail : input an email address
	TextInputContentPurposeEmail TextInputContentPurpose = 6
	// TextInputContentPurposeName : input a name of a person
	TextInputContentPurposeName TextInputContentPurpose = 7
	// TextInputContentPurposePassword : input a password (combine with sensitive_data hint)
	TextInputContentPurposePassword TextInputContentPurpose = 8
	// TextInputContentPurposePin : input is a numeric password (combine with sensitive_data hint)
	TextInputContentPurposePin TextInputContentPurpose = 9
	// TextInputContentPurposeDate : input a date
	TextInputContentPurposeDate TextInputContentPurpose = 10
	// TextInputContentPurposeTime : input a time
	TextInputContentPurposeTime TextInputContentPurpose = 11
	// TextInputContentPurposeDatetime : input a date and time
	TextInputContentPurposeDatetime TextInputContentPurpose = 12
	// TextInputContentPurposeTerminal : input for a terminal
	TextInputContentPurposeTerminal TextInputContentPurpose = 13
)

func (e TextInputContentPurpose) Name() string {
	switch e {
	case TextInputContentPurposeNormal:
		return "normal"
	case TextInputContentPurposeAlpha:
		return "alpha"
	case TextInputContentPurposeDigits:
		return "digits"
	case TextInputContentPurposeNumber:
		return "number"
	case TextInputContentPurposePhone:
		return "phone"
	case TextInputContentPurposeUrl:
		return "url"
	case TextInputContentPurposeEmail:
		return "email"
	case TextInputContentPurposeName:
		return "name"
	case TextInputContentPurposePassword:
		return "password"
	case TextInputContentPurposePin:
		return "pin"
	case TextInputContentPurposeDate:
		return "date"
	case TextInputContentPurposeTime:
		return "time"
	case TextInputContentPurposeDatetime:
		return "datetime"
	case TextInputContentPurposeTerminal:
		return "terminal"
	default:
		return ""
	}
}

func (e TextInputContentPurpose) Value() string {
	switch e {
	case TextInputContentPurposeNormal:
		return "0"
	case TextInputContentPurposeAlpha:
		return "1"
	case TextInputContentPurposeDigits:
		return "2"
	case TextInputContentPurposeNumber:
		return "3"
	case TextInputContentPurposePhone:
		return "4"
	case TextInputContentPurposeUrl:
		return "5"
	case TextInputContentPurposeEmail:
		return "6"
	case TextInputContentPurposeName:
		return "7"
	case TextInputContentPurposePassword:
		return "8"
	case TextInputContentPurposePin:
		return "9"
	case TextInputContentPurposeDate:
		return "10"
	case TextInputContentPurposeTime:
		return "11"
	case TextInputContentPurposeDatetime:
		return "12"
	case TextInputContentPurposeTerminal:
		return "13"
	default:
		return ""
	}
}

func (e TextInputContentPurpose) String() string {
	return e.Name() + "=" + e.Value()
}

// TextInputEnterEvent : enter event
//
// Notification that this seat's text-input focus is on a certain surface.
//
// If client has created multiple text input objects, compositor must send
// this event to all of them.
//
// When the seat has the keyboard capability the text-input focus follows
// the keyboard focus. This event sets the current surface for the
// text-input object.
type TextInputEnterEvent struct {
	Surface *client.Surface
}
type TextInputEnterHandlerFunc func(TextInputEnterEvent)

// SetEnterHandler : sets handler for TextInputEnterEvent
func (i *TextInput) SetEnterHandler(f TextInputEnterHandlerFunc) {
	i.enterHandler = f
}

// TextInputLeaveEvent : leave event
//
// Notification that this seat's text-input focus is no longer on a
// certain surface. The client should reset any preedit string previously
// set.
//
// The leave notification clears the current surface. It is sent before
// the enter notification for the new focus. After leave event, compositor
// must ignore requests from any text input instances until next enter
// event.
//
// When the seat has the keyboard capability the text-input focus follows
// the keyboard focus.
type TextInputLeaveEvent struct {
	Surface *client.Surface
}
type TextInputLeaveHandlerFunc func(TextInputLeaveEvent)

// SetLeaveHandler : sets handler for TextInputLeaveEvent
func (i *TextInput) SetLeaveHandler(f TextInputLeaveHandlerFunc) {
	i.leaveHandler = f
}

// TextInputPreeditStringEvent : pre-edit
//
// Notify when a new composing text (pre-edit) should be set at the
// current cursor position. Any previously set composing text must be
// removed. Any previously existing selected text must be removed.
//
// The argument text contains the pre-edit string buffer.
//
// The parameters cursor_begin and cursor_end are counted in bytes
// relative to the beginning of the submitted text buffer. Cursor should
// be hidden when both are equal to -1.
//
// They could be represented by the client as a line if both values are
// the same, or as a text highlight otherwise.
//
// Values set with this event are double-buffered. They must be applied
// and reset to initial on the next zwp_text_input_v3.done event.
//
// The initial value of text is an empty string, and cursor_begin,
// cursor_end and cursor_hidden are all 0.
type TextInputPreeditStringEvent struct {
	Text        string
	CursorBegin int32
	CursorEnd   int32
}
type TextInputPreeditStringHandlerFunc func(TextInputPreeditStringEvent)

// SetPreeditStringHandler : sets handler for TextInputPreeditStringEvent
func (i *TextInput) SetPreeditStringHandler(f TextInputPreeditStringHandlerFunc) {
	i.preeditStringHandler = f
}

// TextInputCommitStringEvent : text commit
//
// Notify when text should be inserted into the editor widget. The text to
// commit could be either just a single character after a key press or the
// result of some composing (pre-edit).
//
// Values set with this event are double-buffered. They must be applied
// and reset to initial on the next zwp_text_input_v3.done event.
//
// The initial value of text is an empty string.
type TextInputCommitStringEvent struct {
	Text string
}
type TextInputCommitStringHandlerFunc func(TextInputCommitStringEvent)

// SetCommitStringHandler : sets handler for TextInputCommitStringEvent
func (i *TextInput) SetCommitStringHandler(f TextInputCommitStringHandlerFunc) {
	i.commitStringHandler = f
}

// TextInputDeleteSurroundingTextEvent : delete surrounding text
//
// Notify when the text around the current cursor position should be
// deleted.
//
// Before_length and after_length are the number of bytes before and after
// the current cursor index (excluding the selection) to delete.
//
// If a preedit text is present, in effect before_length is counted from
// the beginning of it, and after_length from its end (see done event
// sequence).
//
// Values set with this event are double-buffered. They must be applied
// and reset to initial on the next zwp_text_input_v3.done event.
//
// The initial values of both before_length and after_length are 0.
type TextInputDeleteSurroundingTextEvent struct {
	BeforeLength uint32
	AfterLength  uint32
}
type TextInputDeleteSurroundingTextHandlerFunc func(TextInputDeleteSurroundingTextEvent)

// SetDeleteSurroundingTextHandler : sets handler for TextInputDeleteSurroundingTextEvent
func (i *TextInput) SetDeleteSurroundingTextHandler(f TextInputDeleteSurroundingTextHandlerFunc) {
	i.deleteSurroundingTextHandler = f
}

// TextInputDoneEvent : apply changes
//
// Instruct the application to apply changes to state requested by the
// preedit_string, commit_string and delete_surrounding_text events. The
// state relating to these events is double-buffered, and each one
// modifies the pending state. This event replaces the current state with
// the pending state.
//
// The application must proceed by evaluating the changes in the following
// order:
//
// 1. Replace existing preedit string with the cursor.
// 2. Delete requested surrounding text.
// 3. Insert commit string with the cursor at its end.
// 4. Calculate surrounding text to send.
// 5. Insert new preedit text in cursor position.
// 6. Place cursor inside preedit text.
//
// The serial number reflects the last state of the zwp_text_input_v3
// object known to the compositor. The value of the serial argument must
// be equal to the number of commit requests already issued on that object.
//
// When the client receives a done event with a serial different than the
// number of past commit requests, it must proceed with evaluating and
// applying the changes as normal, except it should not change the current
// state of the zwp_text_input_v3 object. All pending state requests
// (set_surrounding_text, set_content_type and set_cursor_rectangle) on
// the zwp_text_input_v3 object should be sent and committed after
// receiving a zwp_text_input_v3.done event with a matching serial.
type TextInputDoneEvent struct {
	Serial uint32
}
type TextInputDoneHandlerFunc func(TextInputDoneEvent)

// SetDoneHandler : sets handler for TextInputDoneEvent
func (i *TextInput) SetDoneHandler(f TextInputDoneHandlerFunc) {
	i.doneHandler = f
}

func (i *TextInput) Dispatch(opcode uint32, fd int, data []byte) {
	switch opcode {
	case 0:
		if i.enterHandler == nil {
			return
		}
		var e TextInputEnterEvent
		l := 0
		e.Surface = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*client.Surface)
		l += 4

		i.enterHandler(e)
	case 1:
		if i.leaveHandler == nil {
			return
		}
		var e TextInputLeaveEvent
		l := 0
		e.Surface = i.Context().GetProxy(client.Uint32(data[l : l+4])).(*client.Surface)
		l += 4

		i.leaveHandler(e)
	case 2:
		if i.preeditStringHandler == nil {
			return
		}
		var e TextInputPreeditStringEvent
		l := 0
		textLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Text = client.String(data[l : l+textLen])
		l += textLen
		e.CursorBegin = int32(client.Uint32(data[l : l+4]))
		l += 4
		e.CursorEnd = int32(client.Uint32(data[l : l+4]))
		l += 4

		i.preeditStringHandler(e)
	case 3:
		if i.commitStringHandler == nil {
			return
		}
		var e TextInputCommitStringEvent
		l := 0
		textLen := client.PaddedLen(int(client.Uint32(data[l : l+4])))
		l += 4
		e.Text = client.String(data[l : l+textLen])
		l += textLen

		i.commitStringHandler(e)
	case 4:
		if i.deleteSurroundingTextHandler == nil {
			return
		}
		var e TextInputDeleteSurroundingTextEvent
		l := 0
		e.BeforeLength = client.Uint32(data[l : l+4])
		l += 4
		e.AfterLength = client.Uint32(data[l : l+4])
		l += 4

		i.deleteSurroundingTextHandler(e)
	case 5:
		if i.doneHandler == nil {
			return
		}
		var e TextInputDoneEvent
		l := 0
		e.Serial = client.Uint32(data[l : l+4])
		l += 4

		i.doneHandler(e)
	}
}

// TextInputManagerInterfaceName is the name of the interface as it appears in the [client.Registry].
// It can be used to match the [client.RegistryGlobalEvent.Interface] in the
// [Registry.SetGlobalHandler] and can be used in [Registry.Bind] if this applies.
const TextInputManagerInterfaceName = "zwp_text_input_manager_v3"

// TextInputManager : text input manager
//
// A factory for text-input objects. This object is a global singleton.
type TextInputManager struct {
	client.BaseProxy
}

// NewTextInputManager : text input manager
//
// A factory for text-input objects. This object is a global singleton.
func NewTextInputManager(ctx *client.Context) *TextInputManager {
	zwpTextInputManagerV3 := &TextInputManager{}
	ctx.Register(zwpTextInputManagerV3)
	return zwpTextInputManagerV3
}

// Destroy : Destroy the wp_text_input_manager
//
// Destroy the wp_text_input_manager object.
func (i *TextInputManager) Destroy() error {
	defer i.Context().Unregister(i)
	const opcode = 0
	const _reqBufLen = 8
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return err
}

// GetTextInput : create a new text input object
//
// Creates a new text-input object for a given seat.
func (i *TextInputManager) GetTextInput(seat *client.Seat) (*TextInput, error) {
	id := NewTextInput(i.Context())
	const opcode = 1
	const _reqBufLen = 8 + 4 + 4
	var _reqBuf [_reqBufLen]byte
	l := 0
	client.PutUint32(_reqBuf[l:4], i.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], uint32(_reqBufLen<<16|opcode&0x0000ffff))
	l += 4
	client.PutUint32(_reqBuf[l:l+4], id.ID())
	l += 4
	client.PutUint32(_reqBuf[l:l+4], seat.ID())
	l += 4
	err := i.Context().WriteMsg(_reqBuf[:], nil)
	return id, err
}
