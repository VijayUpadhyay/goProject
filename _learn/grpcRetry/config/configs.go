package retryconfig

// Time in sec
const (
	GrpcTimeout   = 20
	GrpcServerAdd = "localhost:50052"
	Sleep_Time    = 60
)

// Time in sec
const (
	MIN_CONNECT_TIMEOUT = 20
	INITIAL_BACKOFF     = 1
	MULTIPLIER          = 1.6
	MAX_BACKOFF         = 120
	JITTER              = 0.2
	MAX_RETRY           = 3
)
