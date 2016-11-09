# Page Elements

Proof of concept project that explores a possibility of splitting a big web application / site into manageable peaces that work nicely together using the following stack: docker, consul, traefik and grpc as a transport mechanism.

## What it does

Each page is composed of multiple elements that can be developed and deployed separately. Component `site` pulls all of them together and composes the resulting source that is sent to the browser. As elements are also versioned you can have multiple versions of each element running on your servers and the operator is able to switch to a different version in a simple dashboard.

### Elements

Element is a part of a website that has a clear separate responsibility. It can include other elements which will be rendered by `site` concurrently.

Each element needs to be able to respond to two rpc requests:

1. **describe**

     Describes if any other elements need to be included and a possibility to modify the arguments sent to all elements in the render phase.

2. **render**

    Generates the HTML and returns it together additional instructions (meta tags, title, etc.) back to `site`.
 
### Site
 
Component site manages routes that link to specific elements and contains logic how to communicate with elements and compose their response into a website.

It expects an element `skeleton` to be available and enabled in the system that it will use to wrap the requested element with (to set common layout, load static assets, etc.). It is also possible to skip this by loading the page with `?format=snippet` query parameter.

#### Communication with elements

Page is rendered following steps:

1. Describe

    Description of all the requested elements is collected to get the full tree of elements required to render the page.

2. Render

    Once the tree is composed rendering is triggered on all the elements concurrently and if needed certain elements are postponed until their children are completed (ex: skeleton needs all elements to complete to get all the required parameters back).

3. Compose

    Once all elements are done rendering they are glued together by replacing a special `<element>{elementName}</element>` tag in the rendered responses.

#### Routing

Two different route types are available:
* /pel/:element[/:version][?format=snippet]
    Allows the developer to preview or target a specific element. It is also possible to load a specific version of an element.

* /custom/routes
    These routes should point to a specific page element.

### Dashboard

Dashboard is a simple control panel for setting the current active version on each element. It will pull the current state of consul catalog to present possible options and save the preferred state to consul's key/value storage that `site` monitors and responds to if altered.

### Advantages

* Elements can be developed separately.
* Impact of a deployment is reduced to a specific functionality.
* New features can be easily tested before they are exposed to the public.

## How to use it

### Requirements

* docker
* docker-compose
* echo "127.0.0.1 micro.site dash.micro.site" >> /etc/hosts

```bash
# start everything up
$ make up

# open the dashboard (don't forget to update /etc/hosts) and enable all elements
$ open http://dash.micro.site

# page should work now
$ open http://micro.site

# to reload a single element or service;
$ make up/header

# stop everything
$ make stop

# remove everything
$ make clean
```


