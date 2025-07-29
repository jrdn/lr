# lr

Client for interacting with Lightroom CC's API

file:///private/var/folders/6x/vkmjcvrn0mn0nqz30d7w48n40000gn/T/LrControllerAPI_8.3.1.html

file:///private/var/folders/6x/vkmjcvrn0mn0nqz30d7w48n40000gn/T/LrControllerAPI_8.3.1_Setup.html

# API Docs

appVersionString()

        Retrieves the current version of the application as a user-displayable string
        (for instance, "8.0").

        returns(string)     Lightroom version

appVersionTable()

        Retrieves the current version of the application as a table that can be parsed.
        For example, a version such as 8.1.4 contains three components, the major version,
        the minor version, and the revision number of a specific release. There is also
        be a build number(usually a 6-digit number).

        returns(table) Lightroom version. Contains the following members:
            • major(number) Major version number
            • minor(number) Minor version number
            • revision(number) Revision number
            • build(number) Build number(usually a six-digit number)

showView()

        Switches the app's view mode to one of: "loupe", "square_grid", "compare", "dynamic_grid"

showSource()

        Selects the given source, one of: "all_photos", "trash"

goBack()

        Navigates back to the previous view mode or photo source

goForward()

        Navigates forward to the next view mode or photo source

selectAll()

        Selects all photos in the grid.

selectNone()

        Deselects all photos.

nextPhoto()

        Selects the next photo.

previousPhoto()

        Selects the previous photo.

zoomIn()

        Zooms in to the last level used.

zoomOut()

        Zooms out to the last level used.

zoomToOneToOne()

        Zooms to 100%.

zoomToFit()

        Sets zoom mode to Fit, making the photo as large as possible while keeping all of it visible.

zoomToFill()

        Sets zoom mode to Fill, making the photo completely fill the view.

zoomInSome()

        Increases the zoom factor one step.

zoomOutSome()

        Decreases the zoom factor one step.

toggleZoom()

        Toggles between zoomed in and zoomed out.

setFocusPoint()

        Pans the photo to the given focus point.
        Arguments:
            x        (number)    New x coordinate as a percentage of the photo width [0-1].
                                If nil, the x position is unchanged.
            y        (number)    New y coordinate as a percentage of the photo height [0-1].
                                If nil, the y position is unchanged.
            animate    (boolean)    If true, smoothly animates to the new location.

revealPanel()

revealAdjustedControls()

        Enables a mode where adjusting a parameter causes that panel to be automatically revealed in the panel track.
        Argument:     (bool)    true to enable reveal behavior, false to disable it (the default mode)

setAutoTone()

        Applies Auto Tone to the current photo.

toggleHDR()

        Toggles HDR on/off for the current photo.

isHDROn()

        Returns true if HDR on for the current photo.

toggleBlackAndWhite()

        Toggles B&W conversion on/off for the current photo.

isBlackAndWhiteOn()

        Returns true if B&W conversion is on for the current photo.

resetToOpen()

        Reverts all changes made during this edit session.

resetToDefault()

        Resets a single adjustment for the current photo to its default value.
        Argument: (string) a parameter name

resetAllDevelopAdjustments()

        Resets all adjustments for the current photo.

resetCrop()

        Resets the crop angle and frame for the current photo.

resetSpotRemoval()

        Clears all heal/remove adjustments from the current photo.

resetRedeye()

        Clears all redeye removal adjustments from the current photo.

getParameterNames()

        Returns an ordered list of all parameter names.

getParameterType()

        Returns a parameter's type, one of "boolean", "number", "string".
        Argument:    (string)    Name of a parameter (ex: "Exposure2012").

getPanelNames()

        Returns an ordered list of all panel names.

getRange()

        Gets the min and max value, precision, and increment sizes of a numeric adjustment. If percision is zero,
        the value is an integer.

        Note: the values of the following parameters have special meanings:

        "PerspectiveUpright" :
            0 = Off
            1 = Auto
            2 = Full
            3 = Level
            4 = Vertical
            5 = Guided

        "HDREditMode" :
            0 = Off
            1 = On

        Argument:    (string)    Name of a parameter (ex: "Exposure2012").
        Returns:     (number, number, number, number, number) The min, max, precision, small increment, and large increment for the given parameter.

getDefault()

        Returns the default value for a parameter.

