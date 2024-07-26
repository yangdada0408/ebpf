package examples

// DocRlimit {
import "github.com/yangdada0408/ebpf/rlimit"

func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		panic(err)
	}
}

// }
