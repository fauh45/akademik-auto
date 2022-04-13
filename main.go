package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()

	defer browser.MustClose()

	login_page := browser.MustPage("https://akademik.polban.ac.id")

	log.Println("Starting to do login on akademik")

	login_page.MustWaitLoad()
	login_page.MustElement(`#fm > div:nth-child(1) > input`).MustInput(os.Getenv("USERNAME"))
	login_page.MustElement(`#fm > div:nth-child(2) > input`).MustInput(os.Getenv("PASSWORD"))

	login_page.MustElement(`#fm > div.row > div.col-xs-4 > button`).MustClick()

	log.Println("Waiting for akademik server to response fully")
	wait, handle := login_page.MustHandleDialog()
	wait()
	handle(true, "")

	login_page.MustWaitLoad().MustClose()

	log.Println("Login complete, closed login page, opening absen page")
	absen_page := browser.MustPage("https://akademik.polban.ac.id/ajar/absen")

	absen_page.MustWaitIdle()

	log.Println("Evaluating Absen Script on the browser")

	absen_page.MustAddScriptTag("https://cdnjs.cloudflare.com/ajax/libs/jquery/3.6.0/jquery.min.js")
	absen_page.MustEval(`
		() => $.each($("#jadwal > tbody > tr:nth-child(n)"),(e,a)=>{const t=$(a),d=t.find("td:eq(1)").text().split("-")[0],n=t.find("td:eq(2)").text(),s=t.find("td:eq(4)").text(),l=t.find("td:eq(5)").text(),o=t.find("td:eq(6)").text(),i=document.getElementById("kls").value;$.ajax({url:"https://akademik.polban.ac.id/ajar/absen/absensi_awal",type:"POST",data:{ja:l,jb:o,mk:n,dsn:d,tp:s,kls:i},success:e=>console.log("Kelas "+n+" dosen "+d+" sudah di absenkan"),error:(e,a,t)=>{console.error("Kelas "+n+" dosen "+d+" gagal di absenkan")}})});
	`)
	time.Sleep(time.Minute)

	log.Println("Waiting for network to go to idle")
	absen_page.MustWaitRequestIdle()
	log.Println("Absen should be done, making sure with a picture")

	absen_page.MustScreenshot(fmt.Sprint("AbsenResult-", time.Now().Format(time.RFC3339), ".png"))
	absen_page.MustClose()

	log.Println("Absen done now, closed absen page, and now closing browser")
}
