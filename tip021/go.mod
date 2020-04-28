module github.com/dlsniper/tipsandtricks/tip021

go 1.11

// Step 1. Use the require live template
// E.g. require github.com/gin-gonic/gin v1.6.2
require github.com/gin-gonic/gin v1.6.2

// Step 2. Use the replace live template
// E.g. replace github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.6.1
replace (
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.6.1
)

// Step 3. Replace the repository with a local path
// E.g. replace github.com/gin-gonic/gin => ../mylocalgin

// Step 4. Commit from the IDE.
