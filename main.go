package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	exampleGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "word_cloud",
			Help: "Example Gauge",
		},
		[]string{"word"},
	)
)

var (
	words = []string{
		"二日酔い", "情報提供", "PV", "小中学校", "火縄銃", "大相撲", "バリケード", "ホラ吹き", "深淵", "カルボナーラ",
		"平成", "群青色", "高校野球", "除草剤", "大渋滞", "異口同音", "荷物持ち", "サッカー", "ストリート", "不良生徒",
	}
)

func setRandomValue() {
	for {
		exampleGauge.Reset()
		rand.Seed(time.Now().UnixNano())
		size := rand.Int() % len(words)
		for i := 0; i < size; i++ {
			w := rand.Int() % len(words)
			value := (rand.Int() % 10) + 1
			exampleGauge.With(prometheus.Labels{"word": words[w]}).Set(float64(value))
		}
		time.Sleep(1 * time.Minute)
	}
}

func main() {

	go setRandomValue()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
