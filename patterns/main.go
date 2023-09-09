package main

import "fmt"

func main() {

	// BUILDER PATTERN
	var builderConfig = NewConfigBuilder().
		SetUserId(1).
		SetDomain("localhost").
		SetPort(443).
		SetKey("supersecretkey").
		SetUrl("go/lang/is/awesome").
		Build()

	fmt.Println(builderConfig)

	// FACTORY PATTERN
	var factoryDevConfig = NewFactoryConfig("dev", 1, "http://localhost:3001", "supersecretdevelopmentkey") 
	var factoryProdConfig = NewFactoryConfig("prod", 2, "https://jd-2023.vercel.app", "supersecretproductionkey")

	fmt.Println("Factory Dev Config: ", factoryDevConfig);
	fmt.Println("Factory Prod Config: ", factoryProdConfig);

	// SINGLETON PATTERN
}
