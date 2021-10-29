package week04

var (
	_ Entertainer = &SpoiledBand{}
	_ Setuper     = &SpoiledBand{}
)

// SpoiledBand is a type of entertainer
// who requires a ton of things from the organizers
// and the venue, hence his performance requires a
// lot of preparation.
type SpoiledBand struct{}

func (spo SpoiledBand) Name() string {
	return "Spoiled Band"
}

func (spo SpoiledBand) Setup(v Venue) {
	v.Log.Write([]byte("Spoiled Band: Preparing for the show\n"))
	v.Log.Write([]byte("Spoiled Band: Artists want flowers"))
	v.Log.Write([]byte("Spoiled Band: Wants a Carl Sagan Picture"))
	v.Log.Write([]byte("Spoiled Band: They want a 73F degree room, working on it"))
	v.Log.Write([]byte("Spoiled Band: Finding an air-hockey table for the band"))
	v.Log.Write([]byte("Spoiled Band: Adding 3 Mangos, 3 Hawaiian Papaya, 6 Bananas and 3 Peaches"))
}

func (spo SpoiledBand) Perform(v Venue) {
	v.Log.Write([]byte("Spoiled Band: Performing 3 only songs they have\n"))
	v.Log.Write([]byte("Spoiled Band: Singing first song\n"))
	v.Log.Write([]byte("Spoiled Band: Singing seccond song\n"))
	v.Log.Write([]byte("Spoiled Band: Singing (with performance difficulties) third song\n"))
}
