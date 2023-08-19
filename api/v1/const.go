package v1

type ReportPhase string

var (
	ReportPhaseBackingUp ReportPhase = "BackingUp"
	ReportPhaseCompleted ReportPhase = "Completed"
	ReportPhaseFailed    ReportPhase = "Failed"
)
