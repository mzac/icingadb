package timeperiodrange

import (
	"git.icinga.com/icingadb/icingadb-main/configobject"
	"git.icinga.com/icingadb/icingadb-main/connection"
	"git.icinga.com/icingadb/icingadb-main/utils"
)

var (
	ObjectInformation configobject.ObjectInformation
	Fields         = []string{
		"id",
		"timeperiod_id",
		"range_key",
		"range_value",
		"env_id",
	}
)

type TimeperiodRange struct {
	Id						string 		`json:"id"`
	TimeperiodId			string		`json:"timeperiod_id"`
	RangeKey	 			string 		`json:"range_key"`
	RangeValue	 			string 		`json:"range_value"`
	EnvId           		string		`json:"env_id"`
}

func NewTimeperiodRange() connection.Row {
	t := TimeperiodRange{}
	return &t
}

func (t *TimeperiodRange) InsertValues() []interface{} {
	v := t.UpdateValues()

	return append([]interface{}{utils.Checksum(t.Id)}, v...)
}

func (t *TimeperiodRange) UpdateValues() []interface{} {
	v := make([]interface{}, 0)

	v = append(
		v,
		utils.Checksum(t.TimeperiodId),
		t.RangeKey,
		t.RangeValue,
		utils.Checksum(t.EnvId),
	)

	return v
}

func (t *TimeperiodRange) GetId() string {
	return t.Id
}

func (t *TimeperiodRange) SetId(id string) {
	t.Id = id
}

func (t *TimeperiodRange) GetFinalRows() ([]connection.Row, error) {
	return []connection.Row{t}, nil
}

func init() {
	name := "timeperiod_range"
	ObjectInformation = configobject.ObjectInformation{
		ObjectType: name,
		RedisKey: "timeperiod:range",
		DeltaMySqlField: "id",
		Factory: NewTimeperiodRange,
		HasChecksum: false,
		BulkInsertStmt: connection.NewBulkInsertStmt(name, Fields),
		BulkDeleteStmt: connection.NewBulkDeleteStmt(name),
		BulkUpdateStmt: connection.NewBulkUpdateStmt(name, Fields),
		NotificationListenerType: "timeperiod",
	}
}