// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

/**
 * 数据库连接工具
 * @author 半城风雨
 * @since 2021/9/8
 * @File : db
 */
package utils

import (
	"easygoadmin/conf"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var XormDb *xorm.Engine

func init() {
	fmt.Println("初始化并连接数据库")

	// 获取配置实例
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=3s&parseTime=true", conf.CONFIG.Mysql.Username, conf.CONFIG.Mysql.Password, conf.CONFIG.Mysql.Host, conf.CONFIG.Mysql.Port, conf.CONFIG.Mysql.Database, conf.CONFIG.Mysql.Charset)
	XormDb, err = xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}

	// 通过engine.Ping()来进行数据库的连接测试是否可以连接到数据库。
	err = XormDb.Ping()
	if err == nil {
		fmt.Println("数据库连接成功")
		//关闭连接
		//defer XormDb.Close()
	} else {
		fmt.Printf("数据库连接错误:%v", err.Error())
		return
	}

	XormDb.DatabaseTZ = time.Local // 必须
	XormDb.TZLocation = time.Local // 必须
	// 设置连接池的空闲数大小
	XormDb.SetMaxIdleConns(10)
	// 设置最大打开连接数
	XormDb.SetMaxOpenConns(30)

	// 结构体与数据表的映射
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	XormDb.SetTableMapper(tbMapper)

	// 开启调试模式和打印日志,会在控制台打印执行的sql
	if conf.CONFIG.Mysql.Debug {
		XormDb.ShowSQL(conf.CONFIG.Mysql.Debug)
		XormDb.Logger().SetLevel(core.LOG_DEBUG)
	}
}
