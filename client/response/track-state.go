package response

type TrackState struct {
	LeaveMarbles         bool `json:"leave_marbles"`
	PracticeRubber       int  `json:"practice_rubber"`
	QualifyRubber        int  `json:"qualify_rubber"`
	WarmupRubber         int  `json:"warmup_rubber"`
	RaceRubber           int  `json:"race_rubber"`
	PracticeGripCompound int  `json:"practice_grip_compound"`
	QualifyGripCompound  int  `json:"qualify_grip_compound"`
	WarmupGripCompound   int  `json:"warmup_grip_compound"`
	RaceGripCompound     int  `json:"race_grip_compound"`
}
