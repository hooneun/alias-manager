package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "am",
	Short: "셸 히스토리 분석을 통한 스마트 별칭 관리 도구",
	Long: `Alias Manager (am)은 셸 히스토리를 분석하여 자주 사용하는 명령어에 대한
유용한 별칭(alias)을 자동으로 제안하는 스마트 명령행 도구입니다.

지능적으로 명령어 빈도와 복잡도를 분석하여 더 짧고 기억하기 쉬운 별칭을
제안하며, 기존 별칭과의 충돌을 방지하고 안전하게 셸 설정 파일에 적용할 수 있습니다.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🔍 Alias Manager - 스마트 별칭 관리 도구")
		fmt.Println("사용 가능한 명령어를 보려면 'am --help'를 실행하세요.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
