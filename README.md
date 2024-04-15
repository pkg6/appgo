# description

The source code comes from a third-party project and is independently used to parse the app package information. Competent people are welcome to submit code.

# According to file suffix

> AppParsePath
>
> AppParseFile
>
> AppParseReader
>
> AppParseMultipartFile

~~~
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg6/appgo"
)

func main() {
	filename := "./.test_data/ipa.ipa"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	info, _ := appgo.AppParseFile(f)
	// ico 保存在本地
	filename, buf, _ := info.Icon()
	file, _ := os.Create(filename)
	defer file.Close()
	io.Copy(file, buf)
}
~~~

# Obtain APK information

> APKParsePath
>
> APKParseFile
>
> APKParseReader
>
> APKParseMultipartFile

~~~
package main

import (
	"fmt"

	"github.com/pkg6/appgo"
)

func main() {
	info, err := appgo.APKParseFile(".test_data/helloworld.apk")
	fmt.Println(info, err)
}

~~~

# Obtain IPA information

> IPAParsePath
>
> IPAParseFile
>
> IPAParseReader
>
> IPAParseMultipartFile

~~~
package main

import (
	"fmt"

	"github.com/pkg6/appgo"
)

func main() {
	info, err := appgo.IPAParsePath(".test_data/ipa.ipa")
	fmt.Println(info, err)
}
~~~

# Link

- https://github.com/iineva/ipa-server/
- https://github.com/pkg6/go-lzfse
- https://github.com/pkg6/go-plist

# 

