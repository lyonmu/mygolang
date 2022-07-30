package main

import (
	"gobasic/compositedata"
	// "gobasic/flowcontrol"
	// "log"
	// "gobasic/variables"
	// "go.uber.org/zap"
)

func main() {
	// variables.Init()
	// flowcontrol.Init()
	compositedata.Init()

	/* logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	logger.Info("hello!", zap.String("name", "xiaomin"), zap.Int("age", 20)) */
}
