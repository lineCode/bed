package event

import "github.com/itchyny/bed/mode"

// Event represents the event emitted by UI.
type Event struct {
	Type    Type
	Range   *Range
	Count   int64
	Rune    rune
	CmdName string
	Arg     string
	Error   error
	Mode    mode.Mode
}

// Type ...
type Type int

// Event types
const (
	Nop Type = iota
	Redraw

	CursorUp
	CursorDown
	CursorLeft
	CursorRight
	CursorPrev
	CursorNext
	CursorHead
	CursorEnd
	CursorGoto
	ScrollUp
	ScrollDown
	PageUp
	PageDown
	PageUpHalf
	PageDownHalf
	PageTop
	PageEnd
	JumpTo
	JumpBack

	DeleteByte
	DeletePrevByte
	Increment
	Decrement
	SwitchFocus

	StartInsert
	StartInsertHead
	StartAppend
	StartAppendEnd
	StartReplaceByte
	StartReplace

	ExitInsert
	Backspace
	Delete
	Rune

	Undo
	Redo

	StartVisual
	SwitchVisualEnd
	ExitVisual

	StartCmdlineCommand
	StartCmdlineSearchForward
	StartCmdlineSearchBackward
	BackspaceCmdline
	DeleteCmdline
	DeleteWordCmdline
	ClearToHeadCmdline
	ClearCmdline
	ExitCmdline
	CompleteForwardCmdline
	CompleteBackCmdline
	ExecuteCmdline
	ExecuteSearch
	NextSearch
	PreviousSearch

	Edit
	New
	Vnew
	Wincmd
	FocusWindowUp
	FocusWindowDown
	FocusWindowLeft
	FocusWindowRight
	FocusWindowTopLeft
	FocusWindowBottomRight
	FocusWindowPrevious
	MoveWindowTop
	MoveWindowBottom
	MoveWindowLeft
	MoveWindowRight
	Suspend
	Quit
	QuitAll
	Write
	WriteQuit
	Info
	Error
)
