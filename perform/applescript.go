package perform

import (
	"github.com/everdev/mack"
	//"strconv"
)

func PlayPause() {
	//mack.Say("Playing")
	mack.Tell("Spotify", "playpause")
}

func Next() {
	mack.Tell("Spotify", "next track")
}

func Previous() {
	mack.Tell("Spotify", "previous track")
}

func IncreaseVolume() {
	mack.Tell("Spotify", "set sound volume to sound volume + 5")
}

func DecreaseVolume() {
	mack.Tell("Spotify", "set sound volume to sound volume - 5")
}
