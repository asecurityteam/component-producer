<a id="markdown-component-producer---settings-component-for-publishing-events-to-streams" name="component-producer---settings-component-for-publishing-events-to-streams"></a>
# component-producer - Settings component for publishing events to streams
[![GoDoc](https://godoc.org/github.com/asecurityteam/component-producer?status.svg)](https://godoc.org/github.com/asecurityteam/component-producer)
[![Build Status](https://travis-ci.com/asecurityteam/component-producer.png?branch=master)](https://travis-ci.com/asecurityteam/component-producer)
[![codecov.io](https://codecov.io/github/asecurityteam/component-producer/coverage.svg?branch=master)](https://codecov.io/github/asecurityteam/component-producer?branch=master)
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
constructing an event producer. Currently, the only supported client is HTTP, as provided by
 [`component-httpclient`](https://github.com/asecurityteam/component-httpclient).

**[`Benthos`] (https://github.com/Jeffail/benthos) producer is no longer supported as of version 2 of the component. **

 We use this, for example,
to enable migration to and from servers and serverless deployment models for our
event stream producers. When running on servers we run `benthos` as part of a
pod or docker compose and use the HTTP option to send events from our app to the
`benthos` process. 

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
