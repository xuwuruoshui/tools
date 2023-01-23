package main

var (
	EnumApiInputPath = "./template/end/code/api_resp.txt"
	EnumApiOutPath   = "../end/api/handler/code.go"

	EnumServiceInputPath  = "./template/end/code/service_resp.txt"
	EnumServiceOutputPath = "../end/service/code.go"

	EnumRepositoryInputPath  = "./template/end/code/repo_resp.txt"
	EnumRepositoryOutputPath = "../end/repository/code.go"
)

func main() {

	// enum 转 map
	ConstGenerateMap(EnumApiInputPath, EnumApiOutPath)
	ConstGenerateMap(EnumServiceInputPath, EnumServiceOutputPath)
	ConstGenerateMap(EnumRepositoryInputPath, EnumRepositoryOutputPath)

}
