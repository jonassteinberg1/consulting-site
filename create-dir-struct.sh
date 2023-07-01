#!/bin/bash

mkdir -p frontend/mobile-app/src
mkdir -p frontend/mobile-app/tests
touch frontend/mobile-app/package.json
touch frontend/mobile-app/tsconfig.json

mkdir -p frontend/web-app/src
mkdir -p frontend/web-app/tests
touch frontend/web-app/package.json
touch frontend/web-app/tsconfig.json

mkdir -p backend/user-service/src
mkdir -p backend/user-service/tests
touch backend/user-service/package.json
touch backend/user-service/tsconfig.json

mkdir -p backend/shopping-cart-service/src
mkdir -p backend/shopping-cart-service/tests
touch backend/shopping-cart-service/go.mod
touch backend/shopping-cart-service/go.sum

mkdir -p backend/checkout-service/src
mkdir -p backend/checkout-service/tests
touch backend/checkout-service/go.mod
touch backend/checkout-service/go.sum

mkdir -p backend/consultation-service/src
mkdir -p backend/consultation-service/tests
touch backend/consultation-service/package.json
touch backend/consultation-service/tsconfig.json

mkdir -p infrastructure/kubernetes/user-service
mkdir -p infrastructure/kubernetes/shopping-cart-service
mkdir -p infrastructure/kubernetes/checkout-service
mkdir -p infrastructure/kubernetes/consultation-service

mkdir -p infrastructure/istio
mkdir -p infrastructure/argocd

