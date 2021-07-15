package task

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
)

type taskFunc struct {

}

func init() {
	timeTask := toolbox.NewTask("TimeTask", "0 */5 * * * *",  func () error { return TimeTask() } )
	err := timeTask.Run()
	if err != nil {
		fmt.Println("TimeTask err:", err)
	}

	toolbox.AddTask("TimeTask", timeTask)
	//toolbox.StartTask()
}

func Stop()  {
	toolbox.StopTask()
}

func registerTask()  {

}

func TimeTask() error {
	fmt.Println("this is TimeTask")
	return nil
}