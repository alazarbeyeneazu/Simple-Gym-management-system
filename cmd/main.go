package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	handler "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/handlers/rest/gym"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/admin"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/authz"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/checkin"
	gymgoers "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/gym_goers"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/pyment"
	repor "github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/report"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/modules/user"
	"github.com/alazarbeyeneazu/Simple-Gym-management-system/internals/storage/persistant"
	routers "github.com/alazarbeyeneazu/Simple-Gym-management-system/platforms/routes"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xuri/excelize/v2"
)

func main() {

	dbs := persistant.Init()
	auth := authz.InitService(dbs)
	// auth.InitatePermission()
	service := user.InitService(dbs)
	pymentService := pyment.InitService(dbs)
	gymgoersService := gymgoers.InitService(dbs)
	checkinuser := checkin.InitService(dbs)
	admin := admin.InitService(dbs)
	report := repor.InitService(dbs)
	go BotHandler(report)
	// admin.InitializeSuperAdmin(context.Background(), "0948398647", "passme@123")
	user := handler.Init(service, pymentService, gymgoersService, admin, auth, checkinuser, report)
	routes := user.StartRoutes()
	router := routers.Initialize(":8282", routes)
	router.Serve()
}
func BotHandler(report repor.ReportService) {
	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Println(err)
		return
	}
	// 0911749121
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {

		if update.Message == nil {
			continue
		}
		if !update.Message.IsCommand() {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		f := excelize.NewFile()

		sreamwriter, err := f.NewStreamWriter("Sheet1")
		sreamwriter.SetColWidth(1, 30, 24)
		if err != nil {
			log.Println(err)
			continue
		}
		switch update.Message.Command() {
		case "help":
			msg.Text = `
      Available Commands 
/today  To Get Today's Report
/week   To Get this week's Report
/month  To Get This month's Report
/year  To Get This Year's Report		
/report  To Get all the reports
/backup  To Backup The database 				`
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg.Text))
		case "start":
			msg.Text = `
      Available Commands
today  To Get Today's Report
/week   To Get this week's Report
/month  To Get This month's Report
/year  To Get This Year's Report			
/report  To Get all the reports
/backup  To Backup The database  				`
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg.Text))
		case "report":

			sreamwriter.SetRow("A1", []interface{}{"First Name"})
			sreamwriter.SetRow("B1", []interface{}{"Last Name"})
			sreamwriter.SetRow("C1", []interface{}{"Registered By"})
			sreamwriter.SetRow("D1", []interface{}{"Created At"})
			sreamwriter.SetRow("E1", []interface{}{"Paid By"})
			sreamwriter.SetRow("F1", []interface{}{"Amount"})

			response, err := report.GetAllReports(context.Background())
			if err != nil {
				log.Panicln(err)
				continue
			}
			b := 2
			for i, rep := range response {
				b = b + i

				sreamwriter.SetRow(fmt.Sprintf("A%d", b), []interface{}{rep.FirstName})
				sreamwriter.SetRow(fmt.Sprintf("B%d", b), []interface{}{rep.LastName})
				sreamwriter.SetRow(fmt.Sprintf("C%d", b), []interface{}{rep.CreatedBy})
				sreamwriter.SetRow(fmt.Sprintf("D%d", b), []interface{}{rep.CreatedAt})
				sreamwriter.SetRow(fmt.Sprintf("E%d", b), []interface{}{rep.PaidBy})
				sreamwriter.SetRow(fmt.Sprintf("F%d", b), []interface{}{rep.Amount})

			}
			sreamwriter.Flush()
			report_name := fmt.Sprintf("%v_all_reports.xlsx", time.Now().Format(time.ANSIC))

			f.SaveAs("report.xlsx")

			report, err := ioutil.ReadFile("report.xlsx")
			if err != nil {
				continue
			}

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: report_name, Bytes: report}))
		case "today":

			sreamwriter.SetRow("A1", []interface{}{"First Name"})
			sreamwriter.SetRow("B1", []interface{}{"Last Name"})
			sreamwriter.SetRow("C1", []interface{}{"Registered By"})
			sreamwriter.SetRow("D1", []interface{}{"Created At"})
			sreamwriter.SetRow("E1", []interface{}{"Paid By"})
			sreamwriter.SetRow("F1", []interface{}{"Amount"})

			response, err := report.GetAllReports(context.Background())
			if err != nil {
				log.Panicln(err)
				continue
			}
			b := 1
			for _, rep := range response {
				if rep.CreatedAt.Year() == time.Now().Year() && rep.CreatedAt.Month() == time.Now().Month() && rep.CreatedAt.Day() == time.Now().Day() {

					b = b + 1

					sreamwriter.SetRow(fmt.Sprintf("A%d", b), []interface{}{rep.FirstName})
					sreamwriter.SetRow(fmt.Sprintf("B%d", b), []interface{}{rep.LastName})
					sreamwriter.SetRow(fmt.Sprintf("C%d", b), []interface{}{rep.CreatedBy})
					sreamwriter.SetRow(fmt.Sprintf("D%d", b), []interface{}{rep.CreatedAt})
					sreamwriter.SetRow(fmt.Sprintf("E%d", b), []interface{}{rep.PaidBy})
					sreamwriter.SetRow(fmt.Sprintf("F%d", b), []interface{}{rep.Amount})

				}
			}
			sreamwriter.Flush()
			report_name := fmt.Sprintf("%v_daily_reports.xlsx", time.Now().Format(time.ANSIC))

			f.SaveAs("report.xlsx")

			report, err := ioutil.ReadFile("report.xlsx")
			if err != nil {
				continue
			}
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: report_name, Bytes: report}))

		case "week":

			sreamwriter.SetRow("A1", []interface{}{"First Name"})
			sreamwriter.SetRow("B1", []interface{}{"Last Name"})
			sreamwriter.SetRow("C1", []interface{}{"Registered By"})
			sreamwriter.SetRow("D1", []interface{}{"Created At"})
			sreamwriter.SetRow("E1", []interface{}{"Paid By"})
			sreamwriter.SetRow("F1", []interface{}{"Amount"})

			response, err := report.GetAllReports(context.Background())
			if err != nil {
				log.Panicln(err)
				continue
			}
			_, week := time.Now().ISOWeek()
			b := 1
			for _, rep := range response {
				_, respweek := rep.CreatedAt.ISOWeek()
				if rep.CreatedAt.Year() == time.Now().Year() && rep.CreatedAt.Month() == time.Now().Month() && week == respweek {

					b = b + 1

					sreamwriter.SetRow(fmt.Sprintf("A%d", b), []interface{}{rep.FirstName})
					sreamwriter.SetRow(fmt.Sprintf("B%d", b), []interface{}{rep.LastName})
					sreamwriter.SetRow(fmt.Sprintf("C%d", b), []interface{}{rep.CreatedBy})
					sreamwriter.SetRow(fmt.Sprintf("D%d", b), []interface{}{rep.CreatedAt})
					sreamwriter.SetRow(fmt.Sprintf("E%d", b), []interface{}{rep.PaidBy})
					sreamwriter.SetRow(fmt.Sprintf("F%d", b), []interface{}{rep.Amount})

				}
			}
			sreamwriter.Flush()
			report_name := fmt.Sprintf("%v_weekly_reports.xlsx", time.Now().Format(time.ANSIC))

			f.SaveAs("report.xlsx")

			report, err := ioutil.ReadFile("report.xlsx")
			if err != nil {
				continue
			}
			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: report_name, Bytes: report}))

		case "month":

			sreamwriter.SetRow("A1", []interface{}{"First Name"})
			sreamwriter.SetRow("B1", []interface{}{"Last Name"})
			sreamwriter.SetRow("C1", []interface{}{"Registered By"})
			sreamwriter.SetRow("D1", []interface{}{"Created At"})
			sreamwriter.SetRow("E1", []interface{}{"Paid By"})
			sreamwriter.SetRow("F1", []interface{}{"Amount"})

			response, err := report.GetAllReports(context.Background())
			if err != nil {
				log.Panicln(err)
				continue
			}
			b := 1
			for _, rep := range response {

				if rep.CreatedAt.Year() == time.Now().Year() && rep.CreatedAt.Month() == time.Now().Month() {

					b = b + 1

					sreamwriter.SetRow(fmt.Sprintf("A%d", b), []interface{}{rep.FirstName})
					sreamwriter.SetRow(fmt.Sprintf("B%d", b), []interface{}{rep.LastName})
					sreamwriter.SetRow(fmt.Sprintf("C%d", b), []interface{}{rep.CreatedBy})
					sreamwriter.SetRow(fmt.Sprintf("D%d", b), []interface{}{rep.CreatedAt})
					sreamwriter.SetRow(fmt.Sprintf("E%d", b), []interface{}{rep.PaidBy})
					sreamwriter.SetRow(fmt.Sprintf("F%d", b), []interface{}{rep.Amount})

				}
			}
			sreamwriter.Flush()
			report_name := fmt.Sprintf("%v_monthly_reports.xlsx", time.Now().Format(time.ANSIC))

			f.SaveAs("report.xlsx")

			report, err := ioutil.ReadFile("report.xlsx")
			if err != nil {
				continue
			}

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: report_name, Bytes: report}))
		case "year":

			sreamwriter.SetRow("A1", []interface{}{"First Name"})
			sreamwriter.SetRow("B1", []interface{}{"Last Name"})
			sreamwriter.SetRow("C1", []interface{}{"Registered By"})
			sreamwriter.SetRow("D1", []interface{}{"Created At"})
			sreamwriter.SetRow("E1", []interface{}{"Paid By"})
			sreamwriter.SetRow("F1", []interface{}{"Amount"})

			response, err := report.GetAllReports(context.Background())
			if err != nil {
				log.Panicln(err)
				continue
			}
			b := 1
			for _, rep := range response {

				if rep.CreatedAt.Year() == time.Now().Year() {

					b := b + 2

					sreamwriter.SetRow(fmt.Sprintf("A%d", b), []interface{}{rep.FirstName})
					sreamwriter.SetRow(fmt.Sprintf("B%d", b), []interface{}{rep.LastName})
					sreamwriter.SetRow(fmt.Sprintf("C%d", b), []interface{}{rep.CreatedBy})
					sreamwriter.SetRow(fmt.Sprintf("D%d", b), []interface{}{rep.CreatedAt})
					sreamwriter.SetRow(fmt.Sprintf("E%d", b), []interface{}{rep.PaidBy})
					sreamwriter.SetRow(fmt.Sprintf("F%d", b), []interface{}{rep.Amount})

				}
			}
			sreamwriter.Flush()
			report_name := fmt.Sprintf("%v_yearly_reports.xlsx", time.Now().Format(time.ANSIC))

			f.SaveAs("report.xlsx")

			report, err := ioutil.ReadFile("report.xlsx")
			if err != nil {
				continue
			}

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: report_name, Bytes: report}))
		case "backup":
			dbname := fmt.Sprintf("%v_backup.db", time.Now().Format(time.ANSIC))
			report, err := ioutil.ReadFile("gym.db")
			if err != nil {
				continue
			}

			bot.Send(tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FileBytes{Name: dbname, Bytes: report}))

		default:
			msg.Text = `
			Please Use one of these commands 
/report  To Get all the reports 
/today  To Get Today's Report
/week   To Get this week's Report
/month  To Get This month's Report
/year  To Get This Year's Report			
/backup  To Backup The database 				`
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg.Text))
		}
	}
}
