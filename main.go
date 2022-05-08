/*
High level flow:

-> user makes a new plan on website, can input a list of phone numbers to send invite texts to
-> people who have the link sent in the invite text will be able to join the plan on the website
-> button to vote to cancel the plan (storing some sort of state, want users to be able to change their minds)
-> once plan reaches critical mass for cancellation, plan gets marked as cancelled in the web portal, and a
	text is sent to all participants that the plan has been voted to be cancelled
*/
package main

import "github.com/csdavenport6/cancelplans/sms"

type phoneNumber string

func main() {
	sms.SendMessage("quinn sux")
}
