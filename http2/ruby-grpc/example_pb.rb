# frozen_string_literal: true
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: example.proto

require 'google/protobuf'


descriptor_data = "\n\rexample.proto\"\t\n\x07Request\"\x1b\n\x08Response\x12\x0f\n\x07message\x18\x01 \x01(\t2\'\n\x07\x45xample\x12\x1c\n\x03Run\x12\x08.Request\x1a\t.Response\"\x00\x62\x06proto3"

pool = Google::Protobuf::DescriptorPool.generated_pool
pool.add_serialized_file(descriptor_data)

Request = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Request").msgclass
Response = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("Response").msgclass
