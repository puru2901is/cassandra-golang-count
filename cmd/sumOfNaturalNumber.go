/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var number int

// sumOfNaturalNumberCmd represents the sumOfNaturalNumber command
var sumOfNaturalNumberCmd = &cobra.Command{
	Use:   "sumOfNaturalNumber",
	Short: "Sum of n natural number",
	Long: `This is for the sum of n natural number. you can provide n with parameter number. By default the value is 100. For example:
	cassandra-golang-count sumOfNaturalNumber --number 10
	It will give you 55 as output.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var sum int
		fmt.Println(number)
		bar := progressbar.Default(int64(number))
		for i := 1; i <= number; i++ {
			bar.Add(1)
			time.Sleep(40 * time.Millisecond)
			sum += i
		}
		color.Magenta("sum of natural numbers: %v", sum)
		time.Sleep(200000 * time.Millisecond)
	},
}

func init() {
	rootCmd.AddCommand(sumOfNaturalNumberCmd)

	sumOfNaturalNumberCmd.PersistentFlags().IntVar(&number, "number", 100, "sum of n natural numbers")

}
