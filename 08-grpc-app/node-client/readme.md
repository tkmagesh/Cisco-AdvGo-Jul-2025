## to install the grpc tool for generating proxy and stub
```shell
npm install -g grpc-tools
```

## To generate the proxy and the stub
- Run the following command in the `grpc-app` folder
```shell
grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./node-client/static_codegen/ --grpc_out=grpc_js:./node-client/static_codegen/ ./proto/service.proto
```

## To setup
```shell
npm run setup
```

## To run the client
- Make sure the grpc server is running

```shell
npm run client
```