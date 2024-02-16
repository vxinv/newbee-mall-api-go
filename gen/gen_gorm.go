package gen

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func Dsn(username, password, host, dbname string) string {
	return username + ":" + password + "@tcp(" + host + ":" + "3306" + ")/" + dbname + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
}

func Gen_gorm() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		// 如果你希望为可为null的字段生成属性为指针类型, 设置 FieldNullable 为 true
		FieldNullable: true,
		// 如果你希望在 `Create` API 中为字段分配默认值, 设置 FieldCoverable 为 true, 参考: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// 如果你希望生成无符号整数类型字段, 设置 FieldSignable 为 true
		FieldSignable: true,
		// 如果你希望从数据库生成索引标记, 设置 FieldWithIndexTag 为 true
		FieldWithIndexTag: true,
		// 如果你希望从数据库生成类型标记, 设置 FieldWithTypeTag 为 true
		FieldWithTypeTag: true,
		// 如果你需要对查询代码进行单元测试, 设置 WithUnitTest 为 true
		WithUnitTest: true,
	})

	gormdb, _ := gorm.Open(mysql.Open(Dsn("root", "WdRedis5188@@", "localhost", "yami_shops")))
	g.UseDB(gormdb) // reuse your gorm db

	//// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})
	//
	//// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})
	g.ApplyBasic(
		// 从当前数据库中生成所有表的结构
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
