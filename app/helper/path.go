package helper

import (
	"github.com/koushamad/goro-core/app/conf"
	"path"
	"path/filepath"
	"runtime"
	"time"
)

func Path(Path string) string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(path.Dir(path.Dir(path.Dir(b)))))
	return filepath.Dir(filepath.Dir(filepath.Dir(d))) + Path
}

func RootPath() string {
	return Path("")
}

func StoragePath() string {
	return Path("/storage")
}

func LogPath() string {
	return "storage/log/" + time.Now().Format("2006-01-02") + "-goro.log"
}

func CachePath() string {
	return Path("/storage/cache")
}

func SqlitePatch() string {
	return Path("/storage/database/" + conf.Get("db.driver"))
}
