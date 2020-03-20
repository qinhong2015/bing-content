package Service

import (
	"FeedImport/db"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type accountModel struct {
	accountPdo *sql.DB
	AccountName, AccountId, AccountSource, MysqlHost, Timezone string
	AdvanceModifierEnabled, ModifyPromotionIdEnabled, ModifyGoogleExpressEnabled bool
	Config map[string]interface{}
	settings map[string]interface{}
}

type settingModel struct {
	id string
	value interface{}
}

func (a *accountModel) GetAccountPdo() *sql.DB{
	if a.accountPdo == nil {
		host := a.MysqlHost
		dbs := a.Config["dbs"].(map[string]interface{})[host]
		a.accountPdo = db.CreateInstance(dbs, a.AccountId)
	}
	return a.accountPdo
}

func (a *accountModel) LoadSettings() {
	rows, err := a.GetAccountPdo().Query("select `id`, `value` from `settings`")
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
	a.settings = settingsMap
}

func (a *accountModel) GetSettings(name string) interface{} {
	if val, ok := a.settings[name]; ok {
		return val
	}
	return nil;
}