getLabel()

        Returns the localized user-visible string describing a parameter.
        Argument:    (string)    Name of a parameter (ex: "Exposure2012").
        Returns:     (string)    Label for the parameter.

setValue()

        Sets the value of an adjustment for the current photo.
        Arguments:
            parameter    (string)    Name of a parameter (ex: "Exposure2012").
            value         (number)     The new value for the parameter.

getValue()

        Gets the current value of an adjustment for the current photo.
        Argument:    (string)    Name of a parameter (ex: "Exposure2012").

increment()

        Increments the value of an adjustment for the current photo.
        Arguments:
            parameter    (string)    Name of a parameter (ex: "Exposure2012").
            amount         (number)    Optional amount by which to increment. If omitted, a default amount will be used.

decrement()

        Decrements the value of an adjustment for the current photo.
        Argument:
            parameter    (string)    Name of a parameter (ex: "Exposure2012").
            amount         (number)    Optional amount by which to decrement. If omitted, a default amount will be used.

startTracking()

        Temporarily puts the Develop module into its tracking state, causing faster, lower-quailty redraw
        and preventing history states from being generated. Tracking will automatically be turned back
        off as soon as a diferent parameter is adjusted, or two seconds after the last adjustment is made.
        Argument:     (string)    a parameter name

stopTracking()

        Causes Develop module to exit its tracking state immediately, creating a single history state
        for all changes that were made to the parameter that was being tracked.

setTrackingDelay()

        Sets the number of seconds that tracking remains enabled after each adjustment is made.
        The default is 2 seconds.
        Argument:     (number)     number of seconds to wait before turning off tracking

setTrackingModifiers()

        Sets/clears modifier keys that may affect the behavior of the tracking parameter (for
        example, to enable the slider visualization shown when the Alt/Option key is held down).
        Argument:    (string)    If "option" then that modifier key is enabled. If not, all tracking modifiers are cleared.

setMultipleAdjustmentThreshold()

        Sets the time threshold that determines when adjustments to different parameters will be grouped
        together into a single history state versus recorded separately. If multiple different parameters
        are changed within a window of time that is less than this threshold, they will be grouped
        together into a single "Multiple Settings" history state. If they occur farther apart than this
        threshold, each will get its own history state. Recording many separate history states in rapid
        succession can degrade Lightroom's performance, so this threshold is very important if simultaneous
        adjustments by the user are a possibility.

        The default is 0.5 seconds. Has no effect if set to a value that is higher than the tracking delay.
        Argument:         (number)    time threshold under which multiple adjustments will be grouped together

addAdjustmentChangeObserver( optParam )

        Registers an observer that will be called any time the develop adjustments changes.
        Argument: (string) A parameter name to observe. If this is omitted, all parameters are observed.
        Returns: (string) The GUID of an object that can be closed when you do not need the observer any more.

rotateLeft()

        Rotates the image 90 degrees counter-clockwise.

rotateRight()

        Rotates the image 90 degrees clockwise.

flipVertical()

        Flips the image vertically.

flipHorizontal()

        Mirrors the image horizontally.

rotateAspectRatio()

        Rotates the aspect ratio between portrait and landscape.
 
toggleConstrainAspectRatio()

        Toggles aspect ratio constraint on/off.

isConstrainAspectRatioOn()

        Returns true if constrain aspect ratio is on.

setCropAspectRatio()

        Sets the crop aspect ratio.
        Argument:    (number or string) "original", "as_shot", "custom", or a number greater than 0 (width/height).

getCropAspectRatio()

        Gets the aspect ratio of crop used when constraint is enabled.
        Returns: (number, string) the crop aspect ratio, and "original" or "as_shot" or "custom".

selectTool()

        Activate a tool mode.
        Argument:    (string) new tool mode, one of: "loupe", "crop", "heal", "whiteBalance", "masking", "upright".

getSelectedTool()

        Reports which tool mode is active.
        Returns: (string) current tool mode, one of: "loupe", "crop", "heal", "whiteBalance", "masking", "upright".

showClipping()

        Toggles clipping indicators on/off.

isClippingShowing()

        Returns true if clipping indicators are enabled.

editInPhotoshop()

        Opens the current photo for editing in Photoshop.

openExport()

        Opens the Export dialog for the selected photos.

canExportWithPrevious()

        Returns true if the selected photos can be exported using the last export settings.

exportWithPrevious()

        Exports the selected photos using the last export settings.

