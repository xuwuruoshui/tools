package main

var (
	EnumApiInputPath = "./template/code/api_resp.txt"
	EnumApiOutPath = "../end/api/handler.txt/code.go"

	EnumServiceInputPath = "./template/code/service_resp.txt"
	EnumServiceOutputPath = "../end/service/code.go"

	EnumRepositoryInputPath = "./template/code/repo_resp.txt"
	EnumRepositoryOutputPath = "../end/repository/code.go"

)

func main(){

	// enum 转 map
	ConstGenerateMap(EnumApiInputPath,EnumApiOutPath)
	ConstGenerateMap(EnumServiceInputPath,EnumServiceOutputPath)
	ConstGenerateMap(EnumRepositoryInputPath,EnumRepositoryOutputPath)

}




