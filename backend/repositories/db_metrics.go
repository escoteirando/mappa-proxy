package repositories

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"gorm.io/gorm"
)

var tables = []string{"associados", "escotistas", "grupos", "secoes"}
var gauges = make(map[string]prometheus.GaugeFunc, len(tables))

func getCountGaugeFunc(db *gorm.DB, table string) func() float64 {
	return func() float64 {
		var count int64
		db.Table(table).Count(&count)
		return float64(count)
	}
}
func SetupDatabaseMetrics(db IRepository) {
	for _, table := range tables {
		var c = promauto.NewGaugeFunc(prometheus.GaugeOpts{
			Name:        "mappa_entities",
			ConstLabels: prometheus.Labels{"table": table},
			Help:        "Table count",
		}, getCountGaugeFunc(db.GetDBFunc(), table))
		gauges[table] = c
	}
}
