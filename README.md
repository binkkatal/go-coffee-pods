# GOLANG CHALLENGE

This repo demonstrates the listing of coffee machines and pods by a coffee company using a go backend server using the following end points

## Endpoints

- `/api/product/listOfCoffeeMachines` -> this endpoint accepts a `size_id` query_parameter, if left blank all coffee machines will be returned. 
- `/api/product/listOfCoffeePods` -> this endpoint accepts a `flavor_id` and/or a `size_id` query parameter. if left blank will return all pods.
- `/api/crosssell/coffeeMachines` -> this endpoint requires a `coffee_machine_id` query parameter. Will return the smallest packs of pods, one per flavor.

- `/api/crosssell/coffeePods` -> this endpoint requires a `pod_id` will return all other pods with a matching `flavor_id` and `size_id` (excluding itself).

## How to build the app

Go into the repo folder build the project using

	go build .

This will create an executable file `GoLangTask` which can be run directly

	./GoLangTask

Or you can simply use

	go run main.go

Which will start listening the server at `localhost:8080`

## Sample Requests

### All large machines

	http://localhost:8080/api/product/listOfCoffeeMachines?size_id=3

```json
[{\"coffee_machine_id\":4,\"size_id\":3,\"size_name\":\"large\",\"sku\":\"CM101\",\"model_id\":1,\"model_name\":\"base model\"},{\"coffee_machine_id\":5,\"size_id\":3,\"size_name\":\"large\",\"sku\":\"CM102\",\"model_id\":2,\"model_name\":\"premium model\",\"water_line\":true},{\"coffee_machine_id\":6,\"size_id\":3,\"size_name\":\"large\",\"sku\":\"CM103\",\"model_id\":3,\"model_name\":\"deluxe model\",\"water_line\":true}]
```

### All cross-sell for large machine, smallest per flavor

	http://localhost:8080/api/crosssell/coffeeMachines?coffee_machine_id=5

```json
[{\"pod_id\":11,\"size_id\":3,\"size_name\":\"large\",\"flavor_id\":1,\"flavor_name\":\"vanilla\",\"sku\":\"CP101\",\"quantity\":12},{\"pod_id\":13,\"size_id\":3,\"size_name\":\"large\",\"flavor_id\":2,\"flavor_name\":\"caramel\",\"sku\":\"CP111\",\"quantity\":12},{\"pod_id\":15,\"size_id\":3,\"size_name\":\"large\",\"flavor_id\":3,\"flavor_name\":\"psl\",\"sku\":\"CP121\",\"quantity\":12},{\"pod_id\":17,\"size_id\":3,\"size_name\":\"large\",\"flavor_id\":4,\"flavor_name\":\"mocha\",\"sku\":\"CP131\",\"quantity\":12},{\"pod_id\":19,\"size_id\":3,\"size_name\":\"large\",\"flavor_id\":5,\"flavor_name\":\"hazelnut\",\"sku\":\"CP141\",\"quantity\":12}]
```

### All choices on espresso vanilla landing page

	http://localhost:8080/api/product/listOfCoffeePods?flavour_id=1&size_id=1

```json
[{\"pod_id\":21,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":1,\"flavor_name\":\"vanilla\",\"sku\":\"EP003\",\"quantity\":36},{\"pod_id\":22,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":1,\"flavor_name\":\"vanilla\",\"sku\":\"EP005\",\"quantity\":60},{\"pod_id\":23,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":1,\"flavor_name\":\"vanilla\",\"sku\":\"EP007\",\"quantity\":84},{\"pod_id\":24,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":2,\"flavor_name\":\"caramel\",\"sku\":\"EP013\",\"quantity\":36},{\"pod_id\":25,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":2,\"flavor_name\":\"caramel\",\"sku\":\"EP015\",\"quantity\":60},{\"pod_id\":26,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":2,\"flavor_name\":\"caramel\",\"sku\":\"EP017\",\"quantity\":84}]
```

### All Espresso Machines

	http://localhost:8080/api/product/listOfCoffeeMachines?size_id=1

```json
[{\"coffee_machine_id\":7,\"size_id\":1,\"size_name\":\"espresso\",\"sku\":\"EM001\",\"model_id\":1,\"model_name\":\"base model\"},{\"coffee_machine_id\":8,\"size_id\":1,\"size_name\":\"espresso\",\"sku\":\"EM002\",\"model_id\":2,\"model_name\":\"premium model\"},{\"coffee_machine_id\":9,\"size_id\":1,\"size_name\":\"espresso\",\"sku\":\"EM003\",\"model_id\":3,\"model_name\":\"deluxe model\",\"water_line\":true}]
```

### All cross-sell for Espresso machine, smallest per flavor

	http://localhost:8080/api/crosssell/coffeeMachines?coffee_machine_id=9

```json
[{\"pod_id\":21,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":1,\"flavor_name\":\"vanilla\",\"sku\":\"EP003\",\"quantity\":36},{\"pod_id\":24,\"size_id\":1,\"size_name\":\"espresso\",\"flavor_id\":2,\"flavor_name\":\"caramel\",\"sku\":\"EP013\",\"quantity\":36}]
```