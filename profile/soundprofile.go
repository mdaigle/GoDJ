package profile

var MAX_TEMPO 		= 500
var MIN_TEMPO 		= 0
var MAX_DURATION 	= 3600
var MIN_DURATION 	= 0
var MAX_LOUDNESS 	= 100
var MIN_LOUDNESS 	= -100
var MAX_FLOAT 		= 1.0
var MIN_FLOAT 		= 0.0

type soundProfile struct {
	style 													[]string
	mood 													[]string
	song_type 												string
	max_tempo, 					min_tempo 					int
	max_duration, 				min_duration 				int
	max_loudness, 				min_loudness 				int
	artist_max_familiarity, 	artist_min_familiarity 		float32
	song_max_hotttnesss, 		song_min_hotttnesss 		float32
	max_danceability, 			min_danceability 			float32
	max_energy, 				min_energy 					float32
	max_liveness, 				min_liveness 				float32
	max_speechiness, 			min_speechiness 			float32
	sort string
}

func NewProfile() soundProfile {
	return soundProfile{
		max_tempo:				MAX_TEMPO,
		min_tempo:				MIN_TEMPO,
		max_duration:			MAX_DURATION,
		min_duration:			MIN_DURATION,
		max_loudness:			MAX_LOUDNESS,
		min_loudness:			MIN_LOUDNESS,
		artist_max_familiarity:	MAX_FLOAT,
		artist_min_familiarity:	MIN_FLOAT,
		song_max_hotttnesss:	MAX_FLOAT,
		song_min_hotttnesss:	MIN_FLOAT,
		max_danceability:		MAX_FLOAT,
		min_danceability:		MIN_FLOAT,
		max_energy:				MAX_FLOAT,
		min_energy:				MIN_FLOAT,
		max_liveness:			MAX_FLOAT,
		min_liveness:			MIN_FLOAT,
		max_speechiness:		MAX_FLOAT,
		min_speechiness:		MIN_FLOAT,
	};
}

func (field *int) setIntField(value int, ) {

}