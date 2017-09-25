package db

import (
	"errors"
	"fmt"
	"os"
	//"reflect"

	"github.com/asdine/storm"
	"github.com/boltdb/bolt"

	"github.com/qor/admin"
	"github.com/qor/qor"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/roscopecoltran/krakend/config"
	// "github.com/roscopecoltran/krakend/plugins/openapi"
	// "github.com/k0kubun/pp"
)

var (
	RDB   *gorm.DB
	KVS   *bolt.DB
	STORM *storm.DB
	ADMIN *admin.Admin
)

type RDBStorage struct {
	RDB   *gorm.DB
	Ready bool
}

type KVSStorage struct {
	Bolt struct {
		Client *bolt.DB
		Ready  bool
	}
	Storm struct {
		Client *storm.DB
		Ready  bool
	}
}

type AdminHelper struct {
	UI    *admin.Admin
	Ready bool
}

func NewStack() error {
	// pp.Println(config.Config)
	dbConfig := config.Config.Datastore.RDB
	if dbConfig.Disabled == false {
		if err := InitGorm(); err != nil {
			return err
		}
	}
	kvConfig := config.Config.Datastore.KVS
	if kvConfig.Disabled == false {
		if kvConfig.Adapter == "boltdb" {
			if err := InitBolt(); err != nil {
				return err
			}
		}
	}
	return nil
	// app := GatewayCluster{}
	// app.initGatewayCluster()
}

func InitGorm() error {
	var err error
	dbConfig := config.Config.Datastore.RDB
	if dbConfig.Disabled == false {
		if dbConfig.Adapter == "mysql" {
			RDB, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		} else if dbConfig.Adapter == "postgres" {
			RDB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		} else if dbConfig.Adapter == "sqlite" {
			if err := os.MkdirAll(dbConfig.PrefixPath, 0777); err != nil {
				return err
			}
			RDB, err = gorm.Open("sqlite3", fmt.Sprintf("%v/%v.db", dbConfig.PrefixPath, dbConfig.Name))
		} else {
			return errors.New("Not supported database adapter")
		}
		if err == nil {
			if dbConfig.Debug {
				RDB.LogMode(true)
			}
			// RDB.SetLogger(gormrus.New())

			// RDB.DB().SetMaxIdleConns(dbConfig.MaxIdleConns)
			// This should be about the number of cores the machine has (and should
			// be lower than the max_connection specified by postgresql.conf)
			// Since we have two applications running, this should really be
			// number of cores / 2
			// TODO Make configurable
			// RDB.DB().SetMaxOpenConns(dbConfig.MaxOpenConns)

		} else {
			return err
		}
		if RDB.DB().Ping() != nil {
			return errors.New("Could not access to RDB database")
		}
	}
	return nil
}

func GetRDB() *gorm.DB {
	return RDB
}

func MigrateDB(isAdminUI bool, isAutoMigrate bool, isTruncate bool, tables ...interface{}) error {
	if isAdminUI {
		ADMIN = admin.New(&qor.Config{DB: RDB})
	}
	for _, table := range tables {
		if isTruncate {
			// pp.Println(table)
			if err := RDB.DropTableIfExists(table).Error; err != nil {
				fmt.Println("var.isTruncate", isTruncate)
				fmt.Println("var.table", table)
				return err
			}
		}
		if isAutoMigrate {
			fmt.Println("isAutoMigrate, new table: ", table)
			// pp.Println(table)
			if err := RDB.AutoMigrate(table).Error; err != nil {
				// pp.Println(err)
				fmt.Println("var.isAutoMigrate", isAutoMigrate)
				fmt.Println("var.table", table)
				return err
			}
		}
		if isAdminUI {
			fmt.Println("isAdminUI, creating hook for table: ", table)
			ADMIN.AddResource(table)
			fmt.Println("var.isAdminUI", isAdminUI)
			fmt.Println("var.resource", table)
		}
	}
	if isAdminUI {
		if len(ADMIN.GetResources()) > 0 {
			for _, resource := range ADMIN.GetResources() {
				fmt.Println("var.adminui.resource", resource)
			}
		}
	}
	return nil
}

// InitBolt init bolt
func InitBolt() error {
	kvsConfig := config.Config.Datastore.KVS
	if kvsConfig.Disabled == false {
		var err error
		if err := os.MkdirAll(kvsConfig.PrefixPath, 0777); err != nil {
			return err
		}
		KVS, err = bolt.Open(fmt.Sprintf("%v/%v.boltdb", kvsConfig.PrefixPath, kvsConfig.Name), 0600, nil)
		if err != nil {
			return err
		}
		return KVS.Update(func(tx *bolt.Tx) error {
			if _, err = tx.CreateBucketIfNotExists([]byte("koip")); err != nil {
				return err
			}
			return nil
		})
	}
	return nil
}

func GetKVS() *bolt.DB {
	return KVS
}

func (kvss *KVSStorage) InitStorm(driver string) (*KVSStorage, error) {
	var err error
	kvsConfig := config.Config.Datastore.KVS
	if err := os.MkdirAll(kvsConfig.PrefixPath, 0777); err != nil {
		return kvss, err
	}
	dbFilePath := fmt.Sprintf("%v/%v.boltdb", kvsConfig.PrefixPath, kvsConfig.Name)
	switch drv := driver; drv {
	case "storm":
		kvss.Storm.Client, err = storm.Open(dbFilePath, storm.AutoIncrement())
		if err != nil {
			return kvss, err
		}
		kvss.Storm.Ready = true
	case "bolt":
	case "default":
		kvss.Bolt.Client, err = bolt.Open(dbFilePath, 0600, nil)
		if err != nil {
			return kvss, err
		}
		kvss.Bolt.Ready = true
	}
	return kvss, nil
}

func (kvss *KVSStorage) Close(driver string) error {
	switch drv := driver; drv {
	case "storm":
		kvss.Storm.Ready = false
		err := kvss.Storm.Client.Close()
		if err != nil {
			return err
		}
	case "bolt":
	case "default":
		kvss.Bolt.Ready = false
		err := kvss.Bolt.Client.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

/*
func GetDB() *dynamodb.DynamoDB {
	return db
}
*/

/*
// FindServices finds all services
func FindServices(Database *gorm.DB, match string) ([]Service, error) {
	var services []Service
	if match != "" {
		db.Where("service LIKE ?",
			strings.ToLower(fmt.Sprintf("%%%s%%", match))).Order("service").Find(&services)
	} else {
		db.Order("service").Find(&services)
	}
	return services, db.Error
}
*/
