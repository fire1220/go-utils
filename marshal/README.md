# marshal

---

### describe
> 使用json.Marshal把结构体转换成json时  
> 会自动将time.Time类型解析成"2024-03-14 18:10:09"格式  
> tag中datetime表示自定义结构，后面添加,omitempty表示如果空值，格式化为空字符串   


## Usage

```
import "github.com/fire1220/goutils/marshal"
```

### example:
``` go
package test

import (
	"encoding/json"
	"fmt"
	"github.com/fire1220/goutils/marshal"
	"testing"
	"time"
)

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	PlayTime  time.Time `json:"play_time"`
	CreatedAt time.Time `json:"created_at" datetime:"2006-01-02,omitempty"`
	UpdatedAt time.Time `json:"updated_at" datetime:"2006-01-02 15:04:05,omitempty"`
}

func (g Good) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(g)
}

func TestMarshal(t *testing.T) {
	d, _ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
	good := Good{ID: 123, Name: "jock", PlayTime: time.Now(), CreatedAt: time.Now(), UpdatedAt: d}
	bytes, _ := json.Marshal(good)
	// {"id":123,"name":"jock","play_time":"2024-03-15 10:28:38","created_at":"2024-03-15","updated_at":""}
	fmt.Printf("%s\n", bytes)
}
```