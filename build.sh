#!/bin/bash
echo "hello world1"
npm install --golabal yarn
mv resource_types website_code/resource_types
cd website_code/warehouse/web

yarn install && yarn build
