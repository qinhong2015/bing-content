package Service

import (
	"database/sql"
)

type accountService struct {
	account accountModel
}

var accountInstance *accountService

func GetAccountService() *accountService {
	if accountInstance == nil {
		accountInstance = new(accountService)
	}
	//db.GetPdo()
	return accountInstance
}

func (this *accountService) SetAccount(pdo *sql.DB, accountId string, config map[string]interface{}) {
	rows, err := pdo.Query("SELECT `a`.`name`, `a`.`account_id`, `a`.`source`, `a`.`mysql_host`, `a`.`timezone`, `a`.`source`, CASE WHEN ap.permissions LIKE '%\"product_modify_advance_modifier\": true%' THEN 1 ELSE 0 END AS `advance_modifier_enabled`, CASE WHEN ap.permissions LIKE '%\"product_modify_promotion_id\": true%' THEN 1 ELSE 0 END AS `modify_promotion_id_enabled`, CASE WHEN ap.permissions LIKE '%\"product_modify_google_express\": true%' THEN 1 ELSE 0 END AS `product_modify_google_express_enabled` FROM `accounts` AS `a` INNER JOIN `account_permissions` AS `ap` ON `ap`.`account_id` = `a`.`account_id` WHERE `a`.`account_id` = ?;", accountId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	var account accountModel

	accountFound := rows.Next()
	if !accountFound {
		panic("Account not found")
	}

	err = rows.Scan(
		&account.accountName,
		&account.accountId,
		&account.accountSource,
		&account.mysqlHost,
		&account.timezone,
		&account.accountSource,
		&account.advanceModifierEnabled,
		&account.modifyGoogleExpressEnabled,
		&account.modifyPromotionIdEnabled,
	)

	if err != nil {
		panic(err.Error())
	}

	account.config = config
	account.LoadSettings()
	this.account = account
}