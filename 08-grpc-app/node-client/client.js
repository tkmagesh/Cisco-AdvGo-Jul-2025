/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

var parseArgs = require("minimist");
var messages = require("./static_codegen/proto/service_pb");
var services = require("./static_codegen/proto/service_grpc_pb");

var grpc = require("@grpc/grpc-js");

function main() {
  const target = "localhost:50052";
  
  var client = new services.AppServiceClient(
    target,
    grpc.credentials.createInsecure()
  );
  var request = new messages.AddRequest();
  request.setX(100)
  request.setY(200)
  
  client.add(request, function (err, response) {
    if (err !== null) {
        console.log("err :", err)
        return
    }
    console.log("Add Result:", response.getResult());
  });
}

main();
