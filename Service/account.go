package Service

import (
	"FeedImport/db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type accountModel struct {
	accountPdo *sql.DB
	accountName, accountId, accountSource, mysqlHost, timezone string
	advanceModifierEnabled, modifyPromotionIdEnabled, modifyGoogleExpressEnabled bool
	config map[string]interface{}
	settings map[string]interface{}
}

type settingModel struct {
	id string
	value interface{}
}

func (this *accountModel) GetAccountPdo() *sql.DB{
	if this.accountPdo == nil {
		host := this.mysqlHost
		dbs := this.config["dbs"].(map[string]interface{})[host]
		this.accountPdo = db.CreateInstance(dbs, this.accountId)
	}
	return this.accountPdo
}

func (this *accountModel) LoadSettings() {
	rows, err := this.GetAccountPdo().Query("select `id`, `value` from `settings`")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var setting settingModel
	settingsMap := make(map[string]interface{})
	for rows.Next() {
		err = rows.Scan(&setting.id, &setting.value)

		if err != nil {
			panic(err.Error())
		}

		settingsMap[setting.id] = setting.value
	}
	this.settings = settingsMap
}

func (this *accountModel) getSettings(name string) interface{} {
	return this.settings[name];
}