package leadlogger

import "testing"
import "time"
import "fmt"

func TestDefaultFormat(t *testing.T) {
	//We just output this to see, if it's correct
	fmt.Println(DEFAULT_FORMAT(LEVEL_ALERT, "TAG", "message"))
}

func TestCycle(t *testing.T) {
	Start()
	//Waiting for the routine to start
	time.Sleep(time.Millisecond)
	if !Running() {
		t.Error("Expected the routine to run after a call to Start, but it doesn't")
	}
	Stop()
	//Waiting for the routine to react
	time.Sleep(time.Millisecond)
	if Running() {
		t.Error("Expected the routine to be stoppeda fter a call to Stop, but it doesn't")
	}

}

func TestIdeal(t *testing.T) {
	Start(TargetStdOut)

	Emergency("TAG", "Emergency")
	Alert("TAG", "Alert")
	Critical("TAG", "Critical")
	Error("TAG", "Error")
	Warning("TAG", "Warning")
	Notice("TAG", "Notice")
	Informational("TAG", "Informational")
	Debug("TAG", "Debug")
	Stop()
}

func TestFile(t *testing.T) {
	Start(&FileTarget{
		Path:       "log.log",
		MaxLevel:   LEVEL_NOTICE,
		BufferSize: 0,
	})

	defer Stop()

	for i := 0; i < 40; i++ {
		Emergency("TAG", i)
		Alert("TAG", i)
		Critical("TAG", i)
		Error("TAG", i)
		Warning("TAG", i)
		Notice("TAG", i)
		Informational("TAG", i)
		Debug("TAG", i)
	}
}

func TestTimeout(t *testing.T) {
	Start(&StdOutTarget{
		MaxLevel:   LEVEL_DEBUG,
		BufferSize: 10,
	})
	defer Stop()

	Alert("TAG", "One")
	time.Sleep(10 * time.Second)
	Alert("TAG", "Two")

}
