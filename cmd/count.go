/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

/*
to do:
1. add progress bar
2. colors
*/

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gocql/gocql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var table string
var keyspace string

// countCmd represents the count command
var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Count in cassanra",
	Long: `This can be used for cassandra count. You can provide table and keyspace flags. For example:
	cassandra-golang-count count --keyspace example --table table1.
	It will return count of table1 in example keypsace.
`,

	Run: func(cmd *cobra.Command, args []string) {
		hosts := viper.GetStringSlice("cassandra.hosts")
		cluster := gocql.NewCluster(hosts...)
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum

		// connect to the cluster
		session, err := cluster.CreateSession()
		if err != nil {
			color.Red("fatal error %s", err)
		}
		defer session.Close()

		color.Green("connected to cluster")

		ctx := context.Background()

		var count int = 0
		query := fmt.Sprintf("SELECT * FROM %s", table)

		scanner := session.Query(query).PageSize(1000).WithContext(ctx).Iter().Scanner()
		for scanner.Next() {
			count += 1
		}

		color.Magenta("count from cassandra table is : %v", count)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		fmt.Println("Please provide the config path in env variable 'CONFIG_PATH'. Trying defaults..")
		configPath = "conf"
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigName("count") // name of config file (without extension)
	viper.SetConfigType("json")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	countCmd.PersistentFlags().StringVar(&keyspace, "keyspace", "feeds", "keyspace in cassandra")
	countCmd.PersistentFlags().StringVar(&table, "table", "track_account_feeds", "table in cassandra")

	viper.BindPFlag("table", countCmd.PersistentFlags().Lookup("table"))
}
