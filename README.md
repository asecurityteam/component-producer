<a id="markdown-component-producer---settings-component-for-publishing-events-to-streams" name="component-producer---settings-component-for-publishing-events-to-streams"></a>
# component-producer - Settings component for publishing events to streams
[![GoDoc](https://godoc.org/github.com/asecurityteam/component-producer?status.svg)](https://godoc.org/github.com/asecurityteam/component-producer)


[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=bugs)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=code_smells)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=coverage)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=duplicated_lines_density)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Lines of Code](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=ncloc)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=alert_status)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=security_rating)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=sqale_index)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=asecurityteam_component-producer&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=asecurityteam_component-producer)

<!-- TOC -->

- [component-producer - Settings component for publishing events to streams](#component-producer---settings-component-for-publishing-events-to-streams)
    - [Overview](#overview)
    - [Quick Start](#quick-start)
    - [Status](#status)
    - [Contributing](#contributing)
        - [Building And Testing](#building-and-testing)
        - [License](#license)
        - [Contributing Agreement](#contributing-agreement)

<!-- /TOC -->

<a id="markdown-overview" name="overview"></a>
## Overview

This is a [`settings`](https://github.com/asecurityteam/settings) that enables
constructing a an event producer. The resulting client may be either an HTTP
client provided by
[`component-httpclient`](https://github.com/asecurityteam/component-httpclient)
or [`benthos`](https://github.com/Jeffail/benthos). We use this, for example,
to enable migration to and from servers and serverless deployment models for our
event stream producers. When running on servers we run `benthos` as part of a
pod or docker compose and use the HTTP option to send events from our app to the
`benthos` process. When running in serverless we statically bind the `benthos`
functionality to the app by using the `benthos` option.

<a id="markdown-quick-start" name="quick-start"></a>
## Quick Start

```golang
package main

import (
    "context"

    producer "github.com/asecurityteam/component-producer"
    "github.com/asecurityteam/settings"
)

func main() {
    ctx := context.Background()
    envSource := settings.NewEnvSource(os.Environ())

    p := producer.New(ctx, envSource)
    tr := components.NewHTTP(ctx, envSource)
}
```

<a id="markdown-status" name="status"></a>
## Status

This project is in incubation which means we are not yet operating this tool in
production and the interfaces are subject to change.

<a id="markdown-contributing" name="contributing"></a>
## Contributing

<a id="markdown-building-and-testing" name="building-and-testing"></a>
### Building And Testing

We publish a docker image called [SDCLI](https://github.com/asecurityteam/sdcli) that
bundles all of our build dependencies. It is used by the included Makefile to help
make building and testing a bit easier. The following actions are available through
the Makefile:

-   make dep

    Install the project dependencies into a vendor directory

-   make lint

    Run our static analysis suite

-   make test

    Run unit tests and generate a coverage artifact

-   make integration

    Run integration tests and generate a coverage artifact

-   make coverage

    Report the combined coverage for unit and integration tests

<a id="markdown-license" name="license"></a>
### License

This project is licensed under Apache 2.0. See LICENSE.txt for details.

<a id="markdown-contributing-agreement" name="contributing-agreement"></a>
### Contributing Agreement

Atlassian requires signing a contributor's agreement before we can accept a patch. If
you are an individual you can fill out the [individual
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=3f94fbdc-2fbe-46ac-b14c-5d152700ae5d).
If you are contributing on behalf of your company then please fill out the [corporate
CLA](https://na2.docusign.net/Member/PowerFormSigning.aspx?PowerFormId=e1c17c66-ca4d-4aab-a953-2c231af4a20b).
