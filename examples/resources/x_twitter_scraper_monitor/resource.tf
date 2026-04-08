resource "x-twitter-scraper_monitor" "example_monitor" {
  event_types = ["tweet.new", "follower.gained"]
  username = "elonmusk"
}
