# Multi container components

In this example we have a component requiring two docker images, trying to reproduce the typical pattern where a component has different processes.
We'll use the typical php-fpm + nginx example.

* _Dockerfile_ is our php application, it depends on the src folder only
* _Dockerfile.nginx_ is our nginx component, it depends on the nginx subfolder
