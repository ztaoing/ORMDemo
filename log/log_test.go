/**
* @Author:zhoutao
* @Date:2021/1/14 下午12:59
* @Desc:
 */

package log

import (
	"os"
	"testing"
)

func TestSetLevel(t *testing.T) {
	SetLevel(ErrorLevel)
	if infoLog.Writer() == os.Stdout || errLog.Writer() != os.Stdout {
		t.Fatalf("failed to set log level: %v", infoLog.Writer())
	}
	SetLevel(Disabled)
	if infoLog.Writer() == os.Stdout || errLog.Writer() == os.Stdout {
		t.Fatal("failed to set log level")
	}
}
