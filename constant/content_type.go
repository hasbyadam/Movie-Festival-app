package constant

const (
	Image              string = "IMG"
	Video              string = "VID"
	MaxImageUploadSize        = 10 * 1024 * 1024
	MaxMovieUploadSize        = 100 * 1024 * 1024
)

var (
	AudioVideoTypes = []string{
		"video/mp4",
		"video/webm",
		"audio/aiff",
		"audio/mpeg",
		"application/ogg",
		"audio/midi",
		"video/avi",
		"audio/wave",
	}

	ImageTypes = []string{
		"image/x-icon",
		"image/bmp",
		"image/gif",
		"image/webp",
		"image/png",
		"image/jpeg",
	}
)
