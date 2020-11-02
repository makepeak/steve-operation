package config

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source"
	"github.com/micro/go-micro/v2/config/source/file"
	log "github.com/micro/go-micro/v2/logger"
)

var (
	err error
)

var (
	defaultRootPath         = "app"
	defaultConfigFilePrefix = "application-"
	etcdConfig              defaultEtcdConfig
	mysqlConfig             defaultMysqlConfig
	consulConfig            defaultConsulConfig
	profiles                defaultProfiles
	m                       sync.RWMutex
	inited                  bool
	sp                      = string(filepath.Separator)
)

// Init init config file
func Init() {
	m.Lock()
	defer m.Unlock()

	if inited {
		log.Info("[Init] config inited")
		return
	}

	// load config files
	// firt load basic config files
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("."+sp, sp)))

	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)

	// find application.yml file
	if err = config.Load(file.NewSource(file.WithPath(pt + sp + "application.yml"))); err != nil {
		panic(err)
	}

	// find profile infos
	if err = config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Infof("[Init] Load config file：path: %s, %+v\n", pt+sp+"application.yml", profiles)

	// load config files
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"

			log.Infof("[Init] Load config filePath: %s\n", filePath)

			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// load include files
		if err = config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// set files
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	config.Get(defaultRootPath, "consul").Scan(&consulConfig)
	// init done
	inited = true
}

// GetMysqlConfig get mysql config
func GetMysqlConfig() (ret MysqlConfig) {
	return mysqlConfig
}

// GetEtcdConfig get etcd config
func GetEtcdConfig() (ret EtcdConfig) {
	return etcdConfig
}

// GetConsulConfig get consul config
func GetConsulConfig() (ret ConsulConfig) {
	return consulConfig
}
