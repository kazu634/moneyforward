/*
Copyright © 2024 Kazuhiro MUSASHI <simoom634@yahoo.co.jp>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gitea.kazu634.com/kazu634/moneyforward/internal/lib/browser"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

const (
	colBank int = iota
	colStock
	colFund
	colReceivables
	colMiles
)

const AmountCol = 1

func convInt(target string) int {
	tmp := strings.ReplaceAll(target, ",", "")
	tmp = strings.ReplaceAll(tmp, "円", "")

	result, _ := strconv.Atoi(tmp)

	return result
}

// portfolioCmd represents the portfolio command
var portfolioCmd = &cobra.Command{
	Use:   "portfolio",
	Short: "A subcommand to fetch the portfolio data from Moneyforward",
	Long:  `A subcommand to fetch the portfolio data from Moneyforward`,
	Run: func(cmd *cobra.Command, args []string) {
		logInfo("Launching a browser..")
		browser := browser.Launch()
		defer browser.Close()

		browser = login(browser)

		logInfo("Navigating to the portfolio page...")
		browser.Click("a[href^='/bs/portfolio']")

		logInfo("Fetching the portfolio data...")
		reader := strings.NewReader(browser.GetElmHTML("#bs-portfolio > div:nth-child(1) > div > div.mf-col-custom-content-body > section > div.row > div > section > table"))

		doc, _ := goquery.NewDocumentFromReader(reader)

		logInfo("Fetching each portfolio item amount...")
		var result []int
		doc.Find("table").Each(func(i int, tablehtml *goquery.Selection) {
			tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
				rowhtml.Find("th, td").Each(func(indexth int, tablecell *goquery.Selection) {
					if indexth == AmountCol {
						result = append(result, convInt(strings.TrimSpace(tablecell.Text())))
					}
				})
			})
		})

		logInfo(fmt.Sprintf("Get this data: %v", result))

		logInfo("Outputting data...")
		dateStr := time.Now().Format("2006-01-02")
		fmt.Printf("%s\t%d\t%d\t%d\t%d\t%d\n", dateStr, result[colBank], result[colStock], result[colFund], result[colMiles], result[colReceivables])

		logInfo("Finished.")
	},
}

func init() {
	rootCmd.AddCommand(portfolioCmd)

	filter := logInit()
	log.SetOutput(filter)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portfolioCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portfolioCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
