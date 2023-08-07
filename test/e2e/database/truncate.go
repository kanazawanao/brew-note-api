package database

func TruncateTable() {
	rows, err := DB.Raw("show tables").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			panic(err.Error())
		}
		// migrations テーブルは削除しない
		if table != "migrations" {
			// 外部キー制約に関係なくレコードの中身を空にするため、一時的に外部キー制約を無効化
			DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
			DB.Exec("TRUNCATE TABLE coffee-paws." + table)
			DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
		}
	}
}