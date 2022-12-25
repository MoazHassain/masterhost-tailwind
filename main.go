package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/component", component)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/reset-password-step-1", resetPassOne)
	http.HandleFunc("/reset-password-step-2", resetPassTwo)
	http.HandleFunc("/reset-password-step-3", resetPassThree)
	http.HandleFunc("/login-old", loginOld)

	http.HandleFunc("/dashboard", dashboardHome)
	http.HandleFunc("/dashboard-old", dashboardOld)
	http.HandleFunc("/client", clientRoot)
	http.HandleFunc("/client/client-list", clientList)
	http.HandleFunc("/client/subscription", subscription)
	http.HandleFunc("/domain", domainRoot)
	http.HandleFunc("/domain/domain-old", domainOld)
	http.HandleFunc("/domain/domain-list", domain)
	http.HandleFunc("/domain/subdomain-list", subdomain)
	http.HandleFunc("/domain/domain-redirect", domainRedirent)
	http.HandleFunc("/domain/add", addDomain)
	http.HandleFunc("/domain/hosting-selection", hosting)
	http.HandleFunc("/database", databaseRoot)
	http.HandleFunc("/database/database-list", database)
	http.HandleFunc("/database/database-user", databaseUser)
	http.HandleFunc("/database/add", addDatabase)
	http.HandleFunc("/database/add-user", addDbUser)
	http.HandleFunc("/email", emailRoot)
	http.HandleFunc("/email/email-list", emailList)
	http.HandleFunc("/email/forward-email", emailforward)
	http.HandleFunc("/email/add", addEmail)

	http.HandleFunc("/tools", toolsRoot)
	http.HandleFunc("/tools/ssl-manager", sslManager)
	http.HandleFunc("/tools/dns-zone-editor", dnsEditor)
	http.HandleFunc("/tools/add-custom-ssl", addCustomssl)
	http.HandleFunc("/tools/add-auto-ssl", addAutossl)
	http.HandleFunc("/tools/install-ssl", installssl)

	http.HandleFunc("/account", account)
	http.HandleFunc("/account/profile", accountProfile)
	http.HandleFunc("/account/preference", accountPreference)
	http.HandleFunc("/account/authentication", accountAuthentication)
	http.HandleFunc("/account/ssh-keys", accountSSH)
	http.HandleFunc("/account/api", accountAPI)
	http.HandleFunc("/account/users", accountUser)
	http.HandleFunc("/account/notification", accountNotification)

	http.HandleFunc("/packages", packageRoot)
	http.HandleFunc("/packages/package-list", packages)
	http.HandleFunc("/packages/add", addPackages)
	http.HandleFunc("/applications", applicationRoot)
	http.HandleFunc("/applications/application-list", application)
	http.HandleFunc("/billing", billing)
	http.HandleFunc("/billing/payment", billPayment)
	http.HandleFunc("/billing/limit", billLimit)
	http.HandleFunc("/billing/history", billHistory)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./dist/assets"))))

	fmt.Println("starting server on: 8500")
	http.ListenAndServe(":8500", nil)

}

func component(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/components.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/signup.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func login(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/signin.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func resetPassOne(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/forget-pass-s1.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func resetPassTwo(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/forget-pass-s2.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func resetPassThree(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/forget-pass-s3.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func loginOld(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/registration/signin-old.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func dashboardHome(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}
func dashboardOld(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-old.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func clientRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/client/client-list", http.StatusFound)

}
func clientList(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/client/clients.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func subscription(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/client/subscription.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func domainRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/domain/domain-list", http.StatusFound)

}
func domainOld(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-old.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/domain-old.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func domain(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/domain.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func subdomain(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/domain.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/subdomain.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func domainRedirent(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/domain.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/domain-redirect.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func addDomain(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/add-domain.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func hosting(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/domain/hosting-selection.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func packageRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/packages/package-list", http.StatusFound)

}
func packages(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/package/packages.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addPackages(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/package/add-package.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func applicationRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/applications/application-list", http.StatusFound)

}
func application(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/application/application.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func databaseRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/database/database-list", http.StatusFound)

}
func database(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/database/databases.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func databaseUser(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/database/databases.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/database/dbuser.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addDatabase(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/database/add-database.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addDbUser(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/database/add-dbUser.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func emailRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/email/email-list", http.StatusFound)

}
func emailList(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/email/email.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func emailforward(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/email/email.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/email/forward-email.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addEmail(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/email/add-email.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func toolsRoot(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/tools/ssl-manager", http.StatusFound)

}
func sslManager(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/ssl-manager.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func dnsEditor(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/ssl-manager.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/dns-zone-editor.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addCustomssl(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/add-custom-ssl.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addAutossl(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/add-custom-ssl.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/add-auto-ssl.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func installssl(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/add-custom-ssl.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/tools/install-ssl.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func billing(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/billing/payment", http.StatusFound)

}
func billPayment(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/billing/payment.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func billLimit(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/billing/payment.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/billing/limit.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func billHistory(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/billing/payment.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/billing/history.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func account(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/account/profile", http.StatusFound)

}
func accountProfile(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountPreference(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/preference.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountAuthentication(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/authentication.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountSSH(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/ssh-keys.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountAPI(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/api.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountUser(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/users.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func accountNotification(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/account/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("dist/wpage/account/notification.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
