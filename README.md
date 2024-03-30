# description

The source code comes from a third-party project and is independently used to parse the app package information. Competent people are welcome to submit code.

# Obtain APK information

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

# reference

- https://github.com/iineva/ipa-server/
- https://github.com/pkg6/go-lzfse
- https://github.com/pkg6/go-plist

# 

