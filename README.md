# Let's Go! React with Go Complete Fullstack App - TypeScript, React Query, MongoDB, ChakraUI

![Demo App](https://i.ibb.co/JvRTWmW/Group-93.png)

[Video Tutorial on Youtube](https://youtu.be/zw8z_o_kDqc)

Some Features:

-   ⚙️ Tech Stack: Go, React, TypeScript, MongoDB, TanStack Query, ChakraUI
-   ✅ Create, Read, Update, and Delete (CRUD) functionality for todos
-   🌓 Light and Dark mode for user interface
-   📱 Responsive design for various screen sizes
-   🌐 Deployment
-   🔄 Real-time data fetching, caching, and updates with TanStack Query
-   🎨 Stylish UI components with ChakraUI
-   ⏳ And much more!

### .env file

```shell
MONGO_URI=<your_mongo_uri>
PORT=5000
ENV=development
```

### Compile and run

```shell
go run main.go
```

Packages installed:
github.com/aws/aws-lambda-go/lambda
github.com/aws/aws-lambda-go/events
github.com/aws/aws-sdk-go-v2/config
github.com/aws/aws-sdk-go-v2/service/dynamodb
github.com/aws/aws-sdk-go-v2/service/dynamodb/types
github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
github.com/google/uuid

to build lamdazip:
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o bootstrap
Compress-Archive bootstrap function.zip -Force

for client:
npm install
npm run build