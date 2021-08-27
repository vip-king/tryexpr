package main

import (
	"fmt"

	"bitbucket.org/ai69/amoy"
	"github.com/antonmedv/expr"
	"go.uber.org/zap"
)

type Status string

const (
	Success Status = "Success"
	Failure Status = "Failure"
)

func (s Status) String() string {
	if s == Success {
		return "✅"
	} else if s == Failure {
		return "❎"
	} else {
		return "👃🏻"
	}
}

func main() {
	log := amoy.SimpleZapSugaredLogger()

	env := map[string]interface{}{
		"okay": Success,
		"fail": Failure,
		"str":  func(v interface{}) string { return fmt.Sprint(v) },
	}
	queries := []string{
		`true`,
		`false`,
		`str(okay)`,
		`str(okay) == "Success"`,
		`okay`,
		`okay == "Success"`,
		`str(okay) == "✅"`,
	}
	for _, query := range queries {
		prog, err := expr.Compile(query, expr.Env(env))
		if err != nil {
			log.Warnw("failed to compile query", "query", query, zap.Error(err))
			continue
		}

		var output interface{}
		if output, err = expr.Run(prog, env); err != nil {
			log.Warnw("failed to execute query", "query", query, zap.Error(err))
			continue
		} else {
			log.Infow("run successfully", "query", query, "output", output)
		}
	}
}
