package yconf

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Conf Config

type Config struct {
	KafkaServer  []string `yaml:"kafka_server"`
	KafkaGroup   string   `yaml:"kafka_group"`
	KafkaTopicId string   `yaml:"kafka_topic_id"`

	RedisAddr     string `yaml:"redis_addr"`
	RedisPassword string `yaml:"redis_password"`
	RedisDB       int    `yaml:"redis_db"`

	PgHost     string `yaml:"pg_host"`
	PgPort     string `yaml:"pg_port"`
	PgUser     string `yaml:"pg_user"`
	PgPassword string `yaml:"pg_password"`
	PgDbName   string `yaml:"pg_db_name"`

	DepthServer string `yaml:"depth_server"`
}

func InitSetting() error {

	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		return err
	}

	Conf.KafkaServer = make([]string, 0)

	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		return err
	}

	if len(Conf.KafkaServer) == 0 {
		return errors.New("len(Conf.KafkaServer) == 0")
	}
	return nil
}
