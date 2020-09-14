module api

go 1.15

replace github.com/fang/video-streaming-media/defs => ./defs

replace github.com/fang/video-streaming-media/utils => ./utils

require (
	github.com/fang/video-streaming-media/defs v0.0.0-00010101000000-000000000000
	github.com/fang/video-streaming-media/utils v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/julienschmidt/httprouter v1.3.0
)