openEnhance()

        Opens the Enhance dialog for the selected photo.

openHDRMerge()

        Opens the HDR Merge dialog for the selected photos.

openPanoramaMerge()

        Opens the Panorama Merge dialog for the selected photos.

openHDRPanoramaMerge()

        Opens the HDR Panorama dialog for the selected photos.

clearSearch()

        Clears the search criteria.

clearFilters()

        Clears the search filters.

showFilters()

        Shows/hides the search filters.

setRatingFilter()

        Sets the rating filter, [0-5].

setPickFilter()

        Sets the pick filter, one of: "flagged", "rejected", "unflagged", "unrejected", "nonneutral", "nonflagged".

openPreferences()

        Opens the preferences dialog.

copyEditSettings()

        Copies Edit settings to the clipboard.

showCopyEditSettings()

        Shows the dialog for choosing which Edit settings to copy.

canPasteEditSettings()

        Returns trus if there are Edit settings on the clipboard.

pasteEditSettings()

        Pastes Edit settings from the clipboard.

undo()

        Undoes the last operation.

redo()

        Redoes the last undone operation.

flagPick()

        Sets the flag to 'picked'

flagReject()

        Sets the flag to 'rejected'

flagUnflag()

        Sets the flag to 'unflagged'

flagIncrease()

        Increases the flag ('unflagged' to 'picked' or 'rejected' to 'unflagged')

flagDecrease()

        Decreases the flag ('picked' to 'unflagged' or 'unflagged' to 'rejected')

flagPickToggle()

        Toggles the flag between 'picked' and 'unflagged'

flagRejectToggle()

        Toggles the flag between 'rejected' and 'unflagged'

getFlag()

        Returns the flag state, one of: 'picked', 'unflagged', or 'rejected')

rating0()

        Sets the rating to 0

rating1()

        Sets the rating to 1

rating2()

        Sets the rating to 2

rating3()

        Sets the rating to 3

rating4()

        Sets the rating to 4

rating5()

        Sets the rating to 5

ratingIncrease()

        Increases the rating

ratingDecrease()

        Decreases the rating

getRating()

        Returns the rating, [0-5]

addHapticsObserver()

        Registers an observer that will be notified about haptic events. Each notification
        will contain the following fields:

            category        One of the following:
                                "slider"    a slider handle is being dragged
                                "crop"        a crop handle is being dragged
                                "dialog"    a dialog has opened/closed


            event            If category is "dialog" then event will be "opened" or "closed".

                            If category is "slider" or "crop" then event will be one of the following:
                                "hit"        sent once when the drag first reaches a side.
                                "repeat"    sent repeatedly if dragging continues past the end.


            side            If category is "slider" then side will be one of the following:
                                "min"        sent when the left side of the slider track is hit
                                "default"    sent when the slider crosses over its default value
                                "max".        sent when the right side of the slider track is hit

                            If category is "crop" then side will be one of the following:
                                "left", "top", "right", "bottom", "top-left",
                                "top-right", "bottom-left", "bottom-right".


            id                The id of the slider/dialog.

        Returns: An observer object on which you can call "close" when you don't need the observer any more.

# Setup Docs

Lightroom Controller API Setup

Communication between 3rd party controller apps and Lightroom happens over a WebSocket connection.
Getting started
You will first need to enable the feature. To do this, navigate to the Interface tab in the Preferences and make sure the “Enable external controllers” checkbox is ticked. Restart Lightroom.
Lightroom WebSocket Server
When the feature is enabled, Lightroom will start a WebSocket server at ws://127.0.0.1:7682 by default. When a client connects, Lightroom will show a dialog confirming if the user recognizes this application, with the option to remember it across launches.

If the port is not available, Lightroom falls back to a different port. The port that Lightroom listens on will be saved into the file connections.json file at this location:

    Mac: ~/Library/Application Support/Adobe/Lightroom CC/Connections
    Win: AppData\Local\Adobe\Lightroom CC\Connections

Prototype for Communication
The client can make requests to Lightroom via messages. Here is the protocol for a message:

```json
{
	"requestId": "<guid>",
	"object": "<object handle>",
	"message": "<operation>",
	"params": [
	<param1>,
	...
	]
}
```

The requestId is a new guid string for every request. Response messages that are sent back from the Lightroom server will refer this guid. The object handle is the handle to the object that the client wants to send the message to. By default, messages are sent to the root object for which this parameter can be omitted.

