# gcp-regions

Library to validate GCP regions.

This project is not officially supported or endorsed by Google in any way.

# Example

```go
import(
	"fmt"
  
	regions "github.com/ohkinozomu/gcp-regions"
)

func main() {
	// true
	fmt.Println(regions.IsValid("us-central1"))

	// false
	fmt.Println(regions.IsValid("us-central1000"))
}
```

# License

Apache-2.0

`data.go` is generated from [GoogleCloudPlatform/region-picker](https://github.com/GoogleCloudPlatform/region-picker/blob/main/data/regions.json).