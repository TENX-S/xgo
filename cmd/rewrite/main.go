package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `automake autogen build-essential ca-certificates                    
    gcc-5-arm-linux-gnueabi g++-5-arm-linux-gnueabi libc6-dev-armel-cross                
    gcc-5-arm-linux-gnueabihf g++-5-arm-linux-gnueabihf libc6-dev-armhf-cross            
    gcc-5-aarch64-linux-gnu g++-5-aarch64-linux-gnu libc6-dev-arm64-cross                
    gcc-5-mips-linux-gnu g++-5-mips-linux-gnu libc6-dev-mips-cross                       
    gcc-5-mipsel-linux-gnu g++-5-mipsel-linux-gnu libc6-dev-mipsel-cross                 
    gcc-5-mips64-linux-gnuabi64 g++-5-mips64-linux-gnuabi64 libc6-dev-mips64-cross       
    gcc-5-mips64el-linux-gnuabi64 g++-5-mips64el-linux-gnuabi64 libc6-dev-mips64el-cross 
    gcc-5-multilib g++-5-multilib gcc-mingw-w64 g++-mingw-w64 clang llvm-dev             
    gcc-6-arm-linux-gnueabi g++-6-arm-linux-gnueabi libc6-dev-armel-cross                
    gcc-6-arm-linux-gnueabihf g++-6-arm-linux-gnueabihf libc6-dev-armhf-cross            
    gcc-6-aarch64-linux-gnu g++-6-aarch64-linux-gnu libc6-dev-arm64-cross                
    gcc-6-mips-linux-gnu g++-6-mips-linux-gnu libc6-dev-mips-cross                       
    gcc-6-mipsel-linux-gnu g++-6-mipsel-linux-gnu libc6-dev-mipsel-cross                 
    gcc-6-mips64-linux-gnuabi64 g++-6-mips64-linux-gnuabi64 libc6-dev-mips64-cross       
    gcc-6-mips64el-linux-gnuabi64 g++-6-mips64el-linux-gnuabi64 libc6-dev-mips64el-cross 
    gcc-6-s390x-linux-gnu g++-6-s390x-linux-gnu libc6-dev-s390x-cross 
    gcc-6-powerpc64le-linux-gnu g++-6-powerpc64le-linux-gnu libc6-dev-powerpc-ppc64-cross 
    gcc-8-riscv64-linux-gnu g++-8-riscv64-linux-gnu libc6-dev-riscv64-cross 
    gcc-6-multilib g++-6-multilib gcc-7-multilib g++-7-multilib gcc-mingw-w64 g++-mingw-w64 
    clang llvm-dev libtool libxml2-dev uuid-dev libssl-dev swig openjdk-8-jdk pkg-config patch 
    make xz-utils cpio wget zip unzip p7zip git mercurial bzr texinfo help2man cmake curl mercurial`

	prefix := "RUN apt-get install -y"

	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	var packages []string
	for _, line := range lines {
		lPkgs := strings.Split(line, " ")
		for i, lPkg := range lPkgs {
			lPkgs[i] = strings.TrimSpace(lPkg)
		}
		packages = append(packages, lPkgs...)
	}

	var (
		commands []string
		i        int
	)
	for i = 0; i+1 < len(packages); i += 2 {
		commands = append(commands, strings.Join([]string{prefix, packages[i], packages[i+1]}, " "))
	}

	if i < len(packages)-1 {
		commands = append(commands, strings.Join([]string{prefix, packages[len(packages)-1]}, " "))
	}
	fmt.Println(strings.Join(commands, "\n"))
}
