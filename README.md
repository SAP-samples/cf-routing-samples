# Cloud Foundry Routing Samples

[![REUSE status](https://api.reuse.software/badge/github.com/SAP-samples/cf-routing-samples)](https://api.reuse.software/info/github.com/SAP-samples/cf-routing-samples)

## Description

The SAP Business Technology Platform (BTP) provides a runtime environment for running your applications at scale.

This repository contains samples for applications running on the Cloud Foundry runtime for SAP BTP, related to routing topics.

In order to leverage more advance features of the routing stack of the Cloud Foundry runtime for SAP BTP, such as full support for HTTP/2, your apps need to add support.

The following samples are available in the respective subdirectories:

* [`http2`](http2/README.md) - HTTP/2 and gRPC enabled server samples for Cloud Foundry

## Requirements

The samples provided in this repository are intended to run in a Cloud Foundry environment. Each of the samples may contain further information on specific technical requirements.

Samples are provided in different programming languages and require the respective runtime environment for development.
## Download and Installation

These samples are provided as source code and should be seen as starting point or reference on how to approach or solve a particular requirement.

The download from GitHub, either via Git or archive, is recommended.

## Known Issues

- HTTP/2 support is being rolled out globally but is not complete. In order to use the HTTP/2 samples, you need to enable a custom domain with HTTP/2 support. 
  <!-- todo link to the blog post about that! -->

## How to obtain support
[Create an issue](https://github.com/SAP-samples/cf-routing-samples/issues) in this repository if you find a bug or have questions about the content.
 
For additional support, [ask a question in SAP Community](https://answers.sap.com/questions/ask.html).

## Contributing
If you wish to contribute code, offer fixes or improvements, please send a pull request. Due to legal reasons, contributors will be asked to accept a DCO when they create the first pull request to this project. This happens in an automated fashion during the submission process. SAP uses [the standard DCO text of the Linux Foundation](https://developercertificate.org/).

## License
Copyright (c) 2022 SAP SE or an SAP affiliate company. All rights reserved. This project is licensed under the Apache Software License, version 2.0 except as noted otherwise in the [LICENSE](LICENSES/Apache-2.0.txt) file.