Here is the protocol for the response message:

```json
{
	"requestId": "<request guid>",
	"success": <boolean>,
	"response": {
		<return_value_1>,
		...
	}
}
```

Each request can be served by one or many responses.
Initializing the connection
First call from the plugin should be:

```json
{
  "requestId": "<guid>",
  "object": null,
  "message": "register",
  "params": ["<appName>", "<appVersion>", "<optional_clientGUID>"]
}
```

Lightroom should reply with:

```json
{
	"requestId": "<request guid>",
	"success": true,
	"response": {
		<clientGUID>,
		<Lightroom version>
	}
}
```

When the client initiates the registration process, Lightroom will show a dialog to the user to confirm whether they want to allow it. If allowed, Lightroom will return a GUID to the client. On subsequent connections, if the same GUID is passed back into the register call, then the confirmation dialog will not be shown.
Closing
After the 3rd party app is done using any object, it can close it by calling the following:

```json
{
  "requestId": "<guid>",
  "object": "<object handle>",
  "message": "close",
  "params": []
}
```

# Parameters

AutoGrayscaleMix
AutoLateralCA
Blacks2012
BlueHue
BlueSaturation
CameraProfile
Clarity2012
ColorGradeBlending
ColorNoiseReduction
ColorNoiseReductionDetail
ColorNoiseReductionSmoothness
Contrast2012
CorrectionAmount
CurveRefineSaturation
DefringeGreenAmount
DefringePurpleAmount
Dehaze
DepthCorrectionAmount
DepthSource
Exposure2012
GrainAmount
GrainFrequency
GrainSize
GrayMixerAqua
GrayMixerBlue
GrayMixerGreen
GrayMixerMagenta
GrayMixerOrange
GrayMixerPurple
GrayMixerRed
GrayMixerYellow
GreenHue
GreenSaturation
HDREditMode
HDRMaxValue
Highlights2012
HueAdjustmentAqua
HueAdjustmentBlue
HueAdjustmentGreen
HueAdjustmentMagenta
HueAdjustmentOrange
HueAdjustmentPurple
HueAdjustmentRed
HueAdjustmentYellow
LensBlur.Active
LensBlur.BlurAmount
LensBlur.CatEyeAmount
LensBlur.HighlightsBoost
LensBlur.ShowOverlay
LensProfileDistortionScale
LensProfileEnable
LensProfileVignettingScale
LocalBlacks2012
LocalClarity2012
LocalContrast2012
LocalCurveRefineSaturation
LocalDefringe
LocalDehaze
LocalExposure2012
LocalGrain
LocalHighlights2012
LocalLuminanceNoise
LocalMoire
LocalSaturation
LocalShadows2012
LocalSharpness
LocalTemperature
LocalTexture
LocalTint
LocalWhites2012
Look.Amount
LuminanceAdjustmentAqua
LuminanceAdjustmentBlue
LuminanceAdjustmentGreen
LuminanceAdjustmentMagenta
LuminanceAdjustmentOrange
LuminanceAdjustmentPurple
LuminanceAdjustmentRed
LuminanceAdjustmentYellow
LuminanceNoiseReductionContrast
LuminanceNoiseReductionDetail
LuminanceSmoothing
OverrideLookVignette
PerspectiveUpright
PointColorHueShift
PointColorLumScale
PointColorRangeAmount
PointColorSatScale
PointColorVisualizeRange
PostCropVignetteAmount
PostCropVignetteFeather
PostCropVignetteHighlightContrast
PostCropVignetteMidpoint
PostCropVignetteRoundness
RedHue
RedSaturation
SDRBlend
SDRBrightness
SDRClarity
SDRContrast
SDRHighlights
SDRShadows
SDRWhites
Saturation
SaturationAdjustmentAqua
SaturationAdjustmentBlue
SaturationAdjustmentGreen
SaturationAdjustmentMagenta
SaturationAdjustmentOrange
SaturationAdjustmentPurple
SaturationAdjustmentRed
SaturationAdjustmentYellow
ShadowTint
Shadows2012
SharpenDetail
SharpenEdgeMasking
SharpenRadius
Sharpness
SplitToningBalance
Temperature
Texture
Tint
Vibrance
WhiteBalance
Whites2012
cropBottom
cropLeft
cropRight
cropTop
previewForSDRDisplay
straightenAngle
visualizeHDRRanges
