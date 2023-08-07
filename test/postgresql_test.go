package main

import (
	"douyin-tiktok/common/utils"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresqlTest(t *testing.T) {
	// 测试正常连接的情况
	engine := utils.InitXorm("postgres", utils.Postgresql{Dsn: "postgres://postgres:123456@localhost:5432/douyin_tiktok?sslmode=disable"})
	defer engine.Close()
	assert.NotNil(t, engine, "Engine should not be nil")
	fmt.Println("Test for successful connection passed!")

	// 测试无法连接的情况
	invalidConfig := utils.Postgresql{
		Dsn: "invalid_dsn",
	}
	assert.Panics(t, func() {
		utils.InitXorm("postgres", invalidConfig)
	})
	fmt.Println("Test for failed connection passed!")
}
