// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var proto_service_pb = require('../proto/service_pb.js');

function serialize_proto_AddRequest(arg) {
  if (!(arg instanceof proto_service_pb.AddRequest)) {
    throw new Error('Expected argument of type proto.AddRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AddRequest(buffer_arg) {
  return proto_service_pb.AddRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AddResponse(arg) {
  if (!(arg instanceof proto_service_pb.AddResponse)) {
    throw new Error('Expected argument of type proto.AddResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AddResponse(buffer_arg) {
  return proto_service_pb.AddResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AggregateRequest(arg) {
  if (!(arg instanceof proto_service_pb.AggregateRequest)) {
    throw new Error('Expected argument of type proto.AggregateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AggregateRequest(buffer_arg) {
  return proto_service_pb.AggregateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_AggregateResponse(arg) {
  if (!(arg instanceof proto_service_pb.AggregateResponse)) {
    throw new Error('Expected argument of type proto.AggregateResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_AggregateResponse(buffer_arg) {
  return proto_service_pb.AggregateResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GreetRequest(arg) {
  if (!(arg instanceof proto_service_pb.GreetRequest)) {
    throw new Error('Expected argument of type proto.GreetRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GreetRequest(buffer_arg) {
  return proto_service_pb.GreetRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_GreetResponse(arg) {
  if (!(arg instanceof proto_service_pb.GreetResponse)) {
    throw new Error('Expected argument of type proto.GreetResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_GreetResponse(buffer_arg) {
  return proto_service_pb.GreetResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_PrimeRequest(arg) {
  if (!(arg instanceof proto_service_pb.PrimeRequest)) {
    throw new Error('Expected argument of type proto.PrimeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_PrimeRequest(buffer_arg) {
  return proto_service_pb.PrimeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_proto_PrimeResponse(arg) {
  if (!(arg instanceof proto_service_pb.PrimeResponse)) {
    throw new Error('Expected argument of type proto.PrimeResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_proto_PrimeResponse(buffer_arg) {
  return proto_service_pb.PrimeResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// Service Contract
var AppServiceService = exports.AppServiceService = {
  // Request Response Pattern
add: {
    path: '/proto.AppService/Add',
    requestStream: false,
    responseStream: false,
    requestType: proto_service_pb.AddRequest,
    responseType: proto_service_pb.AddResponse,
    requestSerialize: serialize_proto_AddRequest,
    requestDeserialize: deserialize_proto_AddRequest,
    responseSerialize: serialize_proto_AddResponse,
    responseDeserialize: deserialize_proto_AddResponse,
  },
  // Server streaming
generatePrimes: {
    path: '/proto.AppService/GeneratePrimes',
    requestStream: false,
    responseStream: true,
    requestType: proto_service_pb.PrimeRequest,
    responseType: proto_service_pb.PrimeResponse,
    requestSerialize: serialize_proto_PrimeRequest,
    requestDeserialize: deserialize_proto_PrimeRequest,
    responseSerialize: serialize_proto_PrimeResponse,
    responseDeserialize: deserialize_proto_PrimeResponse,
  },
  // Client Streaming
aggregate: {
    path: '/proto.AppService/Aggregate',
    requestStream: true,
    responseStream: false,
    requestType: proto_service_pb.AggregateRequest,
    responseType: proto_service_pb.AggregateResponse,
    requestSerialize: serialize_proto_AggregateRequest,
    requestDeserialize: deserialize_proto_AggregateRequest,
    responseSerialize: serialize_proto_AggregateResponse,
    responseDeserialize: deserialize_proto_AggregateResponse,
  },
  // Bidirectional streaming
greet: {
    path: '/proto.AppService/Greet',
    requestStream: true,
    responseStream: true,
    requestType: proto_service_pb.GreetRequest,
    responseType: proto_service_pb.GreetResponse,
    requestSerialize: serialize_proto_GreetRequest,
    requestDeserialize: deserialize_proto_GreetRequest,
    responseSerialize: serialize_proto_GreetResponse,
    responseDeserialize: deserialize_proto_GreetResponse,
  },
};

exports.AppServiceClient = grpc.makeGenericClientConstructor(AppServiceService, 'AppService');
// Operation Contracts
