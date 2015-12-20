package profile

var MAX_TEMPO 		= 500
var MIN_TEMPO 		= 0
var MAX_DURATION 	= 3600
var MIN_DURATION 	= 0
var MAX_LOUDNESS 	= 100
var MIN_LOUDNESS 	= -100
var MAX_FLOAT 		= 1.0
var MIN_FLOAT 		= 0.0

var INCREMENT_STEPS = 16

type SoundProfile struct {
	Style 													[]string
	Mood 													[]string
	Song_type 												string
	Max_tempo, 					Min_tempo 					int
	Max_duration, 				Min_duration 				int
	Max_loudness, 				Min_loudness 				int
	Artist_max_familiarity, 	Artist_min_familiarity 		float64
	Song_max_hotttnesss, 		Song_min_hotttnesss 		float64
	Max_danceability, 			Min_danceability 			float64
	Max_energy, 				Min_energy 					float64
	Max_liveness, 				Min_liveness 				float64
	Max_speechiness, 			Min_speechiness 			float64
	Sort 													string
}

func NewProfile() SoundProfile {
	return SoundProfile{
		Max_tempo:				MAX_TEMPO,
		Min_tempo:				MIN_TEMPO,
		Max_duration:			MAX_DURATION,
		Min_duration:			MIN_DURATION,
		Max_loudness:			MAX_LOUDNESS,
		Min_loudness:			MIN_LOUDNESS,
		Artist_max_familiarity:	MAX_FLOAT,
		Artist_min_familiarity:	MIN_FLOAT,
		Song_max_hotttnesss:	MAX_FLOAT,
		Song_min_hotttnesss:	MIN_FLOAT,
		Max_danceability:		MAX_FLOAT,
		Min_danceability:		MIN_FLOAT,
		Max_energy:				MAX_FLOAT,
		Min_energy:				MIN_FLOAT,
		Max_liveness:			MAX_FLOAT,
		Min_liveness:			MIN_FLOAT,
		Max_speechiness:		MAX_FLOAT,
		Min_speechiness:		MIN_FLOAT,
	};
}

func incrementFloat(field float64) {
	field = field + MAX_FLOAT/float64(INCREMENT_STEPS)
}

func incrementInt(field int, max, min int) {
	field = field + (max - min)/INCREMENT_STEPS
}