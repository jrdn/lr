package lr

type Client interface {
	Close(objectHandle string) error

	AppVersionString() (string, error)
	ShowView(view View) error
	ShowSource(source Source) error

	GoBack() error
	GoForward() error
	SelectAll() error
	SelectNone() error
	NextPhoto() error
	PrevPhoto() error

	ZoomIn() error
	ZoomOut() error
	ZoomInSome() error
	ZoomOutSome() error
	Zoom100() error
	ZoomFit() error
	ZoomFill() error
	ZoomToggle() error

	SetFocusPoint(x, y float64, animate bool) error

	ParameterNames() ([]string, error)
	ParameterType(name string) (ParameterType, error)
	ParameterDefault(name string) (*ParameterValue, error)
	Increment(name string, amount float64) error
	Decrement(name string, amount float64) error
	GetValue(name string) (*ParameterValue, error)
	SetValue(name string, value *ParameterValue) error
	ResetToDefault(name string) error
	SetAutoTone() error

	RotateLeft() error
	RotateRight() error
	FlipHorizontal() error
	FlipVertical() error
	RotateAspectRatio() error
	ToggleConstrainAspectRatio() error

	Undo() error
	Redo() error

	FlagPick() error
	FlagReject() error
	FlagUnflag() error
	FlagIncrease() error
	FlagDecrease() error
	RatingIncrease() error
	RatingDecrease() error
	RatingClear() error
	Rating0() error
	Rating1() error
	Rating2() error
	Rating3() error
	Rating4() error
	Rating5() error

	OpenEnhance() error
	OpenHDRMerge() error
	OpenPanoramaMerge() error
	OpenHDRPanoramaMerge() error
	OpenExport() error
	EditInPhotoshop() error
	OpenPreferences() error

	ClearFilters() error
	ShowFilters() error
}